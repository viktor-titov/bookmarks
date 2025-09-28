package deleteCmd

import (
	"github.com/spf13/cobra"
)

func newDeleteLinkCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link",
		Short: "Delete a link",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for deleting a link goes here
			cmd.Println("Deleting a link...")
		},
	}

	return cmd
}
