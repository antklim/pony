package commands

import (
	"fmt"
	"log"

	"github.com/antklim/pony"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build static pages",
		RunE:  buildHandler,
	}

	addOutdirFlag(cmd.Flags())
	addStrictFlag(cmd.Flags())

	return cmd
}

// TODO: return a list of errors one time
func buildHandler(cmd *cobra.Command, args []string) error {
	opts := []pony.Option{
		pony.MetadataFile(meta),
		pony.TemplatesDir(tmpl),
	}
	p := pony.NewPony(opts...)
	if errs := p.LoadAll(); errs != nil {
		log.Println(errs)
		return errors.New("failed to load pony")
	}
	b := &pony.Builder{
		MetadataFile: meta,
		OutDir:       outDir,
		TemplatesDir: tmpl,
		Pony:         p,
	}

	if err := b.Build(); err != nil {
		return err
	}

	// TODO: add benchmark (counter how fast the site was build) and show it in the result
	fmt.Printf("pages are available in %s\n", outDir)

	return nil
}
