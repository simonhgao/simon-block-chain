package main

import (
	"log"
)

func main() {
	println("Let's start a demo block chain example")
	bc, err := NewBlockChain()
	if err != nil {
		println("create block chain err")
		log.Panic(err.Error())
	}
	println("create block chain success")
	// 延迟关闭数据库
	defer func() {
		err := bc.CloseDb()
		if err != nil {
			log.Panic(err.Error())
		}
	}()
	// err = bc.AddBlock("Send 1 BTC to simon")
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// err = bc.AddBlock("Send 2 more BTC to kevin")
	// if err != nil {
	// 	log.Panic(err.Error())
	// }

	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. Hash: %x\n", block.GetHead().PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Transactions)
	// 	fmt.Printf("Hash: %x\n", block.GetHead().Hash)
	// 	fmt.Println()
	// }

	// iter := bc.db.NewIterator(nil, nil)
	// for ok := iter.Seek(bc.tip); ok; ok = iter.Next() {
	// 	data := iter.Value()
	// 	if string(data) == string(bc.tip) {
	// 		break
	// 	}
	// 	lastBlock, err := DeserializeBlock(data)
	// 	if err != nil {
	// 		log.Panic(err.Error())
	// 	}
	// 	fmt.Printf("Prev. Hash: %x\n", lastBlock.GetHead().PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", lastBlock.Transactions)
	// 	fmt.Printf("Hash: %x\n", lastBlock.GetHead().Hash)
	// }
	// iter.Release()
	// err = iter.Error()

	cli := CLI{bc}
	cli.Run()
}
