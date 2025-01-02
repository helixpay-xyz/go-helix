package scan

import (
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/helixpay-xyz/go-helix/pkg/crawler"
)

type Work struct {
	transaction *types.Transaction
	chain       string
}

type worker struct {
	numWorker    int
	transactions chan *crawler.TransactionData
}

func NewWorker(numWorker int, transactions chan *crawler.TransactionData) *worker {
	return &worker{
		numWorker:    numWorker,
		transactions: transactions,
	}
}

func (w *worker) Run() {
	for i := 0; i < w.numWorker; i++ {
		go func() {
			for tx := range w.transactions {
				log.Println("Receive transaction from chain", tx.Chain, "with hash", tx.Transaction.Hash().Hex())
			}
		}()
	}
}
