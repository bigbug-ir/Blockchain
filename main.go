package main

import (
	"fmt"

	"github.com/bigbug-ir/blockchain/packages/controller"
	"github.com/bigbug-ir/blockchain/packages/models"
)

func main() {
	blockchiain := models.NewBlockChain()
	fmt.Println(blockchiain.GetChain())
	controller.CreateBlock(blockchiain)
	controller.CreateBlock(blockchiain)
	controller.CreateBlock(blockchiain)
	controller.CreateBlock(blockchiain)
	controller.CreateBlock(blockchiain)
	controller.CreateBlock(blockchiain)
	controller.CreateBlock(blockchiain)
	fmt.Println(blockchiain.GetChain())
}
