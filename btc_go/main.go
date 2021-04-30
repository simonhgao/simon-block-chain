package main

import "fmt"

func main() {
	println("Let's start a demo block chain example")

	bc := NewBlockChain()
	bc.AddBlock("Send 1 BTC to simon")
	bc.AddBlock("Send 2 more BTC to kevin")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. Hash: %x\n", block.GetHead().PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Transactions)
		fmt.Printf("Hash: %x\n", block.GetHead().Hash)
		fmt.Println()
	}
}
