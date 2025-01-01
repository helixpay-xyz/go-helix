package scan

import (
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

type workers struct {
	numWorker    int
	transactions chan *types.Transaction
}

func NewWorker(numWorker int, transactions chan *types.Transaction) *workers {
	return &workers{
		numWorker:    numWorker,
		transactions: transactions,
	}
}

func (w *workers) Run() {
	for i := 0; i < w.numWorker; i++ {
		go func() {
			for tx := range w.transactions {
				log.Println(tx.Hash().Hex())
			}
		}()
	}
}
