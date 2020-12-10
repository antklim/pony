package commands

import (
	"fmt"
	"os"

	"github.com/antklim/pony/internal"
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
// TODO: move validation to builder
func buildHandler(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(meta); err != nil {
		return errors.Wrap(err, "metadata file read failed")
	}

	if _, err := os.Stat(outdir); err != nil {
		return errors.Wrap(err, "output directory read failed")
	}

	if _, err := os.Stat(tmpl); err != nil {
		return errors.Wrap(err, "template file read failed")
	}

	builder := internal.Builder{
		Meta:   meta,
		Tmpl:   tmpl,
		OutDir: outdir,
	}

	if err := builder.LoadMeta(); err != nil {
		return err
	}

	if err := builder.LoadTemplate(); err != nil {
		return err
	}

	if err := builder.GeneratePages(); err != nil {
		return err
	}

	// TODO: add benchmark (counter how fast the site was build) and show it in the result
	fmt.Printf("pages are available in %s\n", outdir)

	return nil
}
