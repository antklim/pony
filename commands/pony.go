package commands

import (
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(args []string) error {
	cmd := newPonyCmd()
	cmd.SetArgs(args)
	return cmd.Execute()
}

func newPonyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pony",
		Short: "Pony is a simple static site generator and viewer",
	}

	cmd.AddCommand(newBuildCmd(), newRunCmd())

	addMetaFlag(cmd.PersistentFlags())
	addTemplateFlag(cmd.PersistentFlags())
	addStrictFlag(cmd.Flags())

	return cmd
}
