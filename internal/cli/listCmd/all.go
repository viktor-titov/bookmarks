package listCmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viktor-titov/bookmarks/internal/adapter/config"

	actionLink "github.com/viktor-titov/bookmarks/internal/action/link"
	adapterLink "github.com/viktor-titov/bookmarks/internal/adapter/links"

	actionCommand "github.com/viktor-titov/bookmarks/internal/action/command"
	adapterCommand "github.com/viktor-titov/bookmarks/internal/adapter/command"
)

func newAllCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "all",
		Short:   "List all links and commands",
		Aliases: []string{"a"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.NewConfigRepository()
			if err != nil {
				return fmt.Errorf("failed to initialize config repository: %w", err)
			}
			linkRepo, err := adapterLink.NewLinksRepository(*cfg)
			if err != nil {
				return fmt.Errorf("failed to initialize links repository: %w", err)
			}

			act := actionLink.NewList(linkRepo)
			list, err := act.Do()
			if err != nil {
				return fmt.Errorf("failed to list links: %w", err)
			}

			if len(list) > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "Links:\n")
				for key, value := range list {
					fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", key, value)
				}
			}

			commandRepo, err := adapterCommand.NewCommandRepository(*cfg)
			if err != nil {
				return fmt.Errorf("failed to initialize command repository: %w", err)
			}

			actCommand := actionCommand.NewList(commandRepo)
			list, err = actCommand.Do()
			if err != nil {
				return fmt.Errorf("failed to list command: %w", err)
			}

			if len(list) > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "Commands:\n")
				for key, value := range list {
					fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", key, value)
				}
			}

			return nil
		},
	}

	return cmd
}
