package configCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a configuration value by key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Deleting config value...")
		},
	}

	return cmd
}
