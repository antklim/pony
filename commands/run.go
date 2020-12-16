package commands

import (
	"github.com/antklim/pony"
	"github.com/spf13/cobra"
)

var ttarget = pony.SitePreviewServer

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run server to preview pages",
		PreRun: func(cmd *cobra.Command, args []string) {
			if target == "map" {
				ttarget = pony.SiteMapViewServer
			}
		},
		RunE: runHandler,
	}

	addAddressFlag(cmd.Flags())
	addTargetFlag(cmd.Flags())

	return cmd
}

func runHandler(cmd *cobra.Command, args []string) error {
	s := &pony.Server{
		Addr:         address,
		Target:       ttarget,
		MetadataFile: meta,
		TemplatesDir: tmpl,
	}

	if err := s.Start(); err != nil {
		return err
	}

	return nil
}
