package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type API struct {
	gin *gin.Engine
}

func NewAPI() *API {
	return &API{
		gin: gin.Default(),
	}
}

func (a *API) Run() {
	a.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	a.gin.Run(":8080") // Start server on port 8080

	log.Println("API server started at :8080")
}
