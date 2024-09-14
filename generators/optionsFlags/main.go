package main

import (
	"os"

	"github.com/barbell-math/util/generators/common"
)

type (
	Values struct {
		OptionsStruct string `required:"t" default:"" help:"The struct type that holds the options."`
		OptionsEnum   string `required:"t" default:"" help:"The type that holds the flags."`
	}
)

var (
	VALS Values
)

func main() {
	common.Args(&VALS, os.Args)

	// Look through ast for all const declarations of supplied type in the current package
	// Verify that there are no duplicate flag definitions
	// make file with all flag methods + getflag method
}
