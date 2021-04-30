package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-b.GetHead().Bits))
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.GetBlock().GetHead().PrevBlockHash,
			pow.GetBlock().GetTransactions(),
			IntToHex(pow.GetBlock().GetHead().Timestamp.Unix()),
			IntToHex(int64(pow.GetBlock().GetHead().Bits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) GetTarget() *big.Int {
	if pow != nil {
		return pow.target
	}
	return nil
}

func (pow *ProofOfWork) GetBlock() *Block {
	if pow != nil {
		return pow.block
	}
	return nil
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.GetBlock().GetTransactions())
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.GetTarget()) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}
