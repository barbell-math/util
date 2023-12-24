package argparse

import "fmt"

type Options[T FromString[U], U interface{}] struct {
    ShortName rune
    Required bool
    Description string
    Default T
}

type Arg struct {
    Present bool
    Required bool
    Description string
    setVal func(a *Arg, args []string) error
}

type FromString[T interface{}] interface {
    Translate(args []string) (T,error)
    ToVal() T
}

func AddArg[T FromString[U], U interface{}](
    parser *Parser, 
    longName string, 
    opts *Options[T,U],
) *U {
    var rv *U
    err:=parser.addArg(&Arg{
        setVal: func(a *Arg, args []string) error {
            if a.Present {
                if v,err:=opts.Default.Translate(args); err==nil {
                    *rv=v
                } else {
                    return err
                }
            } else {
                *rv=opts.Default.ToVal()
            }
            return nil
        },
        Required: opts.Required,
        Description: opts.Description,
    },longName,opts.ShortName)
    if err!=nil {
        panic(fmt.Errorf("An error occured setting up the parser: %w",err))
    }
    return rv
}
