package translators

import (
	"strconv"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinInt
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinInt8
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinInt16
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinInt32
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinInt64

type (
	// Represents a cmd line argument that will be translated to a int type.
	//gen:ifaceImplCheck ifaceName Translator[int]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinInt struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a int8 type.
	//gen:ifaceImplCheck ifaceName Translator[int8]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinInt8 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a int16 type.
	//gen:ifaceImplCheck ifaceName Translator[int16]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinInt16 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a int32 type.
	//gen:ifaceImplCheck ifaceName Translator[int32]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinInt32 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a int64 type.
	//gen:ifaceImplCheck ifaceName Translator[int64]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinInt64 struct {
		Base int
	}
)

func (i BuiltinInt) Translate(arg string) (int, error) {
	// bit size of 0 corresponds with int
	i64, err := strconv.ParseInt(arg, i.Base, 0)
	return int(i64), err
}

func (i BuiltinInt) Reset() {
	// intentional noop - BuiltinInt has no state
}

func (i BuiltinInt8) Translate(arg string) (int8, error) {
	i64, err := strconv.ParseInt(arg, i.Base, 8)
	return int8(i64), err
}

func (i BuiltinInt8) Reset() {
	// intentional noop - BuiltinInt8 has no state
}

func (i BuiltinInt16) Translate(arg string) (int16, error) {
	i64, err := strconv.ParseInt(arg, i.Base, 16)
	return int16(i64), err
}

func (i BuiltinInt16) Reset() {
	// intentional noop - BuiltinInt16 has no state
}

func (i BuiltinInt32) Translate(arg string) (int32, error) {
	i64, err := strconv.ParseInt(arg, i.Base, 32)
	return int32(i64), err
}

func (i BuiltinInt32) Reset() {
	// intentional noop - BuiltinInt32 has no state
}

func (i BuiltinInt64) Translate(arg string) (int64, error) {
	i64, err := strconv.ParseInt(arg, i.Base, 64)
	return i64, err
}

func (i BuiltinInt64) Reset() {
	// intentional noop - BuiltinInt64 has no state
}
