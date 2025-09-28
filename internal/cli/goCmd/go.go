package goCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go",
		Short: "Go to a link",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Going to link...")
		},
	}

	return cmd
}
