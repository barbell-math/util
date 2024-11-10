package argparse

import (
	"github.com/barbell-math/util/argparse/translators"
	containerBasic "github.com/barbell-math/util/container/basic"
	mathBasic "github.com/barbell-math/util/math/basic"
	"github.com/barbell-math/util/container/containers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/widgets"
)

type (
	// Used to create a list of arguments that are then used to build the parser.
	ArgBuilder []Arg
)

// Appends an argument to the supplied builder without performing any validation
// of the argument or builder as a whole.
func AddArg[T any, U Translater[T]](
	val *T,
	builder *ArgBuilder,
	longName string,
	opts *opts[T, U],
) {
	*builder=append(*builder, NewArg[T, U](val, longName, opts))
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
	opts.argType=FlagArgType
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
	opts.argType=MultiFlagArgType
	AddArg[T, *translators.FlagCntr[T]](val, builder, longName, opts)
}

// Creates a parser using the arg builder. Note that the arg builder will be
// modified and should not be used again after calling ToParser. The previously
// added arguments will be validated. Validation can return one of the below
// errors, all warpped in a top level [ParserConfigErr]:
//
//  - [ReservedShortNameErr]
//  - [ReservedLongNameErr]
//  - [DuplicateShortNameErr]
//  - [DuplicateLongNameErr]
//  - [LongNameToShortErr]
func (b *ArgBuilder) ToParser(progName string, progDesc string) (Parser, error) {
	// After calling this function the args slice must not reallocate due to the
	// maps containing pointers to the slice values.
	// TODO - delete if subparsers work
	// *b=append(*b, reservedArgs...)

	rv:=Parser{progName: progName, progDesc: progDesc, subParsers: [][]Arg{*b}}
	rv.requiredArgs, _=containers.NewHashMap[
		*string,
		*longArg,
		widgets.BasePntr[string, widgets.BuiltinString],
		widgets.BasePntr[longArg, *longArg],
	](len(*b))
	rv.shortArgs, _=containers.NewHashMap[
		*byte,
		*shortArg,
		widgets.BasePntr[byte, widgets.BuiltinByte],
		widgets.BasePntr[shortArg, *shortArg],
	](len(*b))
	rv.longArgs, _=containers.NewHashMap[
		*string,
		*longArg,
		widgets.BasePntr[string, widgets.BuiltinString],
		widgets.BasePntr[longArg, *longArg],
	](len(*b))

	// for i:=0; i<len(*b)-len(reservedArgs); i++ { - TODO - delete
	for i:=0; i<len(*b); i++ {
		// TODO - delete if subparsers work
		// if err:=tmpShortName.FromString(string((*b)[i].shortFlag)); err==nil {
		// 	return rv, customerr.AppendError(
		// 		ParserConfigErr,
		// 		customerr.Wrap(ReservedShortNameErr, "'%c'", (*b)[i].shortFlag),
		// 	)
		// }
		// if err:=tmpLongName.FromString((*b)[i].longFlag); err==nil {
		// 	return rv, customerr.AppendError(
		// 		ParserConfigErr,
		// 		customerr.Wrap(ReservedLongNameErr, "'%s'", (*b)[i].longFlag),
		// 	)
		// }

		if _, err:=rv.shortArgs.Get(&(*b)[i].shortFlag); err==nil {
			return rv, customerr.AppendError(
				ParserConfigErr,
				customerr.Wrap(DuplicateShortNameErr, "'%c'", (*b)[i].shortFlag),
			)
		}
		if _, err:=rv.longArgs.Get(&(*b)[i].longFlag); err==nil {
			return rv, customerr.AppendError(
				ParserConfigErr,
				customerr.Wrap(DuplicateLongNameErr, "'%s'", (*b)[i].longFlag),
			)
		}

		if len((*b)[i].longFlag)<2 {
			return rv, customerr.AppendError(
				ParserConfigErr,
				customerr.Wrap(
					LongNameToShortErr,
					"Name: '%s'", (*b)[i].longFlag,
				),
			)
		}

		rv.shortArgs.Emplace(containerBasic.Pair[*byte, *shortArg]{
			&(*b)[i].shortFlag, (*shortArg)(&(*b)[i]),
		})
		rv.longArgs.Emplace(containerBasic.Pair[*string, *longArg]{
			&(*b)[i].longFlag, (*longArg)(&(*b)[i]),
		})
		if (*b)[i].required {
			rv.requiredArgs.Emplace(containerBasic.Pair[*string, *longArg]{
				&(*b)[i].longFlag, (*longArg)(&(*b)[i]),
			})
		}
	}

	return rv, nil
}
