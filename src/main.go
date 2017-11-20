package main

func main() {
	bc := NewBlockchain()

	// Adding a new transaction
	transaction := Transaction{"satoshi", "raul", 30}
	bc.NewTransaction(transaction)
	bc.MineBlock()
}
