package internal

import (
	"html/template"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
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
	// meta, err := readMeta(b.Meta)
	// if err != nil {
	// 	return err
	// }

	// tmpl, err := template.ParseFiles(b.Tmpl)
	// if err != nil {
	// 	return err
	// }

	// for _, page := range meta.Pages {
	// 	fname := filepath.Join(b.OutDir, page.Key+".html")
	// 	f, err := os.Create(fname)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if err := buildPage(f, tmpl, page); err != nil {
	// 		// TODO: keep file only when in debug mode
	// 		// os.RemoveAll(f.Name())
	// 		return err
	// 	}
	// }

	return nil
}

func buildPage(w io.Writer, tmpl *template.Template, page Page) error {
	props := page.Props()
	return tmpl.Execute(w, props)
}

func readMeta(filename string) (*Meta, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	meta := &Meta{}
	if err := yaml.Unmarshal(buf, meta); err != nil {
		return nil, err
	}

	return meta, nil
}
