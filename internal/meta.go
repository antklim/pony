package internal

import (
	"gopkg.in/yaml.v2"
)

// Props content properties.
type Props map[string]string

// TODO: add fields validation

// Meta describes content metadata.
// All fields are required
type Meta struct {
	Pages    map[string]Page `yaml:"pages"`
	Template string          `yaml:"template"`
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
	Name       string     `yaml:"name"`
	Path       string     `yaml:"path"`
	Template   *string    `yaml:"template"`
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

// Property content property.
type Property struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
