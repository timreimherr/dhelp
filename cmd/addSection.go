/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/timreimherr/dhelp/internal/data"
)

// addSectionCmd represents the addSection command
var addSectionCmd = &cobra.Command{
	Use:   "addSection",
	Short: "Adds a section under which info key/value pairs can be added",
	Long: `Adds a section under which info key/value pairs can be added
Example: dhelp addSection "Custom Section"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding section: ", args[0])
		err := data.CreateSection(args[0])
		if err != nil {
			log.Fatal(err)
		}
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
