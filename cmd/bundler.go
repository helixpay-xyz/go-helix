/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helixpay-xyz/go-helix/client"
	"github.com/helixpay-xyz/go-helix/config"
	"github.com/helixpay-xyz/go-helix/mempool"
	"github.com/helixpay-xyz/go-helix/rpc"
	"github.com/spf13/cobra"
)

// bundlerCmd represents the bundler command
var bundlerCmd = &cobra.Command{
	Use:   "bundler",
	Short: "Run basic bundler",
	Long:  `Bundler is the process to bundle UserOperation and submit to the blockchain.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.InitConfig()
		if err != nil {
			panic(err)
		}

		// Start RPC
		r := gin.Default()

		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		for _, chain := range config.ActiveChains {
			log.Print("Active chain: ", config.ChainConfigs[chain].Url)
			mempool := mempool.NewMempool()
			client := client.NewClient(chain, config.ChainConfigs[chain].Url, mempool)
			rpc := rpc.NewRpc(client)
			r.POST("/"+chain, rpc.HandleRequest)
		}

		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	},
}

func init() {
	rootCmd.AddCommand(bundlerCmd)
}
