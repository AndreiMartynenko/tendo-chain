package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

//Adding a new block to the blockchain

type Block struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  []byte
	Hash      []byte
}

// Blockchain structure
type Blockchain struct {
	Blocks []*Block
}

// hash for the block
func (b *Block) GenerateHash() {
	data := bytes.Join([][]byte{
		[]byte(fmt.Sprintf("%d", b.Index)),
		[]byte(fmt.Sprintf("%d", b.Timestamp)),
		[]byte(b.Data),
		b.PrevHash,
	}, []byte{})
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}

// new block
func NewBlock(data string, prevBlock *Block) *Block {
	block := &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	block.GenerateHash()
	return block
}

// Genesis block
func GenesisBlock() *Block {
	return &Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		Data:      "Genesis Block",
		Hash:      []byte{},
	}
}

// Add a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Print blockchain
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d, Data: %s, Hash: %x\n", block.Index, block.Data, block.Hash)
	}
}

func main() {
	// Create a blockchain and add blocks
	bc := &Blockchain{Blocks: []*Block{GenesisBlock()}}
	bc.AddBlock("Block 1")
	bc.AddBlock("Block 2")
	bc.Print()
}
