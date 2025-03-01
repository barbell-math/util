package argparse

import (
	"github.com/barbell-math/util/src/argparse/computers"
	"github.com/barbell-math/util/src/argparse/translators"
	containerBasic "github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/enum"
	mathBasic "github.com/barbell-math/util/src/math/basic"
	"github.com/barbell-math/util/src/widgets"
)

type (
	// Used to create a list of arguments that are then used to build the parser.
	ArgBuilder struct {
		args         []arg
		computedVals []computedArg
	}
)

// Appends an argument to the supplied builder without performing any validation
// of the argument or builder as a whole.
//
// If opts is Nil then the opts will be populated with the default values from
// calling NewOpts.
func AddArg[T translators.Translator[U], U any](
	val *U,
	builder *ArgBuilder,
	longName string,
	opts *opts[T, U],
) {
	if opts == nil {
		opts = NewOpts[T, U]()
	}
	builder.args = append(builder.args, newArg[T, U](val, longName, opts))
}

// Appends a flag argument to the supplied builder without performing any
// validation of the argument or builder as a whole. Counter flags can only be
// supplied once.
//
// The arg type of the opts struct will be set to [FlagArgType].
//
// If opts is Nil then the opts will be populated with the default values from
// calling NewOpts.
func AddFlag(
	val *bool,
	builder *ArgBuilder,
	longName string,
	opts *opts[translators.Flag, bool],
) {
	if opts == nil {
		opts = NewOpts[translators.Flag]()
	}
	opts.argType = FlagArgType
	AddArg[translators.Flag](val, builder, longName, opts)
}

// Appends a flag counter argument to the supplied builder without performing
// any validation of the argument or builder as a whole. Counter flags can be
// supplied many times. The counter will represent the total number of times the
// flag was supplied.
//
// The arg type of the opts struct will be set to [MultiFlagArgType].
//
// The translator value in the opts struct will be initialized to a zero-valued
// [translator.FlagCntr].
//
// If opts is Nil then the opts will be populated with the default values from
// calling NewOpts.
func AddFlagCntr[T mathBasic.Int | mathBasic.Uint](
	val *T,
	builder *ArgBuilder,
	longName string,
	opts *opts[*translators.FlagCntr[T], T],
) {
	if opts == nil {
		opts = NewOpts[*translators.FlagCntr[T], T]()
	}
	opts.argType = MultiFlagArgType
	opts.translator = &translators.FlagCntr[T]{}
	AddArg[*translators.FlagCntr[T]](val, builder, longName, opts)
}

// Appends a list argument to the supplied builder without performing any
// validation of the argument or builder as a whole. List arguments accept many
// values for a single flag and will return a slice of all the translated values.
//
// The arg type of the opts struct will be set to [MultiFlagArgType].
//
// If opts is Nil then the opts will be populated with the default values from
// calling NewOpts.
func AddListArg[T translators.Translator[U], W widgets.BaseInterface[U], U any](
	val *[]U,
	builder *ArgBuilder,
	longName string,
	opts *opts[*translators.ListValues[T, W, U], []U],
) {
	if opts == nil {
		opts = NewOpts[*translators.ListValues[T, W, U], []U]()
	}
	opts.argType = MultiValueArgType
	AddArg[*translators.ListValues[T, W, U]](val, builder, longName, opts)
}

// Appends a selector argument to the supplied builder without performing any
// validation of the argument or builder as a whole. Selector arguments accept
// one value that must be one of a predefined set of values.
//
// The arg type of the opts struct will be set to [ValueArgType].
//
// If opts is Nil then the opts will be populated with the default values from
// calling NewOpts.
func AddSelector[T translators.Translator[U], W widgets.BaseInterface[U], U any](
	val *U,
	builder *ArgBuilder,
	longName string,
	opts *opts[translators.Selector[T, W, U], U],
) {
	if opts == nil {
		opts = NewOpts[translators.Selector[T, W, U], U]()
	}
	opts.argType = ValueArgType
	AddArg[translators.Selector[T, W, U]](val, builder, longName, opts)
}

// Appends a enum selector to the supplied builder without performing any
// validation of the argument or builder as a whole. Enum selector arguments
// accept a value that must map to a valid enum value of the supplied enum
// type.
func AddEnum[EP enum.Pntr[E], E enum.Value](
	val *E,
	builder *ArgBuilder,
	longName string,
	opts *opts[translators.Enum[EP, E], E],
) {
	if opts == nil {
		opts = NewOpts[translators.Enum[EP, E]]()
	}
	opts.argType = ValueArgType
	opts.translator = translators.Enum[EP, E]{}
	AddArg[translators.Enum[EP, E]](val, builder, longName, opts)
}

// Appends a computed argument to the supplied builder without performing any
// validation of computation of the argument or builder as a whole.
func AddComputedArg[T computers.Computer[U], U any](
	val *U,
	builder *ArgBuilder,
	computer T,
) {
	builder.computedVals = append(
		builder.computedVals,
		newComputedArg[T, U](val, computer),
	)
}

// Creates a parser using the arg builder. Note that the arg builder will be
// modified and should not be used again after calling ToParser. The previously
// added arguments will be validated. Validation can return one of the below
// errors, all warpped in a top level [ParserConfigErr]:
//
//   - [ReservedShortNameErr]
//   - [ReservedLongNameErr]
//   - [DuplicateShortNameErr]
//   - [DuplicateLongNameErr]
//   - [LongNameToShortErr]
func (b *ArgBuilder) ToParser(progName string, progDesc string) (Parser, error) {
	// After calling this function the args slice must not reallocate due to the
	// maps containing pointers to the slice values.
	rv := newParser(progName, progDesc, b.args, b.computedVals)
	for i := 0; i < len(b.args); i++ {
		if b.args[i].shortFlag != byte(0) {
			if _, err := rv.shortArgs.Get(b.args[i].shortFlag); err == nil {
				return rv, customerr.AppendError(
					ParserConfigErr,
					customerr.Wrap(
						DuplicateShortNameErr, "'%c'", b.args[i].shortFlag,
					),
				)
			}
		}
		if _, err := rv.longArgs.Get(b.args[i].longFlag); err == nil {
			return rv, customerr.AppendError(
				ParserConfigErr,
				customerr.Wrap(DuplicateLongNameErr, "'%s'", b.args[i].longFlag),
			)
		}

		if b.args[i].longFlag == "config" {
			return rv, customerr.AppendError(
				ParserConfigErr,
				customerr.Wrap(
					ReservedLongNameErr,
					"'config' is reserved for specifying argument config files",
				),
			)
		}
		if len(b.args[i].longFlag) < 2 {
			return rv, customerr.AppendError(
				ParserConfigErr,
				customerr.Wrap(
					LongNameToShortErr,
					"Name: '%s'", b.args[i].longFlag,
				),
			)
		}

		if b.args[i].shortFlag != byte(0) {
			rv.shortArgs.Emplace(containerBasic.Pair[byte, *shortArg]{
				b.args[i].shortFlag, (*shortArg)(&b.args[i]),
			})
		}
		rv.longArgs.Emplace(containerBasic.Pair[string, *longArg]{
			b.args[i].longFlag, (*longArg)(&b.args[i]),
		})
		if b.args[i].required {
			rv.requiredArgs.Emplace(containerBasic.Pair[string, *longArg]{
				b.args[i].longFlag, (*longArg)(&b.args[i]),
			})
		}
	}

	return rv, nil
}
