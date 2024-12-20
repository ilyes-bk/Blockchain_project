package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Struct for KPI results
type KPIMetrics struct {
	TransactionThroughput float64
	Latency               float64
	Scalability           string
	Security              string
	CostEfficiency        string
}

// Energy structure to represent listed energy
type Energy struct {
	ID        *big.Int
	Seller    string
	Amount    *big.Int
	Price     *big.Int
	Purchased bool // Track if the energy is purchased
}

var energies = make(map[*big.Int]Energy)

// Ledger entry to record transactions
type Transaction struct {
	TxHash    string
	Type      string // "list" or "purchase"
	Amount    *big.Int
	Price     *big.Int
	Buyer     string // Only relevant for purchase transactions
	Timestamp time.Time
}

var ledger []Transaction // Ledger to store all transactions

// MeasureTransactionThroughput calculates transactions per second
func MeasureTransactionThroughput(totalTransactions int, duration time.Duration) float64 {
	return float64(totalTransactions) / duration.Seconds()
}

// MeasureLatency calculates average latency
func MeasureLatency(startTime time.Time, endTime time.Time) float64 {
	return endTime.Sub(startTime).Seconds()
}

// AssessScalability simulates scalability assessment
func AssessScalability(loadLevel int) string {
	if loadLevel < 1000 {
		return "High scalability"
	} else if loadLevel < 5000 {
		return "Moderate scalability"
	} else {
		return "Low scalability"
	}
}

// EvaluateSecurity checks basic conditions (simulated)
func EvaluateSecurity(isDataSecure bool) string {
	if isDataSecure {
		return "High security"
	}
	return "Low security"
}

// AnalyzeCostEfficiency simulates cost analysis
func AnalyzeCostEfficiency(cost float64, value float64) string {
	if value/cost > 1 {
		return "Cost-efficient"
	}
	return "Not cost-efficient"
}

const contractAddress = "0xF56aE9cd0E27c8Fe556558820ac63E782461556E"
const privateKeyHex = "12b75d8098b3fe0483fe957c0e10769a9b9254b997cdbdac7efc8d4376b1cc0b"
const contractABI = `[{"constant":false,"inputs":[{"name":"_amount","type":"uint256"},{"name":"_price","type":"uint256"}],"name":"listEnergy","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_energyId","type":"uint256"}],"name":"purchaseEnergy","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"_energyId","type":"uint256"}],"name":"energies","outputs":[{"name":"seller","type":"address"},{"name":"amount","type":"uint256"},{"name":"price","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

var kpi KPIMetrics // Global variable for KPI metrics

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	chainID := big.NewInt(1337) // Chain ID for Ganache
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	contractAddr := common.HexToAddress(contractAddress)

	for {
		// User Input Menu
		fmt.Println("\nSelect an action:")
		fmt.Println("1. List Energy")
		fmt.Println("2. View Listed Energies")
		fmt.Println("3. Purchase Energy")
		fmt.Println("4. Exit")
		fmt.Println("5. View Ledger")
		fmt.Println("6. View KPI Metrics")
		fmt.Println("7. Change KPI Parameters")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// List Energy for sale
			var energyAmount, energyPrice int64
			fmt.Println("Enter energy amount to list (in kWh):")
			fmt.Scan(&energyAmount)
			fmt.Println("Enter energy price (in Wei):")
			fmt.Scan(&energyPrice)

			energyID := big.NewInt(time.Now().Unix()) // Generate a unique ID for each energy listing (use Unix timestamp)

			// Store energy in the map
			energies[energyID] = Energy{
				ID:        energyID,
				Seller:    fromAddress.Hex(),
				Amount:    big.NewInt(energyAmount),
				Price:     big.NewInt(energyPrice),
				Purchased: false, // Initially, the energy is not purchased
			}

			startTime := time.Now()
			listEnergyData, err := parsedABI.Pack("listEnergy", big.NewInt(energyAmount), big.NewInt(energyPrice))
			if err != nil {
				log.Fatalf("Failed to pack listEnergy data: %v", err)
			}

			tx := types.NewTransaction(nonce, contractAddr, big.NewInt(0), auth.GasLimit, auth.GasPrice, listEnergyData)
			signedTx, err := auth.Signer(auth.From, tx)
			if err != nil {
				log.Fatalf("Failed to sign transaction: %v", err)
			}

			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatalf("Failed to send transaction: %v", err)
			}
			endTime := time.Now()
			fmt.Printf("Energy listed successfully. Transaction hash: %s\n", signedTx.Hash().Hex())

			// Add to ledger
			ledger = append(ledger, Transaction{
				TxHash:    signedTx.Hash().Hex(),
				Type:      "list",
				Amount:    big.NewInt(energyAmount),
				Price:     big.NewInt(energyPrice),
				Timestamp: endTime,
			})

			// Measure KPI metrics
			kpi.TransactionThroughput = MeasureTransactionThroughput(1, endTime.Sub(startTime))
			kpi.Latency = MeasureLatency(startTime, endTime)
			kpi.Scalability = AssessScalability(1000)
			kpi.Security = EvaluateSecurity(true)
			kpi.CostEfficiency = AnalyzeCostEfficiency(float64(energyPrice), float64(energyAmount))

		case 2:
			// List all available energy (not purchased)
			fmt.Println("\nListed Energies (Not Purchased):")
			for id, energy := range energies {
				if !energy.Purchased {
					fmt.Printf("ID: %d, Seller: %s, Amount: %d kWh, Price: %d Wei\n", id, energy.Seller, energy.Amount, energy.Price)
				}
			}

		case 3:
			// Purchase Energy
			var energyID int64
			fmt.Println("Enter energy ID to purchase:")
			fmt.Scan(&energyID)

			// Find the energy and mark it as purchased
			for id, energy := range energies {
				if id.Int64() == energyID && !energy.Purchased {
					// Mark energy as purchased
					energy.Purchased = true
					energies[id] = energy

					nonce++
					purchaseEnergyData, err := parsedABI.Pack("purchaseEnergy", id)
					if err != nil {
						log.Fatalf("Failed to pack purchaseEnergy data: %v", err)
					}

					tx := types.NewTransaction(nonce, contractAddr, big.NewInt(0), auth.GasLimit, auth.GasPrice, purchaseEnergyData)
					signedTx, err := auth.Signer(auth.From, tx)
					if err != nil {
						log.Fatalf("Failed to sign transaction: %v", err)
					}

					err = client.SendTransaction(context.Background(), signedTx)
					if err != nil {
						log.Fatalf("Failed to send transaction: %v", err)
					}
					fmt.Printf("Energy purchased successfully. Transaction hash: %s\n", signedTx.Hash().Hex())

					// Add purchase transaction to ledger
					ledger = append(ledger, Transaction{
						TxHash:    signedTx.Hash().Hex(),
						Type:      "purchase",
						Amount:    energy.Amount,
						Price:     energy.Price,
						Buyer:     fromAddress.Hex(),
						Timestamp: time.Now(),
					})
					break
				}
			}

		case 4:
			// Exit
			fmt.Println("Exiting...")
			return

		case 5:
			// View Ledger
			fmt.Println("\nTransaction Ledger:")
			for _, tx := range ledger {
				fmt.Printf("TxHash: %s, Type: %s, Amount: %d kWh, Price: %d Wei, Timestamp: %s\n",
					tx.TxHash, tx.Type, tx.Amount, tx.Price, tx.Timestamp)
			}

		case 6:
			// View KPI Metrics
			fmt.Printf("\nCurrent KPI Metrics:\n")
			fmt.Printf("Transaction Throughput: %.2f transactions/sec\n", kpi.TransactionThroughput)
			fmt.Printf("Latency: %.2f seconds\n", kpi.Latency)
			fmt.Printf("Scalability: %s\n", kpi.Scalability)
			fmt.Printf("Security: %s\n", kpi.Security)
			fmt.Printf("Cost Efficiency: %s\n", kpi.CostEfficiency)

		case 7:
			// Change KPI Parameters
			var throughput, latency, scalability, security, costEfficiency int
			fmt.Println("\nEnter the following KPI parameters:")

			fmt.Print("Transaction Throughput (in transactions/sec): ")
			fmt.Scan(&throughput)
			kpi.TransactionThroughput = float64(throughput)

			fmt.Print("Latency (in seconds): ")
			fmt.Scan(&latency)
			kpi.Latency = float64(latency)

			fmt.Print("Scalability (0: Low, 1: Moderate, 2: High): ")
			fmt.Scan(&scalability)
			if scalability == 0 {
				kpi.Scalability = "Low scalability"
			} else if scalability == 1 {
				kpi.Scalability = "Moderate scalability"
			} else {
				kpi.Scalability = "High scalability"
			}

			fmt.Print("Security (0: Low, 1: High): ")
			fmt.Scan(&security)
			if security == 0 {
				kpi.Security = "Low security"
			} else {
				kpi.Security = "High security"
			}

			fmt.Print("Cost Efficiency (0: Not cost-efficient, 1: Cost-efficient): ")
			fmt.Scan(&costEfficiency)
			if costEfficiency == 0 {
				kpi.CostEfficiency = "Not cost-efficient"
			} else {
				kpi.CostEfficiency = "Cost-efficient"
			}

			fmt.Println("\nKPI parameters updated successfully.")

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
