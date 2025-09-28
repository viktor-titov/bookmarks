package deleteCmd

import (
	"fmt"

	"github.com/spf13/cobra"
	action "github.com/viktor-titov/bookmarks/internal/action/command"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/command"
	"github.com/viktor-titov/bookmarks/internal/adapter/config"
)

func newDeleteCommandCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command",
		Short: "Delete a command",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				fmt.Println("Usage: command <key>")
				return nil
			}

			key := args[0]

			cfg, err := config.NewConfigRepository()
			if err != nil {
				return fmt.Errorf("create config repository: %w", err)
			}
			repo, err := adapter.NewCommandRepository(*cfg)
			if err != nil {
				return fmt.Errorf("create command repository: %w", err)
			}

			act := action.NewDelete(repo)
			err = act.Do(key)
			if err != nil {
				return fmt.Errorf("delete command: %w", err)
			}

			cmd.Println("Command deleted successfully")

			return nil
		},
	}

	return cmd
}
