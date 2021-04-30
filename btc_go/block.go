package main

import (
	"time"
)

const targetBits = 16

// Block
type Block struct {
	Head         *BlockHead
	Transactions []byte
}

// BlockHead
type BlockHead struct {
	// Version of the block.  This is not the same as the protocol version.
	Version int32

	// Hash of the previous block in the block chain.
	PrevBlockHash []byte

	Hash []byte

	// Time the block was created.  This is, unfortunately, encoded as a
	// uint32 on the wire and therefore is limited to 2106.
	Timestamp time.Time

	// Difficulty target for the block.
	Bits uint32

	// Nonce used to generate the block.
	Nonce uint32
}

// 创建创世块，第一个块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// 创建新块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := new(Block)
	block.Head = new(BlockHead)
	block.Head.Version = 0x10000000
	block.Head.Bits = targetBits
	block.Head.PrevBlockHash = prevBlockHash
	block.Head.Timestamp = time.Now()
	block.Transactions = []byte(data)
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Head.Nonce = uint32(nonce)
	block.Head.Hash = hash[:]
	return block
}

// // 为当前的Block增加Hash
// func (b *Block) SetHash() {
// 	timestamp := []byte(b.Head.Timestamp.String())
// 	data := bytes.Join([][]byte{b.Head.PrevBlockHash, b.Transactions, timestamp}, []byte{})
// 	hash := sha256.Sum256(data)
// 	b.Head.Hash = hash[:]
// }

// 获取头部
func (b *Block) GetHead() *BlockHead {
	if b != nil {
		return b.Head
	}
	return nil
}

// 获取头部
func (b *Block) GetTransactions() []byte {
	if b != nil {
		return b.Transactions
	}
	return nil
}
