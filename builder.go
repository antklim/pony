package pony

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Builder builds and stores site artifacts.
type Builder struct {
	MetadataFile string
	OutDir       string
	TemplatesDir string
	pony         *Pony
}

// Build builds site and saves it to a provided output directory.
func (b *Builder) Build() error {
	if err := b.validate(); err != nil {
		return err
	}

	if err := b.init(); err != nil {
		return err
	}

	return b.pony.RenderPages(fileWriter(b.OutDir))
}

func (b *Builder) validate() error {
	errs := make([]string, 0)

	if _, err := os.Stat(b.MetadataFile); err != nil {
		errs = append(errs, errors.WithMessage(err, "metadata file read failed").Error())
	}

	if _, err := os.Stat(b.OutDir); err != nil {
		errs = append(errs, errors.WithMessage(err, "output directory read failed").Error())
	}

	if _, err := os.Stat(b.TemplatesDir); err != nil {
		errs = append(errs, errors.WithMessage(err, "templates directory read failed").Error())
	}

	if len(errs) == 0 {
		return nil
	}

	emsg := strings.Join(errs, "; ")
	return errors.New(emsg)
}

func (b *Builder) init() error {
	opts := []Option{
		MetadataFile(b.MetadataFile),
		TemplatesDir(b.TemplatesDir),
	}
	p := NewPony(opts...)
	if errs := p.LoadAll(); errs != nil {
		log.Println(errs)
		return errors.New("failed to initialize pony")
	}

	b.pony = p

	return nil
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
