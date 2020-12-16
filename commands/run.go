package commands

import (
	"github.com/spf13/cobra"
)

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run server to preview pages",
		RunE:  runHandler,
	}

	addStrictFlag(cmd.Flags())

	return cmd
}

func runHandler(cmd *cobra.Command, args []string) error {
	return nil
}
