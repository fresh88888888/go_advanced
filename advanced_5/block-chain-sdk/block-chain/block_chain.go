package blockchain

import (
	"container/list"
	"time"
)

// Type BlockChain represent a Blockchain
type BlockChain struct {
	chain *list.List
}

func NewBlockchain(initialBlock *Block) *BlockChain {
	chain := list.New()
	chain.PushBack(initialBlock)

	return &BlockChain{chain: chain}
}

func (bc *BlockChain) AddBlock(t Transaction) error {
	newBlock := &Block{}
	currentBlock := bc.GetCurrentBlock()

	newBlock.Index = currentBlock.Index
	newBlock.Transaction = &t
	newBlock.Timestamp = time.Now().String()
	newBlock.PreviousHash = currentBlock.Hash
	bHash, err := GetBlockHash(*newBlock)
	if err != nil {
		return err
	}
	newBlock.Hash = bHash
	bc.chain.PushBack(newBlock)

	return nil
}

func (bc *BlockChain) GetCurrentBlock() *Block {
	b := bc.chain.Back().Value.(*Block)

	return b
}
