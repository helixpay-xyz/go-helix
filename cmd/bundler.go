/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// bundlerCmd represents the bundler command
var bundlerCmd = &cobra.Command{
	Use:   "bundler",
	Short: "Run basic bundler",
	Long:  `Bundler is the process to bundle UserOperation and submit to the blockchain.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bundler called")
	},
}

func init() {
	rootCmd.AddCommand(bundlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bundlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bundlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
