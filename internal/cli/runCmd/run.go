package runCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a command or link",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for running a command or link goes here
			fmt.Println("Running...")
		},
	}

	return cmd
}
