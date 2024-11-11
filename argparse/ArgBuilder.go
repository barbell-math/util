package argparse

import (
	"github.com/barbell-math/util/argparse/computers"
	"github.com/barbell-math/util/argparse/translators"
	containerBasic "github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/customerr"
	mathBasic "github.com/barbell-math/util/math/basic"
)

type (
	// Used to create a list of arguments that are then used to build the parser.
	ArgBuilder struct {
		args         []Arg
		computedVals []ComputedArg
	}
)

// Appends an argument to the supplied builder without performing any validation
// of the argument or builder as a whole.
func AddArg[T any, U translators.Translater[T]](
	val *T,
	builder *ArgBuilder,
	longName string,
	opts *opts[T, U],
) {
	builder.args = append(builder.args, NewArg[T, U](val, longName, opts))
}

// Appends a flag argument to the supplied builder without performing any
// validation of the argument or builder as a whole. Counter flags can only be
// supplied once.
// The arg type of the opts struct will be set to [FlagArgType].
func AddFlag(
	val *bool,
	builder *ArgBuilder,
	longName string,
	opts *opts[bool, translators.Flag],
) {
	opts.argType = FlagArgType
	AddArg[bool, translators.Flag](val, builder, longName, opts)
}

// Appends a flag counter argument to the supplied builder without performing
// any validation of the argument or builder as a whole. Counter flags can be
// supplied many times. The counter will represent the total number of times the
// flag was supplied.
// The arg type of the opts struct will be set to [MultiFlagArgType].
func AddFlagCntr[T mathBasic.Int | mathBasic.Uint](
	val *T,
	builder *ArgBuilder,
	longName string,
	opts *opts[T, *translators.FlagCntr[T]],
) {
	opts.argType = MultiFlagArgType
	AddArg[T, *translators.FlagCntr[T]](val, builder, longName, opts)
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
		NewComputedArg[T, U](val, computer),
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
		if _, err := rv.shortArgs.Get(&b.args[i].shortFlag); err == nil {
			return rv, customerr.AppendError(
				ParserConfigErr,
				customerr.Wrap(
					DuplicateShortNameErr, "'%c'", b.args[i].shortFlag,
				),
			)
		}
		if _, err := rv.longArgs.Get(&b.args[i].longFlag); err == nil {
			return rv, customerr.AppendError(
				ParserConfigErr,
				customerr.Wrap(DuplicateLongNameErr, "'%s'", b.args[i].longFlag),
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

		rv.shortArgs.Emplace(containerBasic.Pair[*byte, *shortArg]{
			&b.args[i].shortFlag, (*shortArg)(&b.args[i]),
		})
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
