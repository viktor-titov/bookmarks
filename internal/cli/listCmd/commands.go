package listCmd

import (
	"fmt"

	"github.com/spf13/cobra"
	action "github.com/viktor-titov/bookmarks/internal/action/command"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/command"
	"github.com/viktor-titov/bookmarks/internal/adapter/config"
)

func newListOfCommandsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "commands",
		Aliases: []string{"command", "com", "c"},
		Short:   "List all commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.NewConfigRepository()
			if err != nil {
				return fmt.Errorf("failed to initialize config repository: %w", err)
			}
			repo, err := adapter.NewCommandRepository(*cfg)
			if err != nil {
				return fmt.Errorf("failed to initialize command repository: %w", err)
			}

			act := action.NewList(repo)
			list, err := act.Do()
			if err != nil {
				return fmt.Errorf("failed to list command: %w", err)
			}

			if len(list) == 0 {
				cmd.Println("No command found.")
				return nil
			}

			for key, value := range list {
				fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", key, value)
			}

			return nil
		},
	}
	return cmd
}
