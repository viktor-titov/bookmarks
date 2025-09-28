package configCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List configs",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing configs...")
		},
	}
	return cmd
}
