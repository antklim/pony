package internal

// Props ...
type Props map[string]string

// Meta ...
type Meta struct {
	Pages []Page `yaml:"pages"`
}

// Props returs all pages properties
func (m *Meta) Props() map[string]Props {
	if len(m.Pages) == 0 {
		return nil
	}

	props := make(map[string]Props)
	for _, page := range m.Pages {
		props[page.Key] = page.Props()
	}

	return props
}

// Page ...
type Page struct {
	Key         string `yaml:"key"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Title       string `yaml:"title"`
	// Template    Template   `yaml:"template"`
	Properties []Property `yaml:"properties"`
}

// Props ...
func (p *Page) Props() Props {
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
