package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/big"
)

func (b *Block) PoW() {
	var hash []byte
	target := big.NewInt(1)
	target.Lsh(target, uint(256-b.Bits))

	fmt.Println("mining...")
	for b.Nonce < math.MaxUint32 {
		hash = b.calcHash()
		fmt.Printf("\rHash: %x", hash)
		if b.validateHash(hash, target) {
			break
		}
		b.Nonce++
	}
	b.Hash = hash[:]
	fmt.Println("\n")
}

func (b *Block) calcHash() []byte {
	headerHex := new(bytes.Buffer)

	headerHex.Write(b.PrevBlock)
	writeElement(headerHex, b.Transactions)
	writeElement(headerHex, b.Timestamp.Unix())
	writeElement(headerHex, b.Bits)
	writeElement(headerHex, b.Nonce)

	hash := sha256.Sum256(headerHex.Bytes())
	return hash[:]
}

func writeElement(w io.Writer, e interface{}) error {
	err := binary.Write(w, binary.BigEndian, e)
	if err != nil {
		return err
	}
	return nil
}

func (b *Block) validateHash(hash []byte, target *big.Int) bool {
	var hashInt big.Int
	hashInt.SetBytes(hash[:])
	if hashInt.Cmp(target) == -1 {
		return true
	}
	return false
}

func getTargetBits() uint32 {
	return 20
}
