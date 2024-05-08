/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// type listOptions struct {
// 	options []string
// }

type Options struct {
	ProjectName string
	ProjectType string
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start up the daetrader CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}

// func init() {
// 	rootCmd.AddCommand(startCmd)
// }
