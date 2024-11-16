package examples

import (
	"fmt"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/computers"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/container/containers"
	"github.com/barbell-math/util/src/widgets"
)

func Example_ValueArgument_WithoutOptions() {
	vals := struct {
		I int
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[int, translators.BuiltinInt](&vals.I, &b, "int", nil)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--int=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	args = []string{"--int", "1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	// Output:
	//Parser error: <nil>
	//Parsing [--int=1]
	//<nil>
	//1
	//Parsing [--int 1]
	//<nil>
	//1
}

func Example_ValueArgument_WithOptions() {
	vals := struct {
		U uint
		I int
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[int, translators.BuiltinInt](
		&vals.I, &b, "int",
		argparse.NewOpts[int, translators.BuiltinInt]().
			SetShortName('i').
			SetRequired(true).
			SetDescription("This is an integer"),
	)
	argparse.AddArg[uint, translators.BuiltinUint](
		&vals.U, &b, "uint",
		argparse.NewOpts[uint, translators.BuiltinUint]().
			SetShortName('u').
			SetDefaultVal(3).
			SetDescription("This is an integer"),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--int=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)
	fmt.Println(vals.U)

	// Output:
	//Parser error: <nil>
	//Parsing [--int=1]
	//<nil>
	//1
	//3
}

func Example_FlagArgument() {
	vals := struct {
		B bool
	}{}

	b := argparse.ArgBuilder{}
	// Flags default value is "false"
	argparse.AddFlag(
		&vals.B, &b, "bool",
		argparse.NewOpts[bool, translators.Flag]().
			SetShortName('b').
			SetDescription("This is a flag argument"),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--bool"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.B)

	args = []string{}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.B)

	// Output:
	//Parser error: <nil>
	//Parsing [--bool]
	//<nil>
	//true
	//Parsing []
	//<nil>
	//false
}

func Example_FlagCounterArgument() {
	vals := struct {
		I int
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddFlagCntr[int](
		&vals.I, &b, "cntr",
		argparse.NewOpts[int, *translators.FlagCntr[int]]().
			SetShortName('c').
			SetDescription("This is a counter flag argument"),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--cntr"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	args = []string{"--cntr", "-c", "-ccc"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	// Output:
	//Parser error: <nil>
	//Parsing [--cntr]
	//<nil>
	//1
	//Parsing [--cntr -c -ccc]
	//<nil>
	//5
}

func Example_ListArgument() {
	vals := struct {
		L []int
	}{}

	b := argparse.ArgBuilder{}
	// The SetTranslator method must be called because the ListValues translator
	// has state that needs to be initialized.
	argparse.AddListArg[int, translators.BuiltinInt, widgets.BuiltinInt](
		&vals.L, &b, "list",
		argparse.NewOpts[
			[]int,
			*translators.ListValues[int, translators.BuiltinInt, widgets.BuiltinInt],
		]().
			SetShortName('l').
			SetDescription("This is a list flag argument").
			SetTranslator(&translators.ListValues[
				int,
				translators.BuiltinInt,
				widgets.BuiltinInt,
			]{
				ValueTranslator: translators.BuiltinInt{},
			}),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--list=1", "2", "3", "--list", "4", "-l", "5", "6"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.L)

	// Output:
	//Parser error: <nil>
	//Parsing [--list=1 2 3 --list 4 -l 5 6]
	//<nil>
	//[1 2 3 4 5 6]
}

func Example_ListArgument_AllowedValsSet() {
	vals := struct {
		L []int
	}{}

	b := argparse.ArgBuilder{}
	// The SetTranslator method must be called because the ListValues translator
	// has state that needs to be initialized.
	argparse.AddListArg[int, translators.BuiltinInt, widgets.BuiltinInt](
		&vals.L, &b, "list",
		argparse.NewOpts[
			[]int,
			*translators.ListValues[int, translators.BuiltinInt, widgets.BuiltinInt],
		]().
			SetShortName('l').
			SetDescription("This is a list flag argument").
			SetTranslator(&translators.ListValues[
				int,
				translators.BuiltinInt,
				widgets.BuiltinInt,
			]{
				ValueTranslator: translators.BuiltinInt{},
				AllowedVals: containers.HashSetValInit[int, widgets.BuiltinInt](
					1, 2, 3, 4, 5, 6,
				),
			}),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	// The list argument can be specified many times, each time with one or more
	// values. All values will be added to the same list.
	args := []string{"--list=1", "2", "3", "--list", "4", "-l", "5", "6"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.L)

	// Output:
	//Parser error: <nil>
	//Parsing [--list=1 2 3 --list 4 -l 5 6]
	//<nil>
	//[1 2 3 4 5 6]
}

func Example_SelectorArgument() {
	vals := struct {
		I int
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddSelector[int, translators.BuiltinInt, widgets.BuiltinInt](
		&vals.I, &b, "selector",
		argparse.NewOpts[
			int,
			translators.Selector[int, translators.BuiltinInt, widgets.BuiltinInt],
		]().
			SetShortName('s').
			SetDescription("This is a selector argument").
			SetTranslator(translators.Selector[
				int,
				translators.BuiltinInt,
				widgets.BuiltinInt,
			]{
				ValueTranslator: translators.BuiltinInt{},
				AllowedVals: containers.HashSetValInit[int, widgets.BuiltinInt](
					1, 2, 3, 4, 5,
				),
			}),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--selector=4"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I)

	// Output:
	//Parser error: <nil>
	//Parsing [--selector=4]
	//<nil>
	//4
}

func Example_ComputedArgument() {
	vals := struct {
		I1  int
		I2  int
		Res int
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[int, translators.BuiltinInt](
		&vals.I1, &b, "int1",
		argparse.NewOpts[int, translators.BuiltinInt]().
			SetShortName('1').
			SetRequired(true).
			SetDescription("This is an integer"),
	)
	argparse.AddArg[int, translators.BuiltinInt](
		&vals.I2, &b, "int2",
		argparse.NewOpts[int, translators.BuiltinInt]().
			SetShortName('2').
			SetDefaultVal(3).
			SetDescription("This is an integer"),
	)
	argparse.AddComputedArg[int, computers.Add[int]](
		&vals.Res, &b,
		computers.Add[int]{L: &vals.I1, R: &vals.I2},
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--int1=1", "--int2=5"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I1)
	fmt.Println(vals.I2)
	fmt.Println(vals.Res)

	args = []string{"--int1=1"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.I1)
	fmt.Println(vals.I2)
	fmt.Println(vals.Res)

	// Output:
	//Parser error: <nil>
	//Parsing [--int1=1 --int2=5]
	//<nil>
	//1
	//5
	//6
	//Parsing [--int1=1]
	//<nil>
	//1
	//3
	//4
}
