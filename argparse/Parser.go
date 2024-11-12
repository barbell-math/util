package argparse

import (
	"errors"
	"fmt"
	"strings"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/widgets"
)

type (
	computedArgsTree struct {
		compedArgs    []ComputedArg
		subCompedArgs []computedArgsTree
	}

	Parser struct {
		progName string
		progDesc string

		numArgs      int
		subParsers   [][]Arg
		compedArgs   computedArgsTree
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

const (
	helpDescriptionWidth int = 80
)

func (c *computedArgsTree) leftRightRootTraversal(
	op func(c *ComputedArg) error,
) error {
	if len(c.subCompedArgs) > 0 {
		for _, v := range c.subCompedArgs {
			if err := v.leftRightRootTraversal(op); err != nil {
				return err
			}
		}
	}
	for _, v := range c.compedArgs {
		if err := op(&v); err != nil {
			return err
		}
	}
	return nil
}

func newParser(
	progName string,
	progDesc string,
	args []Arg,
	computedArgs []ComputedArg,
) Parser {
	rv := Parser{
		progName:   progName,
		progDesc:   progDesc,
		numArgs:    len(args) + len(computedArgs),
		subParsers: [][]Arg{args},
		compedArgs: computedArgsTree{compedArgs: computedArgs},
	}
	rv.requiredArgs, _ = containers.NewHashMap[
		*string,
		*longArg,
		widgets.BasePntr[string, widgets.BuiltinString],
		widgets.BasePntr[longArg, *longArg],
	](len(args))
	rv.shortArgs, _ = containers.NewHashMap[
		*byte,
		*shortArg,
		widgets.BasePntr[byte, widgets.BuiltinByte],
		widgets.BasePntr[shortArg, *shortArg],
	](len(args))
	rv.longArgs, _ = containers.NewHashMap[
		*string,
		*longArg,
		widgets.BasePntr[string, widgets.BuiltinString],
		widgets.BasePntr[longArg, *longArg],
	](len(args))
	return rv
}

func (p *Parser) getShortArg(b byte) (*Arg, error) {
	a, err := p.shortArgs.Get(&b)
	if err != nil {
		return nil, customerr.Wrap(UnrecognizedShortArgErr, "Argument: '%c'", b)
	}
	return (*Arg)(a), nil
}

func (p *Parser) getLongArg(s string) (*Arg, error) {
	a, err := p.longArgs.Get(&s)
	if err != nil {
		return nil, customerr.Wrap(UnrecognizedLongArgErr, "Argument: '%s'", s)
	}
	return (*Arg)(a), nil
}

// Adds sub-parsers to the current parser. All arguments are placed in a global
// namespace and must be unique. No long or short names can collide. All
// computed args are added to a tree like data structure, which is used to
// maintain the desired bottom up order of execution for computed arguments.
func (p *Parser) AddSubParsers(others ...*Parser) error {
	for _, otherP := range others {
		if otherP.numArgs > 0 {
			p.subParsers = append(p.subParsers, otherP.subParsers...)
			if err := containers.MapDisjointKeyedUnion[*byte, *shortArg](
				&p.shortArgs, &otherP.shortArgs,
			); err != nil {
				return customerr.AppendError(
					ParserCombinationErr, DuplicateShortNameErr, err,
				)
			}
			if err := containers.MapDisjointKeyedUnion[*string, *longArg](
				&p.longArgs, &otherP.longArgs,
			); err != nil {
				return customerr.AppendError(
					ParserCombinationErr, DuplicateLongNameErr, err,
				)
			}
			// Required args are a subset of longArgs, no need to check for dups
			containers.MapKeyedUnion[*string, *longArg](
				&p.requiredArgs, &otherP.requiredArgs,
			)
			p.compedArgs.subCompedArgs = append(
				p.compedArgs.subCompedArgs, otherP.compedArgs,
			)
			p.numArgs += otherP.numArgs
		}
	}
	return nil
}

func (p *Parser) Parse(t tokenIter) error {
	for _, subP := range p.subParsers {
		for i, arg := range subP {
			arg.setDefaultVal()
			arg.reset(&arg)
			subP[i] = arg
		}
	}
	p.compedArgs.leftRightRootTraversal(func(c *ComputedArg) error {
		c.reset()
		return nil
	})

	if err := t.toArgValPairs(p).ToIter().ForEach(
		func(
			index int,
			val basic.Pair[*Arg, string],
		) (iter.IteratorFeedback, error) {
			if _, ok := multiSpecificationArgTypes[val.A.argType]; !ok && val.A.present {
				return iter.Break, customerr.Wrap(
					ArgumentPassedMultipleTimesErr,
					"'%s'", val.A.longFlag,
				)
			}

			if err := val.A.setVal(val.A, val.B); err != nil {
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
	); errors.Is(err, HelpErr) {
		fmt.Println(p.Help())
		return err
	} else if err != nil {
		return customerr.AppendError(ParsingErr, err)
	}

	if err := p.requiredArgs.Vals().ForEach(
		func(index int, val *longArg) (iter.IteratorFeedback, error) {
			if !val.present {
				return iter.Break, customerr.Wrap(
					MissingRequiredArgErr, "'%s'", val.longFlag,
				)
			}
			return iter.Continue, nil
		},
	); err != nil {
		return customerr.AppendError(ParsingErr, err)
	}

	if err := p.compedArgs.leftRightRootTraversal(func(c *ComputedArg) error {
		return c.setVal()
	}); err != nil {
		return customerr.AppendError(ParsingErr, ComputedArgumentErr, err)
	}

	return nil
}

func (p *Parser) Help() string {
	start := ""
	longestArg, _ := p.longArgs.Keys().Reduce(
		&start,
		func(accum **string, iter *string) error {
			if len(*iter) > len(**accum) {
				*accum = iter
			}
			return nil
		},
	)

	var sb strings.Builder
	sb.WriteString(p.progName)
	sb.WriteString(": HELP ME!\n")
	sb.WriteString("Description: ")
	sb.WriteString(p.progDesc)
	sb.WriteByte('\n')
	sb.WriteByte('\n')

	p.longArgs.Vals().ForEach(
		func(index int, val *longArg) (iter.IteratorFeedback, error) {
			if val.shortFlag != byte(0) {
				sb.WriteString("-")
				sb.WriteByte(val.shortFlag)
				sb.WriteString("  ")
			}
			sb.WriteString("--")
			sb.WriteString(val.longFlag)
			for i := 0; i < len(*longestArg)-len(val.longFlag); i++ {
				sb.WriteByte(' ')
			}
			sb.WriteString(" ")

			if val.required {
				sb.WriteString(" (required)  ")
			} else {
				sb.WriteString("             ")
			}

			cntr := 0
			for _, s := range strings.Split(val.description, " ") {
				if cntr+len(s) > helpDescriptionWidth {
					sb.WriteByte('\n')
					for i := 0; i < len(*longestArg)+21; i++ {
						sb.WriteByte(' ')
					}
					cntr = 0
				}
				sb.WriteString(s)
				sb.WriteString(" ")
				cntr += len(s) + 1
			}
			sb.WriteByte('\n')
			return iter.Continue, nil
		},
	)

	return sb.String()
}
