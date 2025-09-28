package listCmd

import "github.com/spf13/cobra"

func newListOfLinksCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "links",
		Short:   "List all links",
		Aliases: []string{"link", "lin", "l"},
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for listing links goes here
			cmd.Println("Listing all links...")
		},
	}

	return cmd
}
