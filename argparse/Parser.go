package argparse

import (
	"fmt"

	"github.com/barbell-math/util/iter"
)

type (
	Options[T any, U ArgumentTranslation[T]] struct {
		ShortName byte
    	Required bool
    	Description string
    	Default U
	}

	ArgumentTranslation[T any] interface {
		Translate(arg string) (T, error)
		Default() T
	}

	Parser struct {
		args []Arg
		shortFlags map[byte]*Arg
		longFlags map[string]*Arg
	}
	Arg struct {
		setVal func(a *Arg, arg string) error
		shortFlag byte
		longFlag string
		argType ArgType
		present bool
		required bool
		description string
	}
)

func AddArg[T any, U ArgumentTranslation[T]](
    parser *Parser, 
    longName string, 
	_type ArgType,
    opts *Options[T,U],
) *T {
    var rv *T
    if err:=parser.addArg(
		Arg{
			setVal: func(a *Arg, arg string) error {
        	    if a.present {
        	        if v,err:=opts.Default.Translate(arg); err==nil {
        	            *rv=v
        	        } else {
        	            return err
        	        }
        	    } else {
        	        *rv=opts.Default.Default()
        	    }
        	    return nil
        	},
        	required: opts.Required,
        	description: opts.Description,
			shortFlag: opts.ShortName,
			longFlag: longName,
		},
	); err!=nil {
        panic(fmt.Errorf("%w: %w", ParserConfigErr, err))
    }
    return rv
}

func (p *Parser) addArg(arg Arg) error {
	if _, ok:=p.longFlags[arg.longFlag]; ok {
		return fmt.Errorf(
			"%w: the long name '%s' is has already been specified",
			ParsingErr, arg.longFlag,
		)
	}

	if _, ok:=p.shortFlags[arg.shortFlag]; ok {
		return fmt.Errorf(
			"%w: the short name '%c' is has already been specified",
			ParsingErr, arg.shortFlag,
		)
	}

	p.args = append(p.args, arg)
	// Cannot populate name maps with arg structs until all args have been added
	// due to the underlying slice reallocating
	p.shortFlags[arg.shortFlag]=nil
	p.longFlags[arg.longFlag]=nil
	return nil
}

func (p Parser) Parse(a ArgvIter) {
	for i, _:=range(p.args) {
		p.shortFlags[p.args[i].shortFlag]=&p.args[i]
		p.longFlags[p.args[i].longFlag]=&p.args[i]
	}

	a.ToTokens().ToIter().ForEach(
		func(index int, val token) (iter.IteratorFeedback, error) {
			return iter.Continue, nil
		},
	)
}
