/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timreimherr/jhelp/internal/data"
)

// addInfoCmd represents the addInfo command
var addInfoCmd = &cobra.Command{
	Use:   "addInfo",
	Short: `Adds a key/value info pair under a section`,
	Long: `Adds a key/value info pair under a section.
Example: jhelp addInfo <section-id> <info-action> <info-command>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("\nsection-name, info-action and info-command args are required")
			fmt.Println("please run: 'jhelp addInfo -h' for details")
			return
		}
		section := args[0]
		key := args[1]
		value := args[2]

		if section == "" {
			fmt.Println("\nsection-name arg is required")
			fmt.Println("please run: 'jhelp addInfo -h' for details")
			return
		}
		if key == "" {
			fmt.Println("\ninfo-action arg is required")
			fmt.Println("please run: 'jhelp addInfo -h' for details")
			return
		}
		if value == "" {
			fmt.Println("\ninfo-command arg is required")
			fmt.Println("please run: 'jhelp addInfo -h' for details")
			return
		}

		sections := data.GetSectionByName(section)
		if len(sections) == 0 {
			fmt.Println("\narg 1 section-name doesn not exist, please check value for argument or create section")
			return
		}

		sectionName, err := data.AddInfoToSection(sections[0].Id, args[1], args[2])
		if err != nil {
			fmt.Printf("\n%s", err)
			return
		}

		fmt.Printf("\nInfo %s : %s, added to section %s", key, value, sectionName)
	},
}

func init() {
	rootCmd.AddCommand(addInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
