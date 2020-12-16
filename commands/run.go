package commands

import (
	"github.com/antklim/pony"
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

// TODO: add address and target flags

func runHandler(cmd *cobra.Command, args []string) error {
	s := &pony.Server{
		Addr:         ":9000",
		Target:       pony.SiteMapViewServer,
		MetadataFile: meta,
		TemplatesDir: tmpl,
	}

	if err := s.Start(); err != nil {
		return err
	}

	return nil
}
