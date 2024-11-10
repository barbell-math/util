package argparse

import (
	"fmt"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/widgets"
)

type (
	Parser struct {
		progName string
		progDesc string

		subParsers [][]Arg
		requiredArgs containers.HashMap[
			*string,
			*longArg,
			widgets.BasePntr[string, widgets.BuiltinString],
			widgets.BasePntr[longArg, *longArg],
		]
		shortArgs containers.HashMap[
			*byte,
			*shortArg,
			widgets.BasePntr[byte, widgets.BuiltinByte],
			widgets.BasePntr[shortArg, *shortArg],
		]
		longArgs containers.HashMap[
			*string,
			*longArg,
			widgets.BasePntr[string, widgets.BuiltinString],
			widgets.BasePntr[longArg, *longArg],
		]
	}
)

func (p *Parser) getShortArg(b byte) (*Arg, error) {
	a, err:=p.shortArgs.Get(&b)
	if err!=nil {
		return nil, customerr.Wrap(UnrecognizedShortArgErr, "Argument: '%c'", b)
	}
	return (*Arg)(a), nil
}

func (p *Parser) getLongArg(s string) (*Arg, error) {
	a, err:=p.longArgs.Get(&s)
	if err!=nil {
		return nil, customerr.Wrap(UnrecognizedLongArgErr, "Argument: '%s'", s)
	}
	return (*Arg)(a), nil
}

func (p *Parser) Combine(others ...*Parser) (*Parser, error) {
	for _, otherP:=range(others) {
		p.subParsers=append(p.subParsers, otherP.subParsers...)
		if err:=containers.MapDisjointKeyedUnion[*byte, *shortArg](
			&p.shortArgs, &otherP.shortArgs,
		); err!=nil {
			customerr.AppendError(ParserCombinationErr, err)
		}
		if err:=containers.MapDisjointKeyedUnion[*string, *longArg](
			&p.longArgs, &otherP.longArgs,
		); err!=nil {
			customerr.AppendError(ParserCombinationErr, err)
		}
		// Required args are a subset of longArgs, no need to check for dups
		containers.MapKeyedUnion[*string, *longArg](
			&p.requiredArgs, &otherP.requiredArgs,
		)
	}
	return p, nil
}

func (p *Parser) Parse(t tokenIter) error {
	for _, subP:=range(p.subParsers) {
		for _, arg:=range(subP) {
			arg.setDefaultVal()
		}
	}

	if err:=t.toArgValPairs(p).ToIter().ForEach(
		func(
			index int,
			val basic.Pair[*Arg, string],
		) (iter.IteratorFeedback, error) {
			if val.A.present {
				return iter.Break, customerr.Wrap(
					ArgumentPassedMultipleTimesErr,
					"'%s'", val.A.longFlag,
				)
			}

			if err:=val.A.setVal(val.A, val.B); err!=nil {
				return iter.Break, customerr.AppendError(
					customerr.Wrap(
						ArgumentTranslationErr,
						"Argument: '%s'", val.A.longFlag,
					),
					err,
				)
			}

			return iter.Continue, nil
		},
	); err==HelpErr {
		fmt.Println(p.Help())
		return err
	} else if err!=nil {
		return customerr.AppendError(ParsingErr, err)
	}

	if err:=p.requiredArgs.Vals().ForEach(
		func(index int, val *longArg) (iter.IteratorFeedback, error) {
			if !val.present {
				return iter.Break, customerr.Wrap(
					MissingRequiredArgErr, "'%s'", val.longFlag,
				)
			}
			return iter.Continue, nil
		},
	); err!=nil {
		return customerr.AppendError(ParsingErr, err)
	}

	return nil
}

func (p *Parser) Help() string {
	return "HELP ME!"
}
