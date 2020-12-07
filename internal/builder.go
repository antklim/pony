package internal

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Builder defines parameters required to build static pages.
type Builder struct {
	// Path to metadata file or directory.
	Meta string
	// Path to template file or directory.
	Tmpl string
	// Output directory.
	OutDir string
	// Dont build template at loading time, build it on demand.
	LazyLoad bool
	// Strictly validate integrity of meta and template.
	Strict bool
}

// Build builds template based on the data provided in meta file.
func (b *Builder) Build() error {
	buf, err := ioutil.ReadFile(b.Meta)
	if err != nil {
		return err
	}

	content, err := ContentLoad(string(buf))
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(b.Tmpl)
	if err != nil {
		return err
	}

	for key, page := range content.Pages {
		fname := filepath.Join(b.OutDir, key+".html")
		f, err := os.Create(fname)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(f, page.Props()); err != nil {
			// TODO: keep file only when in debug mode
			// os.RemoveAll(f.Name())
			return err
		}
	}

	return nil
}
