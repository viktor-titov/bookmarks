package deleteCmd

import "github.com/spf13/cobra"

func newDeleteCommandCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command",
		Short: "Delete a command",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for deleting a command goes here
			cmd.Println("Deleting a command...")
		},
	}

	return cmd
}
