package commands

import (
	"github.com/spf13/pflag"
)

var (
	// flags used in build|run commands
	meta   = ""    // path to the metadata file
	tmpl   = ""    // path to the template directory
	strict = false // enable metadata and template match validation

	// flags used in build command only
	outDir = "" // path to the output directory [build]

	// flags used in run command only
	address = ":9000"   // listen address for the server (ip:port)
	target  = "preview" // server target [map|preview]
)

func addMetaFlag(flags *pflag.FlagSet) {
	flags.StringVarP(&meta, "metadata", "m", meta, "path to the metadata file")
}

func addTemplateFlag(flags *pflag.FlagSet) {
	flags.StringVarP(&tmpl, "template", "t", tmpl, "path to the template directory")
}

func addStrictFlag(flags *pflag.FlagSet) {
	flags.BoolVarP(&strict, "strict", "s", strict, "enable metadata and template match validation")
}

func addOutdirFlag(flags *pflag.FlagSet) {
	flags.StringVarP(&outDir, "outdir", "o", outDir, "path to the output directory")
}

func addAddressFlag(flags *pflag.FlagSet) {
	flags.StringVarP(&address, "address", "a", address, "listen address for the server (ip:port)")
}

func addTargetFlag(flags *pflag.FlagSet) {
	flags.StringVarP(&target, "target", "T", target, "server target [map|preview]")
}
