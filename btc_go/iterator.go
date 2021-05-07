package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type BlockchainIterator struct {
	currentHash []byte
	db          *leveldb.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

func (i *BlockchainIterator) Next() *Block {
	block, err := i.db.Get(i.currentHash, nil)
	if err != nil {
		fmt.Printf("get tip %x err : %s\n", i.currentHash, err.Error())
		return nil
	}
	currentBlock, err := DeserializeBlock(block)
	if err != nil {
		fmt.Printf("DeserializeBlock err : %s\n", err.Error())
		return nil
	}

	i.currentHash = currentBlock.GetHead().PrevBlockHash

	return currentBlock
}
