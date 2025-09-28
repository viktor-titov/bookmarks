package configCmd

import (
	"fmt"

	"github.com/spf13/cobra"
	action "github.com/viktor-titov/bookmarks/internal/action/config"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/config"
)

func newInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "set default configuration value",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo, err := adapter.NewConfigRepository()
			if err != nil {
				return err
			}

			action := action.NewInit(repo)
			err = action.Do()
			if err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Configuration initialized with default values")

			return nil
		},
	}

	return cmd
}
