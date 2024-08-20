# Blockchain Project

This is a basic blockchain implementation in Go. The project includes functionalities for creating blocks, maintaining a blockchain, handling transactions, and computing Merkle trees.

## Project Structure

- `main.go`: Entry point for the application. Initializes the blockchain, creates new blocks, and prints the chain.
- `models/`: Contains the data models for blocks, transactions, and the blockchain.
- `controller/`: Contains functions for creating blocks, validating blocks, and computing Merkle trees.

## Getting Started

To run the project, ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/).

### Prerequisites

- Go 1.18 or higher

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/blockchain.git
cd blockchain
```

### Running the Project
To run the project, navigate to the root directory of the project and execute:
```bash 
    go run .
```
### Project Details
#### `main.go`
The entry point of the application that initializes the blockchain and creates a series of blocks:
``` go
package main

import (
	"fmt"

	"github.com/bigbug-ir/blockchain/packages/controller"
	"github.com/bigbug-ir/blockchain/packages/models"
)

func main() {
	blockchain := models.NewBlockChain()
	fmt.Println(blockchain.GetChain())
	controller.CreateBlock(blockchain)
	controller.CreateBlock(blockchain)
	controller.CreateBlock(blockchain)
	controller.CreateBlock(blockchain)
	controller.CreateBlock(blockchain)
	controller.CreateBlock(blockchain)
	controller.CreateBlock(blockchain)
	fmt.Println(blockchain.GetChain())
}

```
#### `./package/models/block.go`
Defines the `Block` structure and functions related to blocks, including hash computation and block creation:
``` go
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
	// Implementation for saving the block
}

func (block *Block) SetHash() {
	// Computes and sets the hash of the block
}

func GenesisBlock() *Block {
	// Returns the genesis block
}

```
#### `./package/models/blockchain.go`
Contains the `Blockchain` structure and functions for manipulating the blockchain:
``` go
package models

type Blockchain struct {
	Chain []*Block
}

func (blockchain *Blockchain) AddBlock(newBlock *Block) {
	blockchain.Chain = append(blockchain.Chain, newBlock)
}

func (chain *Blockchain) GetChain() []*Block {
	return chain.Chain
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

```
#### `./package/controller/transaction.go`
Defines the `Transaction` structure used in blocks:
``` go
package models

import "time"

type Transaction struct {
	ID        string
	Amount    int
	Currency  string
	From      string
	To        string
	Timestamp time.Time
}

```
#### `./package/controller/block.go`
Contains functions for creating new blocks, computing Merkle trees, and validating blocks:
```go 
package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"github.com/bigbug-ir/blockchain/packages/models"
)

func hash(data []byte) [32]byte {
	// Computes SHA-256 hash of the data
}

func merkleTree(data [][]byte) [][32]byte {
	// Computes the Merkle tree for the given data
}

func CreateBlock(blockchain *models.Blockchain) {
	// Creates and adds a new block to the blockchain
}

func getTransactions() []models.Transaction {
	// Returns a list of sample transactions
}

func ValidateBlock(block *models.Block, blockchain *models.Blockchain) error {
	// Validates the block based on blockchain rules
}
```
## Contributing
Feel free to open issues or submit pull requests to contribute to the project.
## Contact
For any questions, please contact [amirhosseinbakhtiyari.program@gmail.com] or open an issue on [https://github.com/bigbug-ir/Blockchain/].


### توضیحات

- **پیش‌نیازها**: مشخص می‌کند که برای اجرای پروژه به چه نسخه‌ای از Go نیاز دارید.
- **نصب و اجرا**: راهنمای نصب و اجرای پروژه را توضیح می‌دهد.
- **جزئیات پروژه**: شامل توضیحات مربوط به هر فایل و عملکرد آن در پروژه است.
- **مشارکت**: کاربران را تشویق به مشارکت و ارسال مشکلات یا درخواست‌های تغییر می‌کند.

با استفاده از این فایل `README.md`، کاربران می‌توانند اطلاعات دقیقی در مورد نحوه استفاده و جزئیات پروژه بلاکچین شما کسب کنند.
