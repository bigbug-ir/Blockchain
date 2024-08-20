package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/bigbug-ir/blockchain/packages/models"
)

func hash(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func merkleTree(data [][]byte) [][32]byte {
	// Initial level: leaf nodes
	var leaves [][32]byte
	for _, d := range data {
		leaves = append(leaves, hash(d))
	}

	// Building the tree level by level
	var currentLevel [][32]byte
	for len(leaves) > 1 {
		currentLevel = nil
		for i := 0; i < len(leaves); i += 2 {
			if i+1 < len(leaves) {
				concat := append(leaves[i][:], leaves[i+1][:]...)
				currentLevel = append(currentLevel, hash(concat))
			} else {
				// Handle the case with odd number of leaves
				currentLevel = append(currentLevel, leaves[i])
			}
		}
		leaves = currentLevel
	}

	if len(leaves) == 0 {
		// If the tree was empty, create a root node with empty hash
		return [][32]byte{{}}
	}

	return leaves
}

func CreateBlock(blockchian *models.Blockchain) {
	// Create a new block
	rand.Seed(time.Now().UnixNano())
	block := models.Block{}
	block.Transactions = getTransactions()
	var transactionData [][]byte
	Timestamp := time.Now().Unix()
	PrevHash := blockchian.Chain[len(blockchian.Chain)-1].BlockHash
	Nonce := rand.Int63n(10)
	block = models.Block{
		ID:        blockchian.Chain[len(blockchian.Chain)-1].ID + 1,
		Timestamp: Timestamp,
		PrevHash:  PrevHash,
		Nonce:     Nonce,
	}
	for _, tx := range block.Transactions {
		txData := []byte(tx.From + tx.To + string(tx.Amount) + tx.Currency)
		transactionData = append(transactionData, txData)
	}
	Root := merkleTree(transactionData)
	if len(Root) > 0 {
		block.Root = Root[0]
	}
	block.SetHash()
	fmt.Println(hex.EncodeToString(block.BlockHash[:]), block.Timestamp, block.Nonce, block.BlockHash, block.PrevHash)
	blockchian.Chain = append(blockchian.Chain, &block)
}
func getTransactions() []models.Transaction {
	// Get all transactions
	transaction := []models.Transaction{
		{
			Amount:   100,
			From:     "Alice",
			To:       "Bob",
			Currency: "bitcoin",
		},
		{
			Amount:   200,
			From:     "Bob",
			To:       "Charlie",
			Currency: "Ethereum",
		},
	}
	return transaction
}

func ValidateBlock(block *models.Block, blockchain *models.Blockchain) error {
	// Validate the block
	if block.Timestamp < blockchain.Chain[len(blockchain.Chain)-1].Timestamp {
		return errors.New("invalid timestamp")
	}
	if block.PrevHash != blockchain.Chain[len(blockchain.Chain)-1].BlockHash {
		return errors.New("invalid previous hash")
	}
	var testBlock models.Block
	testBlock.Timestamp = block.Timestamp
	testBlock.PrevHash = block.PrevHash
	testBlock.Root = block.Root
	testBlock.Nonce = block.Nonce
	testBlock.SetHash()
	if testBlock.BlockHash != block.BlockHash {
		return errors.New("invalid hash")
	}
	return nil
}
