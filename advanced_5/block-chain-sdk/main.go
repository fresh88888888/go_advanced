package main

import (
	"fmt"

	blockchain "umbrella.github.com/advanced_go/advanced_5/block-chain-sdk/block-chain"
)

func main() {
	bc := blockchain.NewBlockchain(blockchain.NewGenesisBlock())
	fmt.Println(bc.GetCurrentBlock().Hash)
	fmt.Println(blockchain.GetTransactionHash(*bc.GetCurrentBlock().Transaction))
	bc.AddBlock(*blockchain.NewTransaction([]byte{4, 5, 6, 7}))
	fmt.Println(bc.GetCurrentBlock().PreviousHash)
	fmt.Println(bc.GetCurrentBlock().Hash)
}
