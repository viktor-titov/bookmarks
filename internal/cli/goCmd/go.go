package goCmd

import (
	"fmt"

	"github.com/spf13/cobra"
	action "github.com/viktor-titov/bookmarks/internal/action/link"
	"github.com/viktor-titov/bookmarks/internal/adapter/config"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/links"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go",
		Short: "Go to a link",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				fmt.Println("Usage: go <key>")
				return nil
			}

			key := args[0]

			cfg, err := config.NewConfigRepository()
			if err != nil {
				return fmt.Errorf("create config repository: %w", err)
			}
			repo, err := adapter.NewLinksRepository(*cfg)
			if err != nil {
				return fmt.Errorf("create links repository: %w", err)
			}

			act := action.NewGo(repo, cfg)
			err = act.Do(key)
			if err != nil {
				return fmt.Errorf("go to link: %w", err)
			}

			return nil
		},
	}

	return cmd
}
