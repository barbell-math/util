package translators

// Code generated by ../../../bin/structDefaultInit - DO NOT EDIT.
import (
	"io/fs"
	"os"
)

// Returns a new OpenFile struct initialized with the default values.
func NewOpenFile() *OpenFile {
	return &OpenFile{
		flags:       os.O_RDONLY,
		permissions: 0644,
	}
}

// The flags used to determine the file mode. See [os.RDONLY] and
// friends.
func (o *OpenFile) SetFlags(v int) *OpenFile {
	o.flags = v
	return o
}

// The permissions used to open the file with. See [os.OpenFile] for
// reference.
func (o *OpenFile) SetPermissions(v fs.FileMode) *OpenFile {
	o.permissions = v
	return o
}

// The flags used to determine the file mode. See [os.RDONLY] and
// friends.
func (o *OpenFile) GetFlags() int {
	return o.flags
}

// The permissions used to open the file with. See [os.OpenFile] for
// reference.
func (o *OpenFile) GetPermissions() fs.FileMode {
	return o.permissions
}
