/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/timreimherr/jhelp/internal/service"
)

// getCmd represents the getSection command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "\"get\" retreives a section by section name (case insensitive)",
	Long:  `\"get\" retreives a section by section name (case insensitive)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Please provide a <section-name>")
		}
		name := args[0]
		service.PrintSection(name)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getSectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getSectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
