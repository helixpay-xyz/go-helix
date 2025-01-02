package scan

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/helixpay-xyz/go-helix/pkg/crawler"
)

type Scanner struct {
	clients map[string]*ethclient.Client
	worker  *worker
}

func NewScanner() *Scanner {
	client, _ := ethclient.Dial("https://rpc.viction.xyz")

	clients := make(map[string]*ethclient.Client)
	clients["viction"] = client

	worker := NewWorker(10, make(chan *crawler.TransactionData))

	return &Scanner{
		clients: clients,
		worker:  worker,
	}
}

func (s *Scanner) Run() {
	s.worker.Run()
	for chain, client := range s.clients {
		crawler := crawler.NewCrawler(client, chain, 0, 0, s.worker.transactions)
		go crawler.StartCrawl()
	}
}
