package internal

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// TODO: template is a part of meta and should be loaded as part of MetaLoad.
// TODO: allow builder accept template (to be used in runner)
// TODO: allow builder to output to io writer (to be used in runner)
// TODO: validate that all templates referenced in meta available in template directory

// Builder defines parameters required to build static pages.
type Builder struct {
	// Path to metadata file or directory.
	Meta string
	// Path to templates directory.
	Tmpl string
	// Output directory.
	OutDir string
	// Strictly validate integrity of meta and template.
	Strict bool

	// loaded template
	tmpl *template.Template

	// loaded metadata
	meta *Meta
}

// LoadTemplate loads builder template.
func (b *Builder) LoadTemplate() error {
	tmpl, err := template.ParseGlob(b.Tmpl + "/*.html")
	if err != nil {
		return err
	}

	b.tmpl = tmpl

	return nil
}

// LoadMeta loads pages metadata.
func (b *Builder) LoadMeta() error {
	meta := &Meta{}
	if err := meta.Load(b.Meta); err != nil {
		return err
	}

	b.meta = meta

	return nil
}

// GeneratePages builds pages and saves them to output directory.
func (b *Builder) GeneratePages() error {
	for id, page := range b.meta.Pages {
		outdir := page.OutDir(b.OutDir)
		if _, err := os.Stat(outdir); os.IsNotExist(err) {
			if err := os.Mkdir(outdir, 0755); err != nil {
				return err
			}
		}

		fname := filepath.Join(outdir, id+".html")
		f, err := os.Create(fname)
		if err != nil {
			return err
		}

		if err := b.BuildPage(id, f); err != nil {
			return err
		}
	}

	return nil
}

// BuildPage builds page and writes result to provided writer.
func (b *Builder) BuildPage(id string, w io.Writer) error {
	page, ok := b.meta.Pages[id]
	if !ok {
		return fmt.Errorf("page %s not found in provided configuration", id)
	}

	if err := b.tmpl.ExecuteTemplate(w, page.Template, page.Properties); err != nil {
		return err
	}

	return nil
}
