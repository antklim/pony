package internal

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Props content properties.
type Props map[string]string

// Meta stores pages metadata.
type Meta struct {
	Pages map[string]Page
}

// Load loads pages metadata from file.
func (m *Meta) Load(file string) error {
	meta, err := loadInMeta(file)
	if err != nil {
		return err
	}

	pages := make(map[string]Page)

	for id, page := range meta.Pages {
		tmpl := meta.Template
		if page.Template != nil {
			tmpl = *page.Template
		}

		pages[id] = Page{
			ID:         id,
			Name:       page.Name,
			Path:       page.Path,
			Template:   tmpl,
			Properties: page.Props(),
		}
	}

	m.Pages = pages

	return nil
}

// Page stores page metadata.
type Page struct {
	ID         string
	Name       string
	Path       string
	Template   string // template name
	Properties map[string]string
}

// OutDir builds page output directory path.
func (p Page) OutDir(outdir string) string {
	ppath := strings.TrimSpace(p.Path)
	subdir := ""
	if ppath != "" && ppath != "/" {
		subdir = ppath
	}

	return filepath.Join(outdir, subdir)
}

// TODO: add fields validation
// All fields are required
type inMeta struct {
	Pages    map[string]inPage `yaml:"pages"`
	Template string            `yaml:"template"`
}

func loadInMeta(file string) (*inMeta, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return parseInMeta(buf)
}

func parseInMeta(data []byte) (*inMeta, error) {
	meta := &inMeta{}
	if err := yaml.Unmarshal(data, meta); err != nil {
		return nil, err
	}

	return meta, nil
}

type inPage struct {
	Name       string       `yaml:"name"`
	Path       string       `yaml:"path"`
	Template   *string      `yaml:"template"`
	Properties []inProperty `yaml:"properties"`
}

func (p *inPage) Props() Props {
	if len(p.Properties) == 0 {
		return nil
	}

	props := make(map[string]string)
	for _, prop := range p.Properties {
		props[prop.Key] = prop.Value
	}
	return props
}

type inProperty struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
