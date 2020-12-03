package internal

// Meta ...
type Meta struct {
	Pages []Page `yaml:"pages"`
}

// Page ...
type Page struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Title       string `yaml:"title"`
	// Template    Template   `yaml:"template"`
	Properties []Property `yaml:"properties"`
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
