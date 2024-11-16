package translators

import (
	"os"
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestDir(t *testing.T) {
	d := Dir{}
	_, err := d.Translate("/non-existant-dir")
	test.ContainsError(os.ErrNotExist, err, t)

	p, err := d.Translate(".")
	test.Nil(err, t)
	test.Eq(p, ".", t)
}

func TestFile(t *testing.T) {
	f := File{}
	_, err := f.Translate("/non-existant-file")
	test.ContainsError(os.ErrNotExist, err, t)

	p, err := f.Translate("./FileSys_test.go")
	test.Nil(err, t)
	test.Eq(p, "./FileSys_test.go", t)
}

func TestOpenFile(t *testing.T) {
	f := NewOpenFile().SetFlags(os.O_RDONLY)
	fHandle, err := f.Translate("/non-existant-file")
	test.ContainsError(os.ErrNotExist, err, t)
	test.NilPntr[os.File](fHandle, t)

	f = NewOpenFile().SetFlags(os.O_RDONLY)
	fHandle, err = f.Translate("./FileSys_test.go")
	test.Nil(err, t)
	test.NotNilPntr[os.File](fHandle, t)
	fHandle.Close()
}
