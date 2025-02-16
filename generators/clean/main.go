// A generator program that recursively removes all generated code files in the
// current directory. Generated files are defined to match *.gen.go or
// *.gen_test.go.
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/barbell-math/util/generators/common"
)

func main() {
	dir := "."
	if len(os.Args) == 2 {
		dir = os.Args[1]
	}

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		common.PrintRunningError("The supplied dir does not exist: '%s'", dir)
		os.Exit(1)
	}

	pattern := fmt.Sprintf(
		"(\\%s$)|(\\%s$)",
		common.GeneratedSrcFileExt, common.GeneratedTestFileExt,
	)

	err := filepath.Walk(
		dir,
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
