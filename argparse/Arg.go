package argparse

import (
	"github.com/barbell-math/util/argparse/computers"
	"github.com/barbell-math/util/argparse/translators"
	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/widgets"
)

//go:generate ../bin/structDefaultInit -struct opts

type (
	// The optional values that are associated with an argument.
	opts[T any, U translators.Translater[T]] struct {
		argType     ArgType `default:"ValueArgType" setter:"t" getter:"t"`
		shortName   byte    `default:"byte(0)" setter:"t" getter:"t"`
		required    bool    `default:"false" setter:"t" getter:"t"`
		description string  `default:"\"\"" setter:"t" getter:"t"`
		defaultVal  T       `default:"generics.ZeroVal[T]()" setter:"t" getter:"t" import:"github.com/barbell-math/util/generics"`
		translator  U       `default:"generics.ZeroVal[U]()" setter:"t" getter:"t" import:"github.com/barbell-math/util/generics github.com/barbell-math/util/argparse/translators"`
	}

	// Represents a single argument from the cmd line interface and all the
	// options associated with it.
	Arg struct {
		setVal        func(a *Arg, arg string) error
		setDefaultVal func()
		reset func()
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
		reset func()
	}

	// Used when only the shortName field of an Arg is important
	shortArg Arg
	// Used when only the longName field of an Arg is important
	longArg Arg
)

func NewArg[T any, U translators.Translater[T]](
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
		reset: func() { opts.translator.Reset() },
		argType: opts.argType,
		required:    opts.required,
		description: opts.description,
		shortFlag:   opts.shortName,
		longFlag:    longName,
		present:     false,
	}
}

func NewComputedArg[T any, U computers.Computer[T]](
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
	return l.shortFlag==r.shortFlag
}
func (_ *shortArg) Hash(other *shortArg) hash.Hash {
	return hash.Hash(other.shortFlag)
}
func (_ *shortArg) Zero(other *shortArg) {
	// intentional noop
}

func (_ *longArg) Eq(l *longArg, r *longArg) bool {
	return l.longFlag==r.longFlag
}
func (_ *longArg) Hash(other *longArg) hash.Hash {
	s:=widgets.BuiltinString{}
	return s.Hash(&other.longFlag)
}
func (_ *longArg) Zero(other *longArg) {
	// intentional noop
}
