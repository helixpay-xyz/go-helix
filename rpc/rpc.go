package rpc

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helixpay-xyz/go-helix/client"
)

type Rpc struct {
	client  *client.Client
	methods map[string]func(data map[string]any) (string, error)
}

func jsonrpcError(c *gin.Context, code int, message string, data any, id any) {
	c.JSON(http.StatusOK, gin.H{
		"jsonrpc": "2.0",
		"error": gin.H{
			"code":    code,
			"message": message,
			"data":    data,
		},
		"id": id,
	})
	c.Abort()
}

func NewRpc(client *client.Client) *Rpc {
	r := &Rpc{
		client:  client,
		methods: make(map[string]func(params map[string]any) (string, error)),
	}
	r.methods["eth_sendUserOperation"] = client.Eth_sendUserOperation // Register methods here
	return r
}

func (r *Rpc) handleRequest(c *gin.Context, data map[string]any) (id any, result any, success bool) {
	method, ok := data["method"].(string)
	if !ok {
		jsonrpcError(c, -32600, "Invalid Request", "No or invalid 'method' in request", data["id"])
		return data["id"], nil, false
	}

	handler, exists := r.methods[method]
	if !exists {
		jsonrpcError(c, -32601, "Method not found", "The method does not exist / is not available", data["id"])
		return data["id"], nil, false
	}

	params, _ := data["params"].(map[string]any)
	result, err := handler(params)
	if err != nil {
		jsonrpcError(c, -32603, "Internal error", err.Error(), data["id"])
		return data["id"], nil, false
	}

	return data["id"], result, true
}

func (r *Rpc) HandleRequest(c *gin.Context) {
	if c.Request.Method != "POST" {
		jsonrpcError(c, -32700, "Parse error", "POST method expected", nil)
		return
	}
	if c.Request.Body == nil {
		jsonrpcError(c, -32700, "Parse error", "No POST data", nil)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		jsonrpcError(c, -32700, "Parse error", "Error while reading request body", nil)
		return
	}

	data := make(map[string]any)
	if err := json.Unmarshal(body, &data); err != nil {
		jsonrpcError(c, -32700, "Parse error", "Error while parsing request body", nil)
		return
	}

	if id, res, success := r.handleRequest(c, data); success {
		c.JSON(http.StatusOK, gin.H{
			"jsonrpc": "2.0",
			"id":      id,
			"result":  res,
		})
	} else {
		jsonrpcError(c, -32603, "Internal error", res, id)
	}
}
