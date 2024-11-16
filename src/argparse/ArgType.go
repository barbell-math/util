package argparse

//go:generate ../../bin/enum -type=ArgType -package=argparse
//go:generate ../../bin/passThroughWidget -type=ArgType

type (
	//gen:enum unknownValue UnknownArgType
	//gen:enum default UnknownArgType
	//gen:passThroughWidget widgetType Base
	//gen:passThroughWidget baseTypeWidget widgets.BuiltinInt
	//gen:passThroughWidget widgetPackage github.com/barbell-math/util/src/widgets
	ArgType int
)

const (
	//gen:enum string UnknownArgType
	UnknownArgType ArgType = iota
	// Represents a flag type that must accept a value as an argument and must
	// only be supplied once.
	//gen:enum string ValueArgType
	ValueArgType
	// Represents a flag type that can accept many values as an argument and
	// must only be supplied once. At least one argument must be supplied.
	//gen:enum string MultiValueArgType
	MultiValueArgType
	// Represents a flag type that must not accept a value and must only be
	// supplied once.
	//gen:enum string FlagArgType
	FlagArgType
	// Represents a flag type that must not accept a value and may be supplied
	// many times.
	//gen:enum string MultiFlagArgType
	MultiFlagArgType
)

var (
	singleSpecificationArgTypes = map[ArgType]struct{}{
		ValueArgType: struct{}{}, FlagArgType: struct{}{},
	}

	multiSpecificationArgTypes = map[ArgType]struct{}{
		MultiValueArgType: struct{}{}, MultiFlagArgType: struct{}{},
	}
)
