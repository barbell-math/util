// A generator program that recursively removes all generated code files in the
// current directory. Generated files are defined to match *.gen.go or
// *.gen_test.go.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/barbell-math/util/src/generators/common"
)

func main() {
	pattern := fmt.Sprintf(
		"(\\%s$)|(\\%s$)",
		common.GeneratedSrcFileExt, common.GeneratedTestFileExt,
	)

	err := filepath.Walk(
		".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil // Skip directories
			}

			matched, err := regexp.MatchString(pattern, path)
			if err != nil {
				return err
			}

			if matched {
				common.PrintRunningInfo("Removing: %s", path)
				err := os.Remove(path)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)

	if err != nil {
		panic(err)
	}
}
