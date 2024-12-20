// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EnergyTrading {
    // Structure to store energy listing details
    struct Energy {
        uint256 id;
        address seller;
        uint256 amount; // in kWh
        uint256 price;  // in Wei
        bool sold;
        address buyer;  // Track buyer address after purchase
    }

    // Mapping to store energy listings by their unique id
    mapping(uint256 => Energy) public energies;

    // Variable to track the next energy listing id
    uint256 public nextId;

    // Event emitted when energy is listed for sale
    event EnergyListed(uint256 id, address seller, uint256 amount, uint256 price);

    // Event emitted when energy is purchased
    event EnergyPurchased(uint256 id, address buyer);

    // Function to list energy for sale
    function listEnergy(uint256 amount, uint256 price) external {
        require(amount > 0, "Amount must be greater than 0");
        require(price > 0, "Price must be greater than 0");

        // Store the new energy listing
        energies[nextId] = Energy({
            id: nextId,
            seller: msg.sender,
            amount: amount,
            price: price,
            sold: false,
            buyer: address(0) // No buyer initially
        });

        // Emit the EnergyListed event
        emit EnergyListed(nextId, msg.sender, amount, price);

        // Increment the nextId for the next energy listing
        nextId++;
    }

    // Function to purchase energy
    function purchaseEnergy(uint256 id) external payable {
        Energy storage energy = energies[id];

        // Ensure the energy is available and not already sold
        require(!energy.sold, "Energy already sold");
        require(msg.value == energy.price, "Incorrect value sent");

        // Mark the energy as sold
        energy.sold = true;
        energy.buyer = msg.sender; // Store the buyer's address

        // Transfer the value to the seller
        (bool success, ) = energy.seller.call{value: msg.value}("");
        require(success, "Payment failed");

        // Emit the EnergyPurchased event
        emit EnergyPurchased(id, msg.sender);
    }

    // Function to get the details of a specific energy listing
    function getEnergy(uint256 id) external view returns (Energy memory) {
        return energies[id];
    }

    // Optional: Withdraw contract balance (for the owner to withdraw funds if needed)
    function withdraw() external {
        // Ensure only the contract owner can withdraw funds
        // Here we assume the owner is the account that deployed the contract
        require(msg.sender == address(this), "Not authorized to withdraw");

        // Transfer balance to the owner
        payable(msg.sender).transfer(address(this).balance);
    }
}
