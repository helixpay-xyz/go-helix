package client

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/helixpay-xyz/go-helix/mempool"
	"github.com/helixpay-xyz/go-helix/userop"
)

type Client struct {
	chain   string
	rpcUrl  string
	mempool *mempool.Mempool
}

func NewClient(chain string, rpcUrl string, mempool *mempool.Mempool) *Client {
	return &Client{
		chain:   chain,
		rpcUrl:  rpcUrl,
		mempool: mempool,
	}
}

func (c *Client) Eth_sendUserOperations(params []any) (string, error) {
	// Convert each param to UserOperation
	var userOps []userop.UserOperation
	for _, param := range params {
		paramMap, ok := param.(map[string]any)
		if !ok {
			return "", errors.New("invalid param item")
		}

		// Convert the map to JSON bytes
		paramBytes, err := json.Marshal(paramMap)
		if err != nil {
			return "", err
		}

		// Unmarshal into UserOperation
		var userOperation userop.UserOperation
		err = json.Unmarshal(paramBytes, &userOperation)
		if err != nil {
			return "", err
		}

		userOps = append(userOps, userOperation)
	}
	log.Println("User operations: ", userOps)

	for _, userOp := range userOps {
		log.Println("User operation: ", userOp)
	}

	return "", nil
}

func (c *Client) Eth_getUserOperationReceipt(params []any) (string, error) {
	return "", nil
}
