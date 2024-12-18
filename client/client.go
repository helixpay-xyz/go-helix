package client

import "log"

type Client struct {
	chain  string
	rpcUrl string
}

func NewClient(chain string, rpcUrl string) *Client {
	return &Client{
		chain:  chain,
		rpcUrl: rpcUrl,
	}
}

func (c *Client) Eth_sendUserOperation(userOp map[string]any) (string, error) {
	log.Println("Handle eth_sendUserOperation")
	return "", nil
}

func (c *Client) Helix_sendUserOperations(userOps map[string]any) (string, error) {
	log.Println("Handle helix_sendUserOperations")
	return "", nil
}
