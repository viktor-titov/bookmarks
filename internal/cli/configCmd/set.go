package configCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set a configuration value",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Setting config value...")
		},
	}

	return cmd
}
