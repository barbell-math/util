package translators

import (
	"os"

	"github.com/barbell-math/util/src/customerr"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=EnvVar

type (
	// Represents a cmd line argument that will be treated as an environment
	// variable. The value of the environment variable will be returned as the
	// value for the argument.
	//gen:ifaceImplCheck ifaceName Translator[string]
	//gen:ifaceImplCheck valOrPntr both
	EnvVar struct{}
)

func (_ EnvVar) Translate(arg string) (string, error) {
	rv, ok := os.LookupEnv(arg)
	if !ok {
		return "", customerr.Wrap(
			EnvVarNotSetErr,
			"The env var '%s' does not exist but was expected to", arg,
		)
	}
	return rv, nil
}

func (_ EnvVar) Reset() {
	// intentional noop - EnvVar has no state
}
