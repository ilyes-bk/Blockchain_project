package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Transaction represents a single blockchain transaction
type Transaction struct {
	ID        string
	From      string
	To        string
	Amount    float64
	Timestamp time.Time
	Signature string
}

// Block structure with added features
type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []Transaction
	Data         string
	PreviousHash string
	Hash         string
	Nonce        int
	Difficulty   int
}

// Blockchain structure
type Blockchain struct {
	Chain         []Block
	ChainName     string
	CreateTime    time.Time
	PendingTxs    []Transaction
	Difficulty    int
	MiningReward  float64
	ProcessedTxs  int
	TotalTxs      int
	TotalMineTime time.Duration
	mu            sync.RWMutex
}

// Calculate the hash of a block
func (b *Block) calculateHash() string {
	record := fmt.Sprintf("%d%v%v%s%s%d%d",
		b.Index,
		b.Timestamp,
		b.Transactions,
		b.Data,
		b.PreviousHash,
		b.Nonce,
		b.Difficulty,
	)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// Proof of Work implementation
func (b *Block) mineBlock(difficulty int) {
	target := strings.Repeat("0", difficulty)
	for {
		b.Hash = b.calculateHash()
		if strings.HasPrefix(b.Hash, target) {
			break
		}
		b.Nonce++
	}
}

// Create a new block
func createBlock(index int, data string, previousHash string, transactions []Transaction, difficulty int) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now(),
		Data:         data,
		Transactions: transactions,
		PreviousHash: previousHash,
		Difficulty:   difficulty,
	}
	block.Hash = block.calculateHash()
	return block
}

// Validate the blockchain
func (bc *Blockchain) validate() bool {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}

		if currentBlock.Hash != currentBlock.calculateHash() {
			return false
		}
	}
	return true
}

// Add a new block to the chain
func (bc *Blockchain) addBlock(data string) (time.Duration, error) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if len(bc.Chain) == 0 {
		return 0, errors.New("blockchain is empty")
	}

	start := time.Now()
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := createBlock(prevBlock.Index+1, data, prevBlock.Hash, bc.PendingTxs, bc.Difficulty)

	// Mine the block
	newBlock.mineBlock(bc.Difficulty)

	// Update blockchain stats
	bc.ProcessedTxs += len(bc.PendingTxs)
	bc.TotalTxs += len(bc.PendingTxs)
	bc.TotalMineTime += time.Since(start)

	// Add block to chain and clear pending transactions
	bc.Chain = append(bc.Chain, newBlock)
	bc.PendingTxs = []Transaction{}

	return time.Since(start), nil
}

// Create a new transaction
func createTransaction(from, to string, amount float64) Transaction {
	tx := Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	// Generate transaction ID using hash
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s%f%s", from, to, amount, tx.Timestamp)))
	tx.ID = hex.EncodeToString(h.Sum(nil))[:8]
	return tx
}

// Initialize a new blockchain
func createBlockchain(name string) *Blockchain {
	genesisBlock := createBlock(0, "Genesis Block", "", []Transaction{}, 4)
	return &Blockchain{
		Chain:        []Block{genesisBlock},
		ChainName:    name,
		CreateTime:   time.Now(),
		Difficulty:   4,
		MiningReward: 10.0,
		PendingTxs:   []Transaction{},
	}
}

// Aggregate data from sub-chains
func aggregateData(primaryChain *Blockchain, subChains ...*Blockchain) time.Duration {
	start := time.Now()
	var aggregatedData string

	for _, chain := range subChains {
		if len(chain.Chain) > 0 {
			latestBlock := chain.Chain[len(chain.Chain)-1]
			aggregatedData += fmt.Sprintf("%s: %s | ", chain.ChainName, latestBlock.Data)
		}
	}

	primaryChain.addBlock(aggregatedData)
	return time.Since(start)
}

// Export KPIs to JSON
func exportKPIs(chain *Blockchain) {
	kpis := map[string]interface{}{
		"Chain Name":           chain.ChainName,
		"Total Blocks":         len(chain.Chain),
		"Total Transactions":   chain.TotalTxs,
		"Processed Transactions": chain.ProcessedTxs,
		"Total Mining Time":    chain.TotalMineTime.Seconds(),
		"Average Mining Time":  chain.TotalMineTime.Seconds() / float64(len(chain.Chain)),
		"Mining Difficulty":    chain.Difficulty,
		"Mining Reward":        chain.MiningReward,
	}
	jsonData, err := json.MarshalIndent(kpis, "", "  ")
	if err != nil {
		fmt.Println("Error exporting KPIs:", err)
		return
	}
	fmt.Println("\nBlockchain KPIs:")
	fmt.Println(string(jsonData))
}

func main() {
    // Initialize blockchains
    primaryChain := createBlockchain("City Transport")
    tollChain := createBlockchain("Toll System")
    rideChain := createBlockchain("Ride-Sharing")

    // Create and add transactions
    tollTx := createTransaction("Vehicle123", "TollBoothA", 5.0)
    rideTx := createTransaction("RiderX", "DriverY", 15.0)

    tollChain.PendingTxs = append(tollChain.PendingTxs, tollTx)
    rideChain.PendingTxs = append(rideChain.PendingTxs, rideTx)

    // Add blocks to sub-chains
    tollChain.addBlock("Toll paid by Vehicle ID 123 for $5")
    rideChain.addBlock("Ride completed by Driver ID 456 for $15")

    // Aggregate data to the primary chain
    aggregateData(primaryChain, tollChain, rideChain)

    // Print blockchain states
    fmt.Println("Primary Blockchain:")
    for _, block := range primaryChain.Chain {
        fmt.Printf("Index: %d, Data: %s, Hash: %s\n", block.Index, block.Data, block.Hash)
    }

    fmt.Println("\nToll Sub-Chain:")
    for _, block := range tollChain.Chain {
        fmt.Printf("Index: %d, Data: %s, Hash: %s\n", block.Index, block.Data, block.Hash)
    }

    fmt.Println("\nRide-Sharing Sub-Chain:")
    for _, block := range rideChain.Chain {
        fmt.Printf("Index: %d, Data: %s, Hash: %s\n", block.Index, block.Data, block.Hash)
    }

    // Export KPIs for the primary chain
    exportKPIs(primaryChain)
}

