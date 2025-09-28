// Команды для работы с конфигурацией
package configCmd

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration settings",
		Long:  `Manage configuration settings for the application.`,
	}

	cmd.AddCommand(newInitCommand())
	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newGetCommand())
	cmd.AddCommand(newSetCommand())

	return cmd
}
