package pony

import (
	"html/template"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type config struct {
	templatesDir string // templates directory root
	metadataFile string // metadata file name
}

// Option sets pony configuration options such as template directory, schema, etc.
type Option interface {
	apply(*config)
}

type funcOption struct {
	f func(*config)
}

func (fo *funcOption) apply(cfg *config) {
	fo.f(cfg)
}

func newFuncOption(f func(*config)) *funcOption {
	return &funcOption{f}
}

// TemplatesDir sets template directory root.
func TemplatesDir(s string) Option {
	return newFuncOption(func(cfg *config) {
		cfg.templatesDir = s
	})
}

// MetadataFile sets metadata file name.
func MetadataFile(s string) Option {
	return newFuncOption(func(cfg *config) {
		cfg.metadataFile = s
	})
}

// Pony is a static page renderer.
type Pony struct {
	cfg config

	meta *Metadata
	tmpl *template.Template
}

// NewPony creates an instance of Pony which has no metadata or templates loaded.
func NewPony(opts ...Option) *Pony {
	cfg := config{}
	for _, opt := range opts {
		opt.apply(&cfg)
	}

	p := &Pony{
		cfg: cfg,
	}

	return p
}

// LoadMetadata loads metadata file.
func (p *Pony) LoadMetadata() error {
	buf, err := ioutil.ReadFile(p.cfg.metadataFile)
	if err != nil {
		return errors.Wrap(err, "metadata read failed")
	}

	metaInput := &MetadataInput{}
	if err := yaml.Unmarshal(buf, metaInput); err != nil {
		return errors.Wrap(err, "metadata parse failed")
	}

	pages := make(map[string]Page, len(metaInput.Pages))
	for id, pageInput := range metaInput.Pages {
		page := FromPageInput(id, pageInput)
		pages[id] = page
	}

	p.meta = &Metadata{
		Pages: pages,
	}

	return nil
}

// MetadataLoaded returns whether metadata was loaded.
func (p *Pony) MetadataLoaded() bool {
	return p.meta != nil
}

// LoadTemplates loads templates.
func (p *Pony) LoadTemplates() error {
	return nil
}

// TemplatesLoaded returns whether templates were loaded.
func (p *Pony) TemplatesLoaded() bool {
	return p.tmpl != nil
}

// RenderPage renders page by provided metadata and template.
func (p *Pony) RenderPage(w io.Writer) error {
	return nil
}
