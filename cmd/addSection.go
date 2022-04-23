/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addSectionCmd represents the addSection command
var addSectionCmd = &cobra.Command{
	Use:   "addSection",
	Short: "addSection creates a new tool header",
	Long:  `addSection creates a new tool header where commands specific to that tool will be grouped`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addSection called")
	},
}

func init() {
	rootCmd.AddCommand(addSectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addSectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addSectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
