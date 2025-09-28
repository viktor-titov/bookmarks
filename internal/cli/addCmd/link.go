package addCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newAddLinkCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link",
		Short: "Add a new link",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for adding a new link goes here
			fmt.Println("Adding a new link...")
		},
	}

	return cmd
}
