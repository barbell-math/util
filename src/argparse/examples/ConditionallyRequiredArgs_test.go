package examples

import (
	"fmt"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
)

func Example_ConditionallyRequiredArgument_BasedOnExistance() {
	vals := struct {
		U uint
		I int
		F float32
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&vals.I, &b, "int",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('i').
			SetDescription("This is an integer"),
	)
	argparse.AddArg[translators.BuiltinFloat32](
		&vals.F, &b, "float",
		argparse.NewOpts[translators.BuiltinFloat32]().
			SetShortName('f').
			SetDescription("This is a float"),
	)
	argparse.AddArg[translators.BuiltinUint](
		&vals.U, &b, "uint",
		argparse.NewOpts[translators.BuiltinUint]().
			SetShortName('u').
			SetDescription("This is an unsigned integer").
			// The ArgConditionallity struct specifies that if the uint argument
			// was supplied at all then the int and float arguments must have
			// also been supplied.
			SetConditionallyRequired([]argparse.ArgConditionality[uint]{
				argparse.ArgConditionality[uint]{
					Requires: []string{"int", "float"},
					When:     argparse.ArgSupplied[uint],
				},
			}),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--int=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("uint not supplied:")
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)
	fmt.Println(vals.U)
	fmt.Println(vals.F)

	args = []string{"--uint=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("uint supplied without conditionally required arguments:")
	fmt.Println("Parsing", args)
	fmt.Println(err)

	args = []string{"--uint=1", "--float=3.14", "--int=3"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("uint supplied with conditionally required arguments:")
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)
	fmt.Println(vals.U)
	fmt.Println(vals.F)

	// Output:
	//Parser error: <nil>
	//uint not supplied:
	//Parsing [--int=1]
	//<nil>
	//1
	//0
	//0
	//uint supplied without conditionally required arguments:
	//Parsing [--uint=1]
	//An error occurred parsing the supplied arguments
	//Conditionally required argument(s) missing
	//   |- Given the value of 'uint' the following args are required: [int float]
	//uint supplied with conditionally required arguments:
	//Parsing [--uint=1 --float=3.14 --int=3]
	//<nil>
	//3
	//1
	//3.14
}

func Example_ConditionallyRequiredArgument_BasedOnValue() {
	vals := struct {
		U uint
		I int
		F float32
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&vals.I, &b, "int",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('i').
			SetDescription("This is an integer"),
	)
	argparse.AddArg[translators.BuiltinFloat32](
		&vals.F, &b, "float",
		argparse.NewOpts[translators.BuiltinFloat32]().
			SetShortName('f').
			SetDescription("This is a float"),
	)
	argparse.AddArg[translators.BuiltinUint](
		&vals.U, &b, "uint",
		argparse.NewOpts[translators.BuiltinUint]().
			SetShortName('u').
			SetDefaultVal(3).
			SetDescription("This is an unsigned integer").
			// The ArgConditionallity struct specifies that if the uint argument
			// was supplied *and its supplied value is 1* then the int and float
			// arguments must have also been supplied.
			SetConditionallyRequired([]argparse.ArgConditionality[uint]{
				argparse.ArgConditionality[uint]{
					Requires: []string{"int", "float"},
					When:     argparse.ArgEquals[uint](1),
				},
			}),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--int=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("uint not supplied:")
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)
	fmt.Println(vals.U)
	fmt.Println(vals.F)

	args = []string{"--uint=0"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("uint supplied but with different value:")
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)
	fmt.Println(vals.U)
	fmt.Println(vals.F)

	args = []string{"--uint=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("uint supplied without conditionally required arguments:")
	fmt.Println("Parsing", args)
	fmt.Println(err)

	args = []string{"--uint=1", "--float=3.14", "--int=3"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("uint supplied with conditionally required arguments:")
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)
	fmt.Println(vals.U)
	fmt.Println(vals.F)

	// Output:
	//Parser error: <nil>
	//uint not supplied:
	//Parsing [--int=1]
	//<nil>
	//1
	//3
	//0
	//uint supplied but with different value:
	//Parsing [--uint=0]
	//<nil>
	//0
	//0
	//0
	//uint supplied without conditionally required arguments:
	//Parsing [--uint=1]
	//An error occurred parsing the supplied arguments
	//Conditionally required argument(s) missing
	//   |- Given the value of 'uint' the following args are required: [int float]
	//uint supplied with conditionally required arguments:
	//Parsing [--uint=1 --float=3.14 --int=3]
	//<nil>
	//3
	//1
	//3.14
}
