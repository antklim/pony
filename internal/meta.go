package internal

import (
	"gopkg.in/yaml.v2"
)

// Props content properties.
type Props map[string]string

// Meta describes content metadata.
type Meta struct {
	Pages map[string]Page `yaml:"pages"`
}

// MetaLoad creates instance of content metadata from source.
func MetaLoad(data string) (*Meta, error) {
	meta := &Meta{}
	if err := yaml.Unmarshal([]byte(data), meta); err != nil {
		return nil, err
	}

	return meta, nil
}

// Page describes content page structure.
type Page struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	// Template    Template   `yaml:"template"`
	Properties []Property `yaml:"properties"`
}

// Props returns content page properties.
func (p Page) Props() Props {
	if len(p.Properties) == 0 {
		return nil
	}

	props := make(map[string]string)
	for _, prop := range p.Properties {
		props[prop.Key] = prop.Value
	}
	return props
}

// Template ...
// type Template struct {
// 	FilePath string
// 	Tmpl     *template.Template
// }

// Property content property.
type Property struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
