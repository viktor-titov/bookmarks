package listCmd

import (
	"fmt"

	"github.com/spf13/cobra"
	action "github.com/viktor-titov/bookmarks/internal/action/link"
	adapterConfig "github.com/viktor-titov/bookmarks/internal/adapter/config"
	adapterLink "github.com/viktor-titov/bookmarks/internal/adapter/links"
)

func newListOfLinksCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "links",
		Short:   "List all links",
		Aliases: []string{"link", "lin", "l"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := adapterConfig.NewConfigRepository()
			if err != nil {
				return fmt.Errorf("failed to initialize config repository: %w", err)
			}
			linkRepo, err := adapterLink.NewLinksRepository(*cfg)
			if err != nil {
				return fmt.Errorf("failed to initialize links repository: %w", err)
			}

			act := action.NewList(linkRepo)
			list, err := act.Do()
			if err != nil {
				return fmt.Errorf("failed to list links: %w", err)
			}

			if len(list) == 0 {
				cmd.Println("No links found.")
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
