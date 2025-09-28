package listCmd

import "github.com/spf13/cobra"

func newListOfCommandsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "commands",
		Aliases: []string{"command", "com", "c"},
		Short:   "List all commands",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for listing commands goes here
			// For now, just print a placeholder message
			cmd.Println("Listing all commands...")
		},
	}
	return cmd
}
