package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createWalletRequest struct {
	Name                string `json:"name" binding:"required,min=3"` // Field must be exported and have a JSON tag
	ViewingPrivateKey   string `json:"viewing_priv_key" binding:"required"`
	SpendingPublicKey   string `json:"spending_pub_key" binding:"required"`
	RegisterTransaction string `json:"register_transaction" binding:"required"`
}

func RegisterWalletRoutes(api *gin.RouterGroup) {
	api.GET("/wallet", getWallet)
	api.POST("/wallet", registerWallet)
}

func getWallet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get wallet",
	})
}

func registerWallet(c *gin.Context) {
	var request createWalletRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(request.Name)
	log.Println(request.ViewingPrivateKey)
	log.Println(request.RegisterTransaction)

	c.JSON(200, gin.H{
		"message": "create wallet",
	})
}
