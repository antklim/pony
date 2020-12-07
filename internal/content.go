package internal

import (
	"gopkg.in/yaml.v2"
)

// Props ...
type Props map[string]string

// Content ...
type Content struct {
	Pages map[string]Page `yaml:"pages"`
}

// ContentLoad creates instance of content from source.
func ContentLoad(data string) (*Content, error) {
	content := &Content{}
	if err := yaml.Unmarshal([]byte(data), content); err != nil {
		return nil, err
	}

	return content, nil
}

// Page ...
type Page struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	// Template    Template   `yaml:"template"`
	Properties []Property `yaml:"properties"`
}

// Props ...
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

// Property ...
type Property struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
