package argparse

import (
	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/widgets"
)

//go:generate ../bin/structDefaultInit -struct opts

type (
	// An interface that defines what actions can be performed when translating
	// a string argument to a typed value. The translator is expected to perform
	// all validation required to ensure a correct value is returned. It is also
	// expected to return an error if a value is found to be invalid.
	Translater[T any] interface {
		Translate(arg string) (T, error)
	}

	// The optional values that are associated with an argument.
	opts[T any, U Translater[T]] struct {
		argType     ArgType `default:"ValueArgType" setter:"t" getter:"t"`
		shortName   byte    `default:"byte(0)" setter:"t" getter:"t"`
		required    bool    `default:"false" setter:"t" getter:"t"`
		description string  `default:"\"\"" setter:"t" getter:"t"`
		defaultVal  T       `default:"generics.ZeroVal[T]()" setter:"t" getter:"t" import:"github.com/barbell-math/util/generics"`
		translator  U       `default:"generics.ZeroVal[U]()" setter:"t" getter:"t" import:"github.com/barbell-math/util/generics"`
	}

	// Represents a single argument from the cmd line interface and all the
	// options associated with it.
	Arg struct {
		setVal        func(a *Arg, arg string) error
		setDefaultVal func()
		shortFlag     byte
		longFlag      string
		argType       ArgType
		present       bool
		required      bool
		description   string
	}

	shortArg Arg
	longArg Arg
)

func NewArg[T any, U Translater[T]](
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
		argType: opts.argType,
		required:    opts.required,
		description: opts.description,
		shortFlag:   opts.shortName,
		longFlag:    longName,
		present:     false,
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
