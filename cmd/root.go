/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/viktor-titov/bookmarks/internal/cli/addCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/configCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/deleteCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/goCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/listCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/runCmd"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bookmarks",
	Short: "The bookmarks manger CLI application",
	Long:  `Утилита для хранения и запуска short кода команда CLI и открытия в браузере ссылок из командной строки`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.AddCommand(configCmd.NewCommand())
	rootCmd.AddCommand(deleteCmd.NewCommand())
	rootCmd.AddCommand(addCmd.NewCommand())
	rootCmd.AddCommand(listCmd.NewCommand())
	rootCmd.AddCommand(runCmd.NewCommand())
	rootCmd.AddCommand(goCmd.NewCommand())
}
