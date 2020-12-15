package commands

import (
	"fmt"
	"log"
	"os"

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
	if _, err := os.Stat(meta); err != nil {
		return errors.Wrap(err, "metadata file read failed")
	}

	if _, err := os.Stat(outDir); err != nil {
		return errors.Wrap(err, "output directory read failed")
	}

	if _, err := os.Stat(tmpl); err != nil {
		return errors.Wrap(err, "templates directory read failed")
	}

	opts := []pony.Option{
		pony.MetadataFile(meta),
		pony.TemplatesDir(tmpl),
	}
	p := pony.NewPony(opts...)
	if errs := p.LoadAll(); errs != nil {
		log.Println(errs)
		return errors.New("failed to load pony")
	}

	if err := pony.RenderAndStore(p, outDir); err != nil {
		return errors.Wrap(err, "failed to render site")
	}

	// TODO: add benchmark (counter how fast the site was build) and show it in the result
	fmt.Printf("pages are available in %s\n", outDir)

	return nil
}
