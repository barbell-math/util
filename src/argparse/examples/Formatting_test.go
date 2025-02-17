package examples

import (
	"fmt"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
)

func Example_LongArgFormats() {
	vals := struct {
		I int
		B bool
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&vals.I, &b, "int",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('i').
			SetDescription("This is an integer"),
	)
	argparse.AddFlag(
		&vals.B, &b, "bool",
		argparse.NewOpts[translators.Flag]().
			SetShortName('b').
			SetDescription("This is a flag argument"),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	// Long args with no associated values
	args := []string{"--bool"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.B)

	// Long args with associated value after equals sign
	args = []string{"--int=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	// Long args with associated value after space
	args = []string{"--int", "1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	// Output:
	//Parser error: <nil>
	//Parsing [--bool]
	//<nil>
	//true
	//Parsing [--int=1]
	//<nil>
	//1
	//Parsing [--int 1]
	//<nil>
	//1
}

func Example_ShortArgFormats() {
	vals := struct {
		I  int
		B1 bool
		B2 bool
		B3 bool
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&vals.I, &b, "int",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('i').
			SetDescription("This is an integer"),
	)
	argparse.AddFlag(
		&vals.B1, &b, "b1",
		argparse.NewOpts[translators.Flag]().
			SetShortName('1').
			SetDescription("This is a flag argument"),
	)
	argparse.AddFlag(
		&vals.B2, &b, "b2",
		argparse.NewOpts[translators.Flag]().
			SetShortName('2').
			SetDescription("This is a flag argument"),
	)
	argparse.AddFlag(
		&vals.B3, &b, "b3",
		argparse.NewOpts[translators.Flag]().
			SetShortName('3').
			SetDescription("This is a flag argument"),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	// Short args with no associated values
	args := []string{"-1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.B1)

	// Short args with no associated values can be combined
	// If a short arg has a value associated with it, it must be separate
	args = []string{"-123"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.B1)
	fmt.Println(vals.B2)
	fmt.Println(vals.B3)

	// Short args with associated value after equals sign
	args = []string{"-i=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	// Short args with associated value after space
	args = []string{"-i", "1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	// Output:
	//Parser error: <nil>
	//Parsing [-1]
	//<nil>
	//true
	//Parsing [-123]
	//<nil>
	//true
	//true
	//true
	//Parsing [-i=1]
	//<nil>
	//1
	//Parsing [-i 1]
	//<nil>
	//1
}
