package blockchain

import (
	"errors"

	"github.com/SebastianKristof/go-simple-blockchain/block"
)

// BlockChain is a growing list of blocks.
type BlockChain []block.Block

// New initializes a new empty blockchain.
func New() BlockChain {
	return []block.Block{}
}

func (bc *BlockChain) CreateGenesisBlock() error {
	gb := block.New(0, "The chain is born! Seb gets 50 coins")
	gb.PreviousHash = "000"
	*bc = append(*bc, gb)
	return nil
}

func (bc *BlockChain) AddBlock(b block.Block) error {
	if len(*bc) == 0 {
		return errors.New("cannot add block to an empty blockchain")
	}
	b.Index = (*bc)[len(*bc)-1].Index + 1
	b.PreviousHash = (*bc)[len(*bc)-1].Hash

	*bc = append(*bc, b)

	return nil
}
