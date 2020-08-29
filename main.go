package main

import (
	"block-go/blockchain"
	"fmt"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	//to see all the blocks
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
