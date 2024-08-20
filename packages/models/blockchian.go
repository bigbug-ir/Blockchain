package models

type Blockchain struct {
	Chain []*Block
}

func (blockchain *Blockchain) AddBlock(newblock *Block) {
	blockchain.Chain = append(blockchain.Chain, newblock)
}
func (chain *Blockchain) GetChain() []*Block {
	return chain.Chain
}
func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
