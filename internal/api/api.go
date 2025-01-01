package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	gin *gin.Engine
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{
		gin: gin.Default(),
	}
}

func (a *APIHandler) Run() {
	a.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	api := a.gin.Group("/api")
	go RegisterWalletRoutes(api)

	a.gin.Run(":8080") // Start server on port 8080

	log.Println("API server started at :8080")
}
