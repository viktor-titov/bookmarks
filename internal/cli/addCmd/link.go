package addCmd

import (
	"fmt"

	"github.com/spf13/cobra"
	action "github.com/viktor-titov/bookmarks/internal/action/link"
	"github.com/viktor-titov/bookmarks/internal/adapter/config"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/links"
)

func newAddLinkCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link",
		Short: "Add a new link",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				fmt.Println("Usage: add link <key> <value>")
				return nil
			}

			key := args[0]
			value := args[1]

			cfg, err := config.NewConfigRepository()
			if err != nil {
				return fmt.Errorf("create config repository: %w", err)
			}
			repo, err := adapter.NewLinksRepository(*cfg)
			if err != nil {
				return fmt.Errorf("create links repository: %w", err)
			}

			act := action.NewSet(repo)
			err = act.Do(key, value)
			if err != nil {
				return fmt.Errorf("add link: %w", err)
			}

			cmd.Printf("Link added: %s -> %s\n", key, value)
			return nil
		},
	}

	return cmd
}
