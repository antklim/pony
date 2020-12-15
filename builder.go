package pony

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Builder struct {
	MetadataFile string
	OutDir       string
	TemplatesDir string
	Pony         *Pony
}

func (b *Builder) Build() error {
	if _, err := os.Stat(b.MetadataFile); err != nil {
		return errors.Wrap(err, "metadata file read failed")
	}

	if _, err := os.Stat(b.OutDir); err != nil {
		return errors.Wrap(err, "output directory read failed")
	}

	if _, err := os.Stat(b.TemplatesDir); err != nil {
		return errors.Wrap(err, "templates directory read failed")
	}

	return b.Pony.RenderPages(fileWriter(b.OutDir))
}

func fileWriter(dir string) PageWriter {
	return func(page Page) (io.Writer, error) {
		outDir := filepath.Join(dir, page.Path)
		if _, err := os.Stat(outDir); os.IsNotExist(err) {
			if err := os.Mkdir(outDir, 0755); err != nil {
				return nil, err
			}
		}

		fname := filepath.Join(outDir, "index.html")
		return os.Create(fname)
	}
}
