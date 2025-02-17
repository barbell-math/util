package argparse

import (
	"fmt"

	"github.com/barbell-math/util/src/argparse/computers"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/hash"
	"github.com/barbell-math/util/src/widgets"
)

//go:generate ../../bin/structDefaultInit -struct opts

type (
	// A function type that defines when conditional args should be enforced.
	// The value supplied to the function will be the translated value that was
	// supplied on the cmd line.
	ArgConditional[T any] func(v T) bool

	// Defines which arguments are conditionally required as well as when they
	// are required.
	ArgConditionality[T any] struct {
		Requires []string
		When     ArgConditional[T]
	}

	// The optional values that are associated with an argument.
	//gen:structDefaultInit newReturns pntr
	opts[T translators.Translater[U], U any] struct {
		// The type of argument. This value will affect how the parser expects
		// values, so make sure it is the right value. See [ArgType] for
		// descriptions of available types.
		//gen:structDefaultInit default ValueArgType
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		argType ArgType
		// The short name to associate with the argument. These will usually
		// follow a form similar to '-t'.
		//gen:structDefaultInit default byte(0)
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		shortName byte
		// Defines if the argument is required or not.
		//gen:structDefaultInit default false
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		required bool
		// The list of arguments that must also be provided if this argument is
		// provided. All arguments provided are expected to be long names.
		//gen:structDefaultInit default []ArgConditionality[T]{}
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		conditionallyRequired []ArgConditionality[U]
		// Sets the description that will be printed out on the help menu.
		//gen:structDefaultInit default ""
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		description string
		// The default value that should be used if the argument is not supplied.
		// The default defaults to a zero-value initialized value.
		//gen:structDefaultInit default generics.ZeroVal[T]()
		//gen:structDefaultInit getter
		//gen:structDefaultInit imports github.com/barbell-math/util/src/generics
		defaultVal U
		//gen:structDefaultInit default false
		defaultValProvided bool
		// The translator value to use when parsing the cmd line argument's
		// value. Most translators are stateless, but some have state and hence
		// must be able to have there value explicitly set.
		//gen:structDefaultInit default generics.ZeroVal[U]()
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		//gen:structDefaultInit imports github.com/barbell-math/util/src/generics github.com/barbell-math/util/src/argparse/translators
		translator T
	}

	// Represents a single argument from the cmd line interface and all the
	// options associated with it.
	arg struct {
		setVal                func(a *arg, arg string) error
		setDefaultVal         func()
		defaultValAsStr       func() (string, bool)
		reset                 func(a *arg)
		conditionallyRequires func() []string
		allConditionalArgs    func() []string
		shortFlag             byte
		longFlag              string
		description           string
		argType               ArgType
		present               bool
		required              bool
		defaultProvided       bool
	}

	// Represents an argument value that is computed from other arguments rather
	// than being supplied on the cmd line interface.
	computedArg struct {
		setVal func() error
		reset  func()
	}

	// Used when only the shortName field of an Arg is important
	shortArg arg
	// Used when only the longName field of an Arg is important
	longArg arg
)

// A helper function that is intended to be used with the 'When' field of the
// [ArgConditionality] struct.
func ArgSupplied[T any](v T) bool { return true }

// A helper function that is intended to be used with the 'When' field of the
// [ArgConditionality] struct.
func ArgEquals[T comparable](val T) ArgConditional[T] {
	return func(v T) bool {
		return val == v
	}
}

// The default value that should be used if the argument is not supplied.
// The default defaults to a zero-value initialized value.
func (o *opts[T, U]) SetDefaultVal(v U) *opts[T, U] {
	o.defaultVal = v
	o.defaultValProvided = true
	return o
}

func newArg[T translators.Translater[U], U any](
	val *U,
	longName string,
	opts *opts[T, U],
) arg {
	return arg{
		setVal: func(a *arg, arg string) error {
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
		defaultValAsStr: func() (string, bool) {
			if opts.argType != FlagArgType && !opts.required {
				return fmt.Sprint(opts.defaultVal), true
			}
			return "", false
		},
		reset: func(a *arg) {
			opts.translator.Reset()
			a.present = false
		},
		conditionallyRequires: func() []string {
			rv := []string{}
			for _, v := range opts.conditionallyRequired {
				// Explicit copy of val here. It is not intended to be modified
				// by the when function.
				if v.When(*val) {
					rv = append(rv, v.Requires...)
				}
			}
			return rv
		},
		allConditionalArgs: func() []string {
			rv := []string{}
			for _, v := range opts.conditionallyRequired {
				rv = append(rv, v.Requires...)
			}
			return rv
		},
		argType:         opts.argType,
		required:        opts.required,
		description:     opts.description,
		shortFlag:       opts.shortName,
		longFlag:        longName,
		present:         false,
		defaultProvided: opts.defaultValProvided,
	}
}

func newComputedArg[T computers.Computer[U], U any](
	val *U,
	computer T,
) computedArg {
	return computedArg{
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
func (s *shortArg) String() string {
	return string(s.shortFlag)
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
func (l *longArg) String() string {
	return l.longFlag
}
