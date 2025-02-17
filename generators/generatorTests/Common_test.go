package tests

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestGeneratedCodeMatchesExpected(t *testing.T) {
	cmd := exec.Command("go", "generate", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	test.Nil(cmd.Run(), t)

	generatedFiles, _ := filepath.Glob("./*.gen.go")
	expFiles, _ := filepath.Glob("./exp/*.gen.go.exp")

	slices.Sort[[]string](generatedFiles)
	slices.Sort[[]string](expFiles)
	test.True(
		slices.EqualFunc[[]string](
			generatedFiles, expFiles,
			func(gen string, exp string) bool {
				return fmt.Sprintf("exp/%s.exp", gen) == exp
			},
		),
		t,
	)

	for i, iterGenFile := range generatedFiles {
		iterExpFile := expFiles[i]

		genFileContents, err := os.ReadFile(iterGenFile)
		test.Nil(err, t)
		expFileContents, err := os.ReadFile(iterExpFile)
		test.Nil(err, t)
		test.True(bytes.Equal(genFileContents, expFileContents), t)
	}
}
