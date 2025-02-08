package translators

import (
	"os"
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestEnvVar(t *testing.T) {
	e := EnvVar{}

	_, err := e.Translate("__TEST_ENV_VAR")
	test.ContainsError(EnvVarNotSetErr, err, t)

	test.Nil(os.Setenv("__TEST_ENV_VAR", "secret"), t)
	defer func() { test.Nil(os.Unsetenv("__TEST_ENV_VAR"), t) }()
	val, err := e.Translate("__TEST_ENV_VAR")
	test.Nil(err, t)
	test.Eq("secret", val, t)
}
