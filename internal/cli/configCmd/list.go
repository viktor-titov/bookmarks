package configCmd

import (
	"github.com/spf13/cobra"
	action "github.com/viktor-titov/bookmarks/internal/action/config"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/config"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List configs",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			repository, err := adapter.NewConfigRepository()
			if err != nil {
				return err
			}

			act := action.NewList(repository)

			data := act.Do()
			for k, v := range data {
				cmd.Printf("%s: %s\n", k, v)
			}

			return nil
		},
	}
	return cmd
}
