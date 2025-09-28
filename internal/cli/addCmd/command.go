package addCmd

import (
	"fmt"

	"github.com/spf13/cobra"

	action "github.com/viktor-titov/bookmarks/internal/action/command"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/command"
	"github.com/viktor-titov/bookmarks/internal/adapter/config"
)

func newAddCommandCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command",
		Short: "Add a new command",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				fmt.Println("Usage: add command <key> <value>")
				return nil
			}

			key := args[0]
			value := args[1]

			cfg, err := config.NewConfigRepository()
			if err != nil {
				return fmt.Errorf("create config repository: %w", err)
			}
			repo, err := adapter.NewCommandRepository(*cfg)
			if err != nil {
				return fmt.Errorf("create command repository: %w", err)
			}

			act := action.NewSet(repo)
			err = act.Do(key, value)
			if err != nil {
				return fmt.Errorf("add command: %w", err)
			}

			cmd.Printf("Command added: %s -> %s\n", key, value)
			return nil
		},
	}

	return cmd
}
