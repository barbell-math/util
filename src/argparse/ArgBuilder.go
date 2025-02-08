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
func AddArg[T any, U translators.Translater[T]](
	val *T,
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
	opts *opts[bool, translators.Flag],
) {
	if opts == nil {
		opts = NewOpts[bool, translators.Flag]()
	}
	opts.argType = FlagArgType
	AddArg[bool, translators.Flag](val, builder, longName, opts)
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
	opts *opts[T, *translators.FlagCntr[T]],
) {
	if opts == nil {
		opts = NewOpts[T, *translators.FlagCntr[T]]()
	}
	opts.argType = MultiFlagArgType
	opts.translator = &translators.FlagCntr[T]{}
	AddArg[T, *translators.FlagCntr[T]](val, builder, longName, opts)
}

// Appends a list argument to the supplied builder without performing any
// validation of the argument or builder as a whole. List arguments accept many
// values for a single flag and will return a slice of all the translated values.
//
// The arg type of the opts struct will be set to [MultiFlagArgType].
//
// If opts is Nil then the opts will be populated with the default values from
// calling NewOpts.
func AddListArg[T any, U translators.Translater[T], W widgets.BaseInterface[T]](
	val *[]T,
	builder *ArgBuilder,
	longName string,
	opts *opts[[]T, *translators.ListValues[T, U, W]],
) {
	if opts == nil {
		opts = NewOpts[[]T, *translators.ListValues[T, U, W]]()
	}
	opts.argType = MultiValueArgType
	AddArg[[]T, *translators.ListValues[T, U, W]](val, builder, longName, opts)
}

// Appends a selector argument to the supplied builder without performing any
// validation of the argument or builder as a whole. Selector arguments accept
// one value that must be one of a predefined set of values.
//
// The arg type of the opts struct will be set to [ValueArgType].
//
// If opts is Nil then the opts will be populated with the default values from
// calling NewOpts.
func AddSelector[T any, U translators.Translater[T], W widgets.BaseInterface[T]](
	val *T,
	builder *ArgBuilder,
	longName string,
	opts *opts[T, translators.Selector[T, U, W]],
) {
	if opts == nil {
		opts = NewOpts[T, translators.Selector[T, U, W]]()
	}
	opts.argType = ValueArgType
	AddArg[T, translators.Selector[T, U, W]](val, builder, longName, opts)
}

// Appends a enum selector to the supplied builder without performing any
// validation of the argument or builder as a whole. Enum selector arguments
// accept a value that must map to a valid enum value of the supplied enum
// type.
func AddEnum[E enum.Value, EP enum.Pntr[E]](
	val *E,
	builder *ArgBuilder,
	longName string,
	opts *opts[E, translators.Enum[E, EP]],
) {
	if opts == nil {
		opts = NewOpts[E, translators.Enum[E, EP]]()
	}
	opts.argType = ValueArgType
	opts.translator = translators.Enum[E, EP]{}
	AddArg[E, translators.Enum[E, EP]](val, builder, longName, opts)
}

// Appends a computed argument to the supplied builder without performing any
// validation of computation of the argument or builder as a whole.
func AddComputedArg[T any, U computers.Computer[T]](
	val *T,
	builder *ArgBuilder,
	computer U,
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
			if _, err := rv.shortArgs.Get(&b.args[i].shortFlag); err == nil {
				return rv, customerr.AppendError(
					ParserConfigErr,
					customerr.Wrap(
						DuplicateShortNameErr, "'%c'", b.args[i].shortFlag,
					),
				)
			}
		}
		if _, err := rv.longArgs.Get(&b.args[i].longFlag); err == nil {
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
			rv.shortArgs.Emplace(containerBasic.Pair[*byte, *shortArg]{
				&b.args[i].shortFlag, (*shortArg)(&b.args[i]),
			})
		}
		rv.longArgs.Emplace(containerBasic.Pair[*string, *longArg]{
			&b.args[i].longFlag, (*longArg)(&b.args[i]),
		})
		if b.args[i].required {
			rv.requiredArgs.Emplace(containerBasic.Pair[*string, *longArg]{
				&b.args[i].longFlag, (*longArg)(&b.args[i]),
			})
		}
	}

	return rv, nil
}
