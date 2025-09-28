package listCmd

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all links and commands",
		Aliases: []string{"ls", "l"},
	}

	cmd.AddCommand(newListOfCommandsCommand())
	cmd.AddCommand(newListOfLinksCommand())
	cmd.AddCommand(newAllCommand())

	return cmd
}
