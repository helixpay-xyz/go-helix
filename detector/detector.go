package detector

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// detect transactions send with stealth address format in block
type Detector struct {
	client           *ethclient.Client
	workerNum        int
	headerChan       chan *types.Header
	transactionChain chan *types.Transaction
}

func NewDetector(client *ethclient.Client, headerChan chan *types.Header) *Detector {
	return &Detector{
		client:           client,
		workerNum:        5,
		headerChan:       headerChan,
		transactionChain: make(chan *types.Transaction),
	}
}

func (d *Detector) RunWorkers() {
	for i := 0; i < d.workerNum; i++ {
		go func() {
			for tx := range d.transactionChain {
				// if transaction in block is stealth address format, process it
				log.Println(tx.Data())
			}
		}()
	}
}

func (d *Detector) Run() {
	d.RunWorkers()
	for head := range d.headerChan {
		block, err := d.client.BlockByNumber(context.Background(), head.Number)

		if err != nil {
			log.Fatal(err)
			continue
		}

		for _, tx := range block.Transactions() {
			d.transactionChain <- tx
		}

	}
}
