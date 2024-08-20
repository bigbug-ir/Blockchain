package models

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	ID           uint
	Timestamp    int64
	PrevHash     [32]byte
	BlockHash    [32]byte
	Nonce        int64
	Root         [32]byte
	Transactions []Transaction
}

func SaveBlock(block *Block) (newBlock Block, err error) {
	newBlock = Block{
		ID:           block.ID,
		Timestamp:    time.Now().Unix(),
		PrevHash:     block.PrevHash,
		BlockHash:    block.BlockHash,
		Nonce:        block.Nonce,
		Root:         block.Root,
		Transactions: make([]Transaction, 0),
	}
	// I can add to database

	return newBlock, nil
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	nonce := []byte(strconv.FormatInt(int64(block.Nonce), 10))
	headrs := bytes.Join([][]byte{timestamp, block.PrevHash[:], block.Root[:], nonce}, []byte{})
	hash := sha256.Sum256(headrs)
	block.BlockHash = hash
}

func GenesisBlock() *Block {
	var transactions []Transaction

	genesisTx := Transaction{
		From:   "AmirHossein",
		To:     "the world",
		Amount: 100,
	}
	transactions = append(transactions, genesisTx)

	genesisBlock := &Block{
		Timestamp:    time.Now().Unix(),
		BlockHash:    sha256.Sum256([]byte("genesisBlock")),
		PrevHash:     sha256.Sum256([]byte("0")),
		Transactions: transactions,
	}

	return genesisBlock
}
