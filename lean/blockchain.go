package lean

type Blockchain struct {
	blocks       []*Block
	transactions []Transaction
}

func NewBlockchain() *Blockchain {
	genesis := NewBlock([]Transaction{Transaction{"The Genesis Block", "mywallet", 50}}, []byte{})
	return &Blockchain{blocks: []*Block{genesis}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(bc.transactions, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) NewTransaction(tx Transaction) {
	bc.transactions = append(bc.transactions, tx)
}
