package translators

import (
	"strconv"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinUint
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinUint8
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinUint16
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinUint32
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinUint64

type (
	// Represents a cmd line argument that will be translated to a uint type.
	//gen:ifaceImplCheck ifaceName Translator[uint]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinUint struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a uint8 type.
	//gen:ifaceImplCheck ifaceName Translator[uint8]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinUint8 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a uint16 type.
	//gen:ifaceImplCheck ifaceName Translator[uint16]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinUint16 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a uint32 type.
	//gen:ifaceImplCheck ifaceName Translator[uint32]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinUint32 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a uint64 type.
	//gen:ifaceImplCheck ifaceName Translator[uint64]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinUint64 struct {
		Base int
	}
)

func (u BuiltinUint) Translate(arg string) (uint, error) {
	// bit size of 0 corresponds with uint
	u64, err := strconv.ParseUint(arg, u.Base, 0)
	return uint(u64), err
}

func (u BuiltinUint) Reset() {
	// intentional noop - BuiltinUint has no state
}

func (u BuiltinUint8) Translate(arg string) (uint8, error) {
	u64, err := strconv.ParseUint(arg, u.Base, 8)
	return uint8(u64), err
}

func (u BuiltinUint8) Reset() {
	// intentional noop - BuiltinUint8 has no state
}

func (u BuiltinUint16) Translate(arg string) (uint16, error) {
	u64, err := strconv.ParseUint(arg, u.Base, 16)
	return uint16(u64), err
}

func (u BuiltinUint16) Reset() {
	// intentional noop - BuiltinUint16 has no state
}

func (u BuiltinUint32) Translate(arg string) (uint32, error) {
	u64, err := strconv.ParseUint(arg, u.Base, 32)
	return uint32(u64), err
}

func (u BuiltinUint32) Reset() {
	// intentional noop - BuiltinUint32 has no state
}

func (u BuiltinUint64) Translate(arg string) (uint64, error) {
	u64, err := strconv.ParseUint(arg, u.Base, 64)
	return u64, err
}

func (u BuiltinUint64) Reset() {
	// intentional noop - BuiltinUint64 has no state
}
