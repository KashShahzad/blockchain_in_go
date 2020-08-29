package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//elements of block
type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
}

//creating a blockchain type
type Blockchain struct {
	blocks []*Block
}

//creation of hash
func (b *Block) DriveHash() {
	//first we need to remove spaces in bytes
	info := bytes.Join([][]byte{b.Data, b.PreHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//creating a block
func CreateBlock(data string, preHash []byte) *Block {
	//block var using address of Block and taking empty slice of bytes for hash and data which is converted in bytes and prehash
	block := &Block{[]byte{}, []byte(data), preHash}
	//getting hash cooked from Drivehash method
	block.DriveHash()
	return block
}

//method allowing us to add a block to chain
func (chain *Blockchain) AddBlock(data string) {
	//getting previous block in our chain
	prevBlock := chain.blocks[len(chain.blocks)-1]
	//creating a new block var
	new := CreateBlock(data, prevBlock.Hash)
	//adding it to the chain
	chain.blocks = append(chain.blocks, new)
}

//we need to add a genesis block and bC to have something for the start
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	//to see all the blocks
	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
