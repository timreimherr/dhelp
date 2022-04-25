/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/timreimherr/dhelp/internal/data"
	"github.com/timreimherr/dhelp/internal/log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dhelp",
	Short: "dhelp prints cli commands for various dev tools and allows users CRUD operations for additional tool commands",
	Long: `dhelp is a CLI library that prints cli commands
for various dev tools and allows users CRUD operations
for additional tool commands.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		sections := data.GetSections()
		fmt.Println()
		for _, s := range sections {
			infos := data.GetInfosBySectionId(s.Id)
			msg := fmt.Sprintf("%d. %s\n", s.Id, s.Name)
			log.Section(msg)
			for _, i := range infos {
				log.Info(i.Key, i.Value)
			}
			fmt.Println()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dhelp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
