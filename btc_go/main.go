package main

import (
	"fmt"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func main() {
	a := &Block{}
  fmt.Printf("%p\n", a)
}
