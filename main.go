package main

import (
	"fmt"

	"simple-chain/chain"
)

func main() {

	// create a new blockchain instance with a mining difficulty of 5
	blockchain := chain.CreateBlockchain(5)

	// record transactions on the blockchain
	blockchain.AddBlock("Carlos", "Reiz", 10)
	blockchain.AddBlock("Leo", "Juli", 5)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.IsValid())
}
