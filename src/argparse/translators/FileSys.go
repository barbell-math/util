package translators

import (
	"io/fs"
	"os"

	"github.com/barbell-math/util/src/customerr"
)

//go:generate ../../../bin/structDefaultInit -struct OpenFile
//go:generate ../../../bin/structDefaultInit -struct Mkdir

type (
	// A translator that checks that the supplied directory exists.
	Dir struct{}
	// A translator that checks that the supplied file exists.
	File struct{}

	// A translator that makes the supplied directory along with all necessary
	// parent directories.
	Mkdir struct {
		// The permissions used to create all dirs and sub-dirs. See
		// [os.MkdirAll] for reference.
		permissions fs.FileMode `default:"0644" setter:"t" getter:"t" import:"io/fs"`
	}
	// A translator that makes the supplied file.
	OpenFile struct{
		// The flags used to determine the file mode. See [os.RDONLY] and
		// friends.
		flags int `default:"os.O_RDONLY" setter:"t" getter:"t" import:"os"`
		// The permissions used to open the file with. See [os.OpenFile] for
		// reference.
		permissions fs.FileMode `default:"0644" setter:"t" getter:"t" import:"io/fs"`
	}
)

func (_ Dir) Translate(arg string) (string, error) {
	info, err:=os.Stat(arg)
	if err!=nil {
		return "", err
	}
	if !info.IsDir() {
		return "", customerr.Wrap(os.ErrNotExist, arg)
	}
	return arg, nil
}

func (m Mkdir) Translate(arg string) (string, error) {
	return arg, os.MkdirAll(arg, m.permissions)
}

func (_ Dir) Reset() {
	// intentional noop - IsDir has no state
}

func (m Mkdir) Reset() {
	// intentional noop - mkdir has no state
}

func (_ File) Translate(arg string) (string, error) {
	info, err:=os.Stat(arg)
	if err!=nil {
		return "", err
	}
	if !info.Mode().IsRegular() {
		return "", customerr.Wrap(os.ErrNotExist, arg)
	}
	return arg, nil
}

func (o OpenFile) Translate(arg string) (*os.File, error) {
	// If the file is not being created check that it exists.
	if o.flags | os.O_CREATE == 0 {
		info, err:=os.Stat(arg)
		if err!=nil {
			return nil, err
		}
		if info.Mode().IsRegular() {
			return nil, customerr.Wrap(os.ErrNotExist, arg)
		}
	}
	return os.OpenFile(arg, o.flags, o.permissions)
}

func (_ File) Reset() {
	// intentional noop - IsFile has no state
}

func (_ OpenFile) Reset() {
	// intentional noop - OpenFile has no state
}
