package configCmd

import (
	"fmt"

	"github.com/spf13/cobra"

	action "github.com/viktor-titov/bookmarks/internal/action/config"
	adapter "github.com/viktor-titov/bookmarks/internal/adapter/config"
)

func newSetCommand() *cobra.Command {
	var key string
	var value string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set a configuration value",
		RunE: func(cmd *cobra.Command, args []string) error {
			if key == "" || value == "" {
				cmd.Println("both --key and --value flags are required")
				return nil
			}

			repo, err := adapter.NewConfigRepository()
			if err != nil {
				return err
			}

			action := action.NewSet(repo)
			err = action.Do(key, value)
			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Configuration key '%s' set to '%s'\n", key, value)

			return nil
		},
	}

	cmd.Flags().StringVarP(&key, "key", "k", "", "Key name")
	cmd.Flags().StringVarP(&value, "value", "v", "", "Key value")

	return cmd
}
