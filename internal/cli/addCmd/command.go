package addCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newAddCommandCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command",
		Short: "Add a new command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Add command")
		},
	}

	return cmd
}
