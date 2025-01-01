package blockchain

import "github.com/ethereum/go-ethereum/ethclient"

type Blockchain struct {
	client *ethclient.Client
}

func NewBlockchain(client *ethclient.Client) *Blockchain {
	return &Blockchain{
		client: client,
	}
}
