package blockchain

//elements of block
type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
	Nonce   int
}

//creating a blockchain type
type Blockchain struct {
	Blocks []*Block
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

//method allowing us to add a block to chain
func (chain *Blockchain) AddBlock(data string) {
	//getting previous block in our chain
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	//creating a new block var
	new := CreateBlock(data, prevBlock.Hash)
	//adding it to the chain
	chain.Blocks = append(chain.Blocks, new)
}

//we need to add a genesis block and bC to have something for the start
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
