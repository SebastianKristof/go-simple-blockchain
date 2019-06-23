package main

import (
	"encoding/json"
	"fmt"

	"github.com/SebastianKristof/go-simple-blockchain/block"
	"github.com/SebastianKristof/go-simple-blockchain/blockchain"
)

func main() {
	fmt.Println("hello blockchain")

	b := block.New(1, "give me 100 coins")
	bj, _ := json.Marshal(b)
	fmt.Printf("block: %s\n", bj)

	bc := blockchain.New()
	bc.CreateGenesisBlock()

	bc.AddBlock(b)
	bcj, _ := json.MarshalIndent(bc, "", "  ")
	fmt.Printf("Our blockchain: %s\n", bcj)
}
