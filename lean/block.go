package lean

import (
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
	txs = append(txs, Transaction{
		Output: "reward payment",
		Input:  "satoshi",
		Value:  50,
	})
	block := &Block{
		Timestamp:    time.Unix(time.Now().Unix(), 0),
		PrevBlock:    prevBlockHash,
		Transactions: txs,
		Hash:         []byte{},
		Bits:         getTargetBits(),
		Nonce:        1,
	}
	return block
}
