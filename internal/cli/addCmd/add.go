package addCmd

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new link or command",
	}

	cmd.AddCommand(newAddLinkCommand())
	cmd.AddCommand(newAddCommandCommand())
	return cmd
}
