package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build static pages",
		Run:   buildHandler,
	}

	addOutdirFlag(cmd.Flags())
	addStrictFlag(cmd.Flags())

	return cmd
}

func buildHandler(cmd *cobra.Command, args []string) {
	fmt.Println("build >>>")
}
