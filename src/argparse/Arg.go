package argparse

import (
	"github.com/barbell-math/util/src/argparse/computers"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/hash"
	"github.com/barbell-math/util/src/widgets"
)

//go:generate ../../bin/structDefaultInit -struct opts

type (
	// The optional values that are associated with an argument.
	opts[T any, U translators.Translater[T]] struct {
		// The type of argument. This value will affect how the parser expects
		// values, so make sure it is the right value. See [ArgType] for
		// descriptions of available types.
		argType ArgType `default:"ValueArgType" setter:"t" getter:"t"`
		// The short name to associate with the argument. These will usually
		// follow a form similar to '-t'.
		shortName byte `default:"byte(0)" setter:"t" getter:"t"`
		// Defines if the argument is required or not.
		required bool `default:"false" setter:"t" getter:"t"`
		// Sets the description that will be printed out on the help menu.
		description string `default:"\"\"" setter:"t" getter:"t"`
		// The default value that should be used if the argument is not supplied.
		// The default defaults to a zero-value initilized value.
		defaultVal T `default:"generics.ZeroVal[T]()" setter:"t" getter:"t" import:"github.com/barbell-math/util/src/generics"`
		// The translator value to use when parsing the cmd line argument's
		// value. Most translators are stateless, but some have state and hence
		// must be able to have there value explicitly set.
		translator U `default:"generics.ZeroVal[U]()" setter:"t" getter:"t" import:"github.com/barbell-math/util/src/generics github.com/barbell-math/util/src/argparse/translators"`
	}

	// Represents a single argument from the cmd line interface and all the
	// options associated with it.
	Arg struct {
		setVal        func(a *Arg, arg string) error
		setDefaultVal func()
		reset         func(a *Arg)
		shortFlag     byte
		longFlag      string
		argType       ArgType
		present       bool
		required      bool
		description   string
	}

	// Represents an argument value that is computed from other arguments rather
	// than being supplied on the cmd line interface.
	ComputedArg struct {
		setVal func() error
		reset  func()
	}

	// Used when only the shortName field of an Arg is important
	shortArg Arg
	// Used when only the longName field of an Arg is important
	longArg Arg
)

func newArg[T any, U translators.Translater[T]](
	val *T,
	longName string,
	opts *opts[T, U],
) Arg {
	return Arg{
		setVal: func(a *Arg, arg string) error {
			v, err := opts.translator.Translate(arg)
			if err != nil {
				return err
			}
			*val = v
			a.present = true
			return nil
		},
		setDefaultVal: func() {
			*val = opts.defaultVal
		},
		reset: func(a *Arg) {
			opts.translator.Reset()
			a.present = false
		},
		argType:     opts.argType,
		required:    opts.required,
		description: opts.description,
		shortFlag:   opts.shortName,
		longFlag:    longName,
		present:     false,
	}
}

func newComputedArg[T any, U computers.Computer[T]](
	val *T,
	computer U,
) ComputedArg {
	return ComputedArg{
		setVal: func() error {
			v, err := computer.ComputeVals()
			if err != nil {
				return err
			}
			*val = v
			return nil
		},
		reset: func() { computer.Reset() },
	}
}

func (_ *shortArg) Eq(l *shortArg, r *shortArg) bool {
	return l.shortFlag == r.shortFlag
}
func (_ *shortArg) Hash(other *shortArg) hash.Hash {
	return hash.Hash(other.shortFlag)
}
func (_ *shortArg) Zero(other *shortArg) {
	// intentional noop
}

func (_ *longArg) Eq(l *longArg, r *longArg) bool {
	return l.longFlag == r.longFlag
}
func (_ *longArg) Hash(other *longArg) hash.Hash {
	s := widgets.BuiltinString{}
	return s.Hash(&other.longFlag)
}
func (_ *longArg) Zero(other *longArg) {
	// intentional noop
}
