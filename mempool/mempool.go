package mempool

import "github.com/dgraph-io/badger/v4"

type Mempool struct {
	db *badger.DB
}

func NewMempool(db *badger.DB) *Mempool {
	return &Mempool{db: db}
}
