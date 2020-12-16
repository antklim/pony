package commands

import (
	"fmt"
	"time"

	"github.com/antklim/pony"
	"github.com/spf13/cobra"
)

func newBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build static pages",
		RunE:  buildHandler,
	}

	addOutdirFlag(cmd.Flags())

	return cmd
}

func buildHandler(cmd *cobra.Command, args []string) error {
	b := &pony.Builder{
		MetadataFile: meta,
		OutDir:       outDir,
		TemplatesDir: tmpl,
	}

	start := time.Now()

	if err := b.Build(); err != nil {
		return err
	}

	fmt.Printf("pages are available in %s (built in %s)\n", outDir, time.Since(start).Round(time.Millisecond))

	return nil
}
