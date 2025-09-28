package listCmd

import "github.com/spf13/cobra"

func newAllCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "all",
		Short:   "List all links and commands",
		Aliases: []string{"a"},
		Run: func(cmd *cobra.Command, args []string) {
			// This will call the existing list command to show all links and commands
			cmd.Println("Listing all links and commands...")
		},
	}

	return cmd
}
