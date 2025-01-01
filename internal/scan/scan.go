package scan

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/helixpay-xyz/go-helix/pkg/blockchain"
)

type Scanner struct {
	client     *ethclient.Client
	blockchain *blockchain.Blockchain
	workers    *workers
}

func NewScanner() *Scanner {
	client, _ := ethclient.Dial("https://rpc.viction.xyz")
	return &Scanner{
		client:     client,
		blockchain: blockchain.NewBlockchain(client),
		workers:    NewWorker(10, make(chan *types.Transaction)),
	}
}

func (s *Scanner) Run() {
	newHead := make(chan *types.Header)

	go s.blockchain.CrawlBlock(0, 0, newHead)
	go s.workers.Run()

	for header := range newHead {
		go func() {
			block, err := s.client.BlockByNumber(context.Background(), header.Number)

			if err != nil {
				return
			}

			for _, tx := range block.Transactions() {
				s.workers.transactions <- tx
			}
		}()
	}
}
