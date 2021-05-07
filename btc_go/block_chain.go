package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type Blockchain struct {
	tip []byte
	db  *leveldb.DB
}

func NewBlockChain() (*Blockchain, error) {
	var tip []byte
	// check last block Hash
	// 数据库
	db, err := leveldb.OpenFile("./btc_go.db", nil)
	if err != nil {
		fmt.Printf("open db err : %s\n", err.Error())
		return nil, err
	}
	tip, err = db.Get([]byte("l"), nil)
	if err != nil {
		fmt.Printf("get l err : %s\n", err.Error())
	}
	if len(tip) <= 0 {
		genesis := NewGenesisBlock()
		err = insertBlockDb(db, genesis)
		tip = genesis.GetHead().Hash
		if err != nil {
			return nil, err
		}
	}
	// fmt.Printf("tip is : %x\n", tip)
	return &Blockchain{tip, db}, nil
}

func (bc *Blockchain) AddBlock(data string) error {
	prev, err := bc.db.Get(bc.tip, nil)
	if err != nil {
		fmt.Printf("get tip %s err : %s\n", bc.tip, err.Error())
		return err
	}
	prevBlock, err := DeserializeBlock(prev)
	if err != nil {
		fmt.Printf("DeserializeBlock err : %s\n", err.Error())
		return err
	}
	newBlock := NewBlock(data, prevBlock.GetHead().Hash)
	err = insertBlockDb(bc.db, newBlock)
	if err != nil {
		return err
	}
	bc.tip = newBlock.GetHead().Hash
	return nil
}

func (bc *Blockchain) CloseDb() error {
	err := bc.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func insertBlockDb(db *leveldb.DB, b *Block) error {
	tip := b.GetHead().Hash
	batch := new(leveldb.Batch)
	batch.Put(tip, b.Serialize())
	// fmt.Printf("put hash :%x\n", tip)
	// fmt.Printf("put block :%x\n", b.Serialize())
	batch.Put([]byte("l"), tip)
	err := db.Write(batch, nil)
	if err != nil {
		return err
	}
	return nil
}
