package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math/big"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// Validate 证明工作量是否是正确的，验证别人的block
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(int(pow.block.GetHead().Nonce))
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}
