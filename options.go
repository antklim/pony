package pony

type options struct {
	templateDir  string // template directory root
	metadataFile string // metadata file name
	outputDir    string // output directory root
	addr         string // server address for the site preview or site map (ip:port)
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

// OutputDir sets output directory root.
func OutputDir(s string) Option {
	return newFuncOption(func(o *options) {
		o.outputDir = s
	})
}

// Address sets server address for site preview or site map preview.
func Address(s string) Option {
	return newFuncOption(func(o *options) {
		o.addr = s
	})
}
