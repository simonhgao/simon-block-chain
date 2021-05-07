package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to simon's block chain")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		cmds := strings.Split(text, " ")
		if len(cmds) <= 0 {
			fmt.Println("your input is wrong, please type help to get information")
		}
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

		switch cmds[0] {
		case "add":
			if len(cmds) < 2 {
				fmt.Println("your input is wrong, please type help to get information")
			}
			cli.addBlock(text[4:])
		case "print":
			cli.printChain()
		case "help":
			cli.printUsage()
		default:
			fmt.Println("your input is wrong, please type help to get information")
			cli.printUsage()
		}

	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.GetHead().PrevBlockHash)
		fmt.Printf("Data: %s\n", block.GetTransactions())
		fmt.Printf("Hash: %x\n", block.GetHead().Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.GetHead().PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add \"block transactions\" : - add a new block")
	fmt.Println("  print : - print all blocks")
}
