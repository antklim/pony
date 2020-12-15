package commands

import (
	"github.com/spf13/pflag"
)

var (
	meta   = ""    // path to the metadata file
	outDir = ""    // path to the output directory
	strict = false // enable metadata and template match validation [build|run]

	// TODO: validate tmpl to be the path to drectory
	tmpl = "" // path to the template directory
)

func addMetaFlag(flags *pflag.FlagSet) {
	flags.StringVarP(&meta, "metadata", "m", meta, "path to the metadata file")
}

func addOutdirFlag(flags *pflag.FlagSet) {
	flags.StringVarP(&outDir, "outdir", "o", outDir, "path to the output directory")
}

func addStrictFlag(flags *pflag.FlagSet) {
	flags.BoolVarP(&strict, "strict", "s", strict, "enable metadata and template match validation")
}

func addTemplateFlag(flags *pflag.FlagSet) {
	flags.StringVarP(&tmpl, "template", "t", tmpl, "path to the template directory")
}
