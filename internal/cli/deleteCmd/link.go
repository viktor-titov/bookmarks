package deleteCmd

import (
	"fmt"

	"github.com/spf13/cobra"

	action "github.com/viktor-titov/bookmarks/internal/action/link"
	"github.com/viktor-titov/bookmarks/internal/adapter/config"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/links"
)

func newDeleteLinkCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link",
		Short: "Delete a link",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				fmt.Println("Usage: link <key>")
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

			act := action.NewDelete(repo)
			err = act.Do(key)
			if err != nil {
				return fmt.Errorf("delete link: %w", err)
			}

			cmd.Println("Link deleted successfully")

			return nil
		},
	}

	return cmd
}
