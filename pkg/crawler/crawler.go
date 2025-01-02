package crawler

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionData struct {
	Chain       string
	Transaction *types.Transaction
}

type Crawler struct {
	client   *ethclient.Client
	chain    string
	from     int
	to       int
	resChain chan<- *TransactionData
}

func NewCrawler(client *ethclient.Client, chain string, from int, to int, resChain chan<- *TransactionData) *Crawler {
	return &Crawler{
		client:   client,
		chain:    chain,
		from:     from,
		to:       to,
		resChain: resChain,
	}
}

func (c *Crawler) StartCrawl() {
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				data, _ := c.client.BlockNumber(context.Background())
				block, _ := c.client.BlockByNumber(context.Background(), big.NewInt(int64(data)))

				for _, tx := range block.Transactions() {
					c.resChain <- &TransactionData{
						Chain:       c.chain,
						Transaction: tx,
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
