/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/helixpay-xyz/go-helix/detector"
	"github.com/helixpay-xyz/go-helix/listener"
	"github.com/spf13/cobra"
)

// scannerCmd represents the scanner command
var scannerCmd = &cobra.Command{
	Use:   "scanner",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		newHeadChain := make(chan *types.Header)
		wsClient, err := ethclient.Dial("wss://ws.viction.xyz")
		if err != nil {
			panic(err)
		}
		client, err := ethclient.Dial("https://rpc.viction.xyz")
		if err != nil {
			panic(err)
		}

		listener := listener.NewListener(wsClient, newHeadChain)
		detector := detector.NewDetector(client, newHeadChain)

		go detector.Run()
		listener.Run()
	},
}

func init() {
	rootCmd.AddCommand(scannerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scannerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scannerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
