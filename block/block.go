package block

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block describes one block of the chain.
type Block struct {
	Index        int64     `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	Hash         string    `json:"hash"`
	PreviousHash string    `json:"previousHash"`
	Data         string    `json:"data"`
}

func New(index int64, data string) Block {

	b := Block{
		Timestamp:    time.Now(),
		Index:        index,
		Data:         data,
		PreviousHash: "",
		Hash:         "",
	}

	b.calculateHash()

	return b
}

// func getPreviousBlock() block {
// 	return block{}
// }

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(
		fmt.Sprint(b.Index, b.Timestamp, b.PreviousHash, b.Data)))

	b.Hash = fmt.Sprintf("%x", hash)
}
