package main

import (
	"fmt"
	"time"
)

type Block struct {
	Timestamp    time.Time
	PrevBlock    []byte
	Hash         []byte
	Transactions []Transaction
	Bits         uint32
	Nonce        uint32
}

func NewBlock(txs []Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Unix(time.Now().Unix(), 0),
		PrevBlock:    prevBlockHash,
		Transactions: txs,
		Hash:         []byte{},
		Bits:         getTargetBits(),
		Nonce:        1,
	}
	fmt.Printf("Prev. hash: %x\n", block.PrevBlock)
	fmt.Printf("Transactions: %v\n", block.Transactions)
	block.PoW()
	fmt.Println()
	return block
}
