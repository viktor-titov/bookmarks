package configCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get a configuration value by key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Setting config value...")
		},
	}

	return cmd
}
