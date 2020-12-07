package internal

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// TODO: template is a part of meta and should be loaded as part of MetaLoad.

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

// Build applies metadata to template.
func (b *Builder) Build() error {
	buf, err := ioutil.ReadFile(b.Meta)
	if err != nil {
		return err
	}

	meta, err := MetaLoad(string(buf))
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(b.Tmpl)
	if err != nil {
		return err
	}

	for id, page := range meta.Pages {
		if err := buildPage(id, page, tmpl, b.OutDir); err != nil {
			return err
		}
	}

	return nil
}

func buildPage(id string, page Page, tmpl *template.Template, outdir string) error {
	ppath := strings.TrimSpace(page.Path)

	subdir := ""
	if ppath != "" && ppath != "/" {
		subdir = ppath
	}

	dir := filepath.Join(outdir, subdir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			return err
		}
	}

	fname := filepath.Join(outdir, subdir, id+".html")
	f, err := os.Create(fname)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, page.Props()); err != nil {
		// TODO: keep file only when in debug mode
		// os.RemoveAll(f.Name())
		return err
	}

	return nil
}
