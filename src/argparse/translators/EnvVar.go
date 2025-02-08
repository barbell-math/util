package translators

import (
	"os"

	"github.com/barbell-math/util/src/customerr"
)

type (
	// Represents a cmd line argument that will be treated as an environment
	// variable. The value of the environment variable will be returned as the
	// value for the argument.
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
