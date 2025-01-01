package blockchain

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

func (bc *Blockchain) CrawlBlock(from int, to int, newHeader chan<- *types.Header) {
	data, _ := bc.client.BlockNumber(context.Background())

	log.Println(data)
}
