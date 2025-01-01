package scan

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Listener struct {
	client     *ethclient.Client
	headerChan chan *types.Header
}

func NewListener(client *ethclient.Client, headerChain chan *types.Header) *Listener {
	return &Listener{
		client:     client,
		headerChan: headerChain,
	}
}

func (s *Listener) Run() {
	sub, err := s.client.SubscribeNewHead(context.Background(), s.headerChan)

	if err != nil {
		panic(err)
	}
	if err := <-sub.Err(); err != nil {
		panic(err)
	}
}
