package translators

import (
	"io/fs"
	"os"

	"github.com/barbell-math/util/src/customerr"
)

//go:generate ../../../bin/structDefaultInit -struct OpenFile
//go:generate ../../../bin/structDefaultInit -struct Mkdir
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Dir
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=File
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Mkdir
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=OpenFile

type (
	// A translator that checks that the supplied directory exists.
	//gen:ifaceImplCheck ifaceName Translator[string]
	//gen:ifaceImplCheck valOrPntr both
	Dir struct{}
	// A translator that checks that the supplied file exists.
	//gen:ifaceImplCheck ifaceName Translator[string]
	//gen:ifaceImplCheck valOrPntr both
	File struct{}

	// A translator that makes the supplied directory along with all necessary
	// parent directories.
	//gen:structDefaultInit newReturns pntr
	//gen:ifaceImplCheck ifaceName Translator[string]
	//gen:ifaceImplCheck valOrPntr both
	Mkdir struct {
		// The permissions used to create all dirs and sub-dirs. See
		// [os.MkdirAll] for reference.
		//gen:structDefaultInit default 0644
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		//gen:structDefaultInit imports io/fs
		permissions fs.FileMode
	}
	// A translator that makes the supplied file.
	//gen:structDefaultInit newReturns pntr
	//gen:ifaceImplCheck ifaceName Translator[*os.File]
	//gen:ifaceImplCheck imports os
	//gen:ifaceImplCheck valOrPntr both
	OpenFile struct {
		// The flags used to determine the file mode. See [os.RDONLY] and
		// friends.
		//gen:structDefaultInit default os.O_RDONLY
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		//gen:structDefaultInit imports os
		flags int
		// The permissions used to open the file with. See [os.OpenFile] for
		// reference.
		//gen:structDefaultInit default 0644
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		//gen:structDefaultInit imports io/fs
		permissions fs.FileMode
	}
)

func (_ Dir) Translate(arg string) (string, error) {
	info, err := os.Stat(arg)
	if err != nil {
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
	info, err := os.Stat(arg)
	if err != nil {
		return "", err
	}
	if !info.Mode().IsRegular() {
		return "", customerr.Wrap(os.ErrNotExist, arg)
	}
	return arg, nil
}

func (o OpenFile) Translate(arg string) (*os.File, error) {
	// If the file is not being created check that it exists.
	if o.flags|os.O_CREATE == 0 {
		info, err := os.Stat(arg)
		if err != nil {
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
