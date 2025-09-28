/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/viktor-titov/bookmarks/internal/cli/addCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/configCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/deleteCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/goCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/listCmd"
	"github.com/viktor-titov/bookmarks/internal/cli/runCmd"
)

var rootCmd = &cobra.Command{
	Use:   "bookmarks",
	Short: "The bookmarks manger CLI application",
	Long:  `Утилита для хранения и запуска short кода команда CLI и открытия в браузере ссылок из командной строки`,
}

// Execute запускает корневую команду.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
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
