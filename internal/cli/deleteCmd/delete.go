package deleteCmd

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a link or command",
	}

	cmd.AddCommand(newDeleteLinkCommand())
	cmd.AddCommand(newDeleteCommandCommand())

	return cmd
}
