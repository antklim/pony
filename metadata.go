package pony

// TODO: extend property value to interface{}

// Properties content properties.
type Properties map[string]string

// FromPropertyInput creates properties map from the list of property inputs.
func FromPropertyInput(input []PropertyInput) Properties {
	props := make(map[string]string, len(input))
	for _, pi := range input {
		props[pi.Key] = pi.Value
	}
	return Properties(props)
}

// Metadata stores pages metadata.
type Metadata struct {
	Pages    map[string]Page
	Template string
}

// Page stores page metadata.
type Page struct {
	ID         string
	Name       string
	Path       string
	Template   *string
	Properties Properties
}

// FromPageInput creates a page from page input structure.
func FromPageInput(id string, pi PageInput) Page {
	return Page{
		ID:         id,
		Name:       pi.Name,
		Path:       pi.Path,
		Template:   pi.Template,
		Properties: FromPropertyInput(pi.Properties),
	}
}

// TODO: add validation for *Input structures

// MetadataInput describes input format of metadata.
type MetadataInput struct {
	Pages    map[string]PageInput `yaml:"pages"`
	Template string               `yaml:"template"`
}

// PageInput describes input format of page metadata.
type PageInput struct {
	Name       string          `yaml:"name"`
	Path       string          `yaml:"path"`
	Template   *string         `yaml:"template"`
	Properties []PropertyInput `yaml:"properties"`
}

// PropertyInput describes property input format.
type PropertyInput struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
