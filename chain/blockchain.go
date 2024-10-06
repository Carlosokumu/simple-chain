package chain

import (
	"strings"
	"time"
)



type BlockChain struct {
	genesisBlock Block
	chain        []Block //defines the minimum effort miners must undertake to mine and include a block in the blockchain.
	difficulty   int
}



func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
			b.pow++
			b.hash = b.calculateHash()
	}
}


func CreateBlockchain(difficulty int) BlockChain {
	genesisBlock := Block{
			hash:      "0",
			timestamp: time.Now(),
	}
	return BlockChain{
			genesisBlock,
			[]Block{genesisBlock},
			difficulty,
	}
}


func (b *BlockChain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
			"from":   from,
			"to":     to,
			"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
			data:         blockData,
			previousHash: lastBlock.hash,
			timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}


func (b BlockChain) IsValid() bool {
	for i := range b.chain[1:] {
			previousBlock := b.chain[i]
			currentBlock := b.chain[i+1]
			if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
					return false
			}
	}
	return true
}