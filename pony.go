package pony

import (
	"html/template"
	"io"

	"github.com/antklim/pony/internal"
)

type options struct {
	templateDir  string // template directory root
	metadataFile string // metadata file name
}

// Option sets pony options such as template directory, schema, etc.
type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (fo *funcOption) apply(o *options) {
	fo.f(o)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{f}
}

// TemplateDir sets template directory root.
func TemplateDir(s string) Option {
	return newFuncOption(func(o *options) {
		o.templateDir = s
	})
}

// MetadataFile sets metadata file name.
func MetadataFile(s string) Option {
	return newFuncOption(func(o *options) {
		o.metadataFile = s
	})
}

// Pony is a static page renderer.
type Pony struct {
	opts options

	tmpl *template.Template
	meta *internal.Meta
}

// NewPony creates an instance of Pony which has no metadata or templates loaded.
func NewPony(opt ...Option) *Pony {
	return nil
}

func (p *Pony) LoadMetadata() error {
	return nil
}

func (p *Pony) LoadTemplates() error {
	return nil
}

func (p *Pony) RenderPage(w io.Writer) error {
	return nil
}
