package argparse

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/container/containers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/iter"
	"github.com/barbell-math/util/src/strops"
	"github.com/barbell-math/util/src/widgets"
)

type (
	computedArgsTree struct {
		compedArgs    []computedArg
		subCompedArgs []computedArgsTree
	}

	// The type that will take a token stream generated from a sequence of
	// strings and perform all translating and computing tasks. See the
	// [examples] package and README.md for more in depth information.
	Parser struct {
		progName string
		progDesc string

		numArgs      int
		subParsers   [][]arg
		compedArgs   computedArgsTree
		requiredArgs containers.HashMap[
			string,
			*longArg,
			widgets.BuiltinString,
			widgets.BasePntr[longArg, *longArg],
		]
		shortArgs containers.HashMap[
			byte,
			*shortArg,
			widgets.BuiltinByte,
			widgets.BasePntr[shortArg, *shortArg],
		]
		longArgs containers.HashMap[
			string,
			*longArg,
			widgets.BuiltinString,
			widgets.BasePntr[longArg, *longArg],
		]
	}
)

const (
	helpDescriptionWidth int = 80
)

func (c *computedArgsTree) leftRightRootTraversal(
	op func(c *computedArg) error,
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
	args []arg,
	computedArgs []computedArg,
) Parser {
	rv := Parser{
		progName:   progName,
		progDesc:   progDesc,
		numArgs:    len(args) + len(computedArgs),
		subParsers: [][]arg{args},
		compedArgs: computedArgsTree{compedArgs: computedArgs},
	}
	rv.requiredArgs, _ = containers.NewHashMap[
		string,
		*longArg,
		widgets.BuiltinString,
		widgets.BasePntr[longArg, *longArg],
	](len(args))
	rv.shortArgs, _ = containers.NewHashMap[
		byte,
		*shortArg,
		widgets.BuiltinByte,
		widgets.BasePntr[shortArg, *shortArg],
	](len(args))
	rv.longArgs, _ = containers.NewHashMap[
		string,
		*longArg,
		widgets.BuiltinString,
		widgets.BasePntr[longArg, *longArg],
	](len(args))
	return rv
}

func (p *Parser) getShortArg(b byte) (*arg, error) {
	a, err := p.shortArgs.Get(b)
	if err != nil {
		return nil, customerr.Wrap(UnrecognizedShortArgErr, "Argument: '%c'", b)
	}
	return (*arg)(a), nil
}

func (p *Parser) getLongArg(s string) (*arg, error) {
	a, err := p.longArgs.Get(s)
	if err != nil {
		return nil, customerr.Wrap(UnrecognizedLongArgErr, "Argument: '%s'", s)
	}
	return (*arg)(a), nil
}

// Adds sub-parsers to the current parser. All arguments are placed in a global
// namespace and must be unique. No long or short names can collide. All
// computed args are added to a tree like data structure, which is used to
// maintain the desired bottom up order of execution for computed arguments.
func (p *Parser) AddSubParsers(others ...Parser) error {
	for _, otherP := range others {
		if otherP.numArgs > 0 {
			p.subParsers = append(p.subParsers, otherP.subParsers...)
			if err := containers.MapDisjointKeyedUnion[byte, *shortArg](
				&p.shortArgs, &otherP.shortArgs,
			); err != nil {
				return customerr.AppendError(
					ParserCombinationErr, DuplicateShortNameErr, err,
				)
			}
			if err := containers.MapDisjointKeyedUnion[string, *longArg](
				&p.longArgs, &otherP.longArgs,
			); err != nil {
				return customerr.AppendError(
					ParserCombinationErr, DuplicateLongNameErr, err,
				)
			}
			// Required args are a subset of longArgs, no need to check for dups
			containers.MapKeyedUnion[string, *longArg](
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

// Parses the token stream given to it. This is a two step process. The steps
// are as follows:
//
//  1. Consume the tokens and translate all received values, saving the results
//
// to the desired locations.
//  2. Compute all computed arguments in a bottom-up, left-right fashion.
//
// If an error occurs in this process it will be returned wrapped in a top level
// [ParsingErr]. The only exception to this will be the [HelpErr], which will
// stop all further parsing, print the help menu, and return. Any tokens that
// were present before the help flag will be translated.
func (p *Parser) Parse(t tokenIter) error {
	// check the parser state
	if err := p.checkConditionalRequiredArgsExist(); err != nil {
		return err
	}

	// reset the parser state
	for _, subP := range p.subParsers {
		for i, arg := range subP {
			arg.setDefaultVal()
			arg.reset(&arg)
			subP[i] = arg
		}
	}
	p.compedArgs.leftRightRootTraversal(func(c *computedArg) error {
		c.reset()
		return nil
	})

	// compute the parsers new state
	if err := t.toArgValPairs(p).ToIter().ForEach(
		func(
			index int,
			val basic.Pair[*arg, string],
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

	// validate the parsers new state
	if err := p.checkRequiredArgsProvided(); err != nil {
		return customerr.AppendError(ParsingErr, err)
	}
	if err := p.checkConditionallyRequiredArgsProvided(); err != nil {
		return customerr.AppendError(ParsingErr, err)
	}

	// run all computer arguments to finalize state
	if err := p.compedArgs.leftRightRootTraversal(func(c *computedArg) error {
		return c.setVal()
	}); err != nil {
		return customerr.AppendError(ParsingErr, ComputedArgumentErr, err)
	}

	return nil
}

func (p *Parser) checkConditionalRequiredArgsExist() error {
	return p.longArgs.Vals().ForEach(
		func(index int, val *longArg) (iter.IteratorFeedback, error) {
			allConditionallyRequiredArgs := val.allConditionalArgs()
			for _, conditionalArg := range allConditionallyRequiredArgs {
				if _, err := p.longArgs.Get(conditionalArg); err != nil {
					return iter.Break, customerr.AppendError(
						ParserConfigErr,
						customerr.Wrap(
							UnrecognizedConditionallyRequiredArgErr,
							"Argument: %s", conditionalArg,
						),
					)
				}
			}
			return iter.Continue, nil
		},
	)
}

func (p *Parser) checkRequiredArgsProvided() error {
	missingRequiredArgs := []string{}
	p.requiredArgs.Vals().ForEach(
		func(index int, val *longArg) (iter.IteratorFeedback, error) {
			if !val.present {
				missingRequiredArgs = append(missingRequiredArgs, val.longFlag)
			}
			return iter.Continue, nil
		},
	)
	if len(missingRequiredArgs) > 0 {
		sort.Strings(missingRequiredArgs)
		return customerr.Wrap(
			MissingRequiredArgErr, "Missing: %s", missingRequiredArgs,
		)
	}
	return nil
}

func (p *Parser) checkConditionallyRequiredArgsProvided() error {
	// maps requiree to required
	missingRequiredArgs := map[string][]string{}
	p.longArgs.Vals().ForEach(
		func(index int, val *longArg) (iter.IteratorFeedback, error) {
			if !val.present && !val.defaultProvided {
				return iter.Continue, nil
			}
			for _, condArg := range val.conditionallyRequires() {
				if iterArg, _ := p.longArgs.Get(condArg); !iterArg.present {
					if _, ok := missingRequiredArgs[val.longFlag]; !ok {
						missingRequiredArgs[val.longFlag] = []string{}
					}
					missingRequiredArgs[val.longFlag] = append(
						missingRequiredArgs[val.longFlag],
						condArg,
					)
				}
			}
			return iter.Continue, nil
		},
	)

	if len(missingRequiredArgs) == 0 {
		return nil
	}

	keys := []string{}
	for k, _ := range missingRequiredArgs {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	errs := make([]error, len(keys))
	for i, k := range keys {
		errs[i] = customerr.Wrap(
			MissingConditionallyRequiredArgErr,
			"Given the value of '%s' the following args are required: %v",
			k, missingRequiredArgs[k],
		)
	}
	return customerr.AppendError(errs...)
}

const (
	shortValIdx int = iota
	longValIdx
	reqiredIdx
	defaultIdx
	condReqIdx
	descIdx
	reqMarking string = "(req.)"
)

// Returns a string representing the help menu.
func (p *Parser) Help() string {
	tableHeaders := [6]string{
		"", "", "",
		"Default Val", "Conditionally Reqs", "Description",
	}
	colWidths := [6]int{
		2,
		0,
		len(reqMarking),
		len(tableHeaders[defaultIdx]),
		len(tableHeaders[condReqIdx]),
		80,
	}

	// Sorts args alhpabetically
	args, _ := p.longArgs.Vals().Collect()
	sort.Slice(args, func(i, j int) bool {
		return args[i].longFlag < args[j].longFlag
	})

	table := make([][]string, p.longArgs.Length()+1)
	table[0] = tableHeaders[:]
	iter.SliceElems(args).ForEach(
		func(index int, val *longArg) (iter.IteratorFeedback, error) {
			conditionalArgs := val.allConditionalArgs()
			for i, v := range conditionalArgs {
				conditionalArgs[i] = "--" + v
			}

			if len(val.longFlag) > colWidths[longValIdx]+2 {
				colWidths[longValIdx] = len(val.longFlag) + 2
			}
			for _, condArg := range val.allConditionalArgs() {
				if len(condArg) > colWidths[condReqIdx]+2 {
					colWidths[condReqIdx] = len(condArg) + 2
				}
			}

			table[index+1] = make([]string, 6)
			if val.shortFlag != byte(0) {
				table[index+1][0] = fmt.Sprintf("-%c", val.shortFlag)
			}
			table[index+1][1] = fmt.Sprintf("--%s", val.longFlag)
			if val.required {
				table[index+1][2] = reqMarking
			}
			table[index+1][3], _ = val.defaultValAsStr()
			table[index+1][4] = strings.Join(conditionalArgs, " ")
			table[index+1][5] = val.description

			return iter.Continue, nil
		},
	)

	var sb strings.Builder
	sb.WriteString(p.progName)
	sb.WriteString(": HELP ME(nu)!\n")
	sb.WriteString("Description: ")
	sb.WriteString(p.progDesc)
	sb.WriteByte('\n')
	sb.WriteByte('\n')
	// Intentionally ignored err. Left for debugging purposes
	_ = strops.WriteTable(&sb, table, strops.WriteTableOpts{
		ColWidths:     colWidths[:],
		ColSeparators: []bool{false, false, false, true, true, true, true},
		RowSeparators: true,
	})
	return sb.String()
}
