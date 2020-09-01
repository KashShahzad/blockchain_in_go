package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

//elements of block
type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
	Nonce   int
}

// //creation of hash
// func (b *Block) DriveHash() {
// 	//first we need to remove spaces in bytes
// 	info := bytes.Join([][]byte{b.Data, b.PreHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

//creating a block
func CreateBlock(data string, preHash []byte) *Block {
	//block var using address of Block and taking empty slice of bytes for hash and data which is converted in bytes and prehash
	block := &Block{[]byte{}, []byte(data), preHash, 0}
	//getting hash cooked from Drivehash method
	// block.DriveHash()
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//we need to add a genesis block and bC to have something for the start
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//for serialization on block elem
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

//for deserialization on the byte return from db
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}