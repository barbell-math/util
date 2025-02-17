package examples

import (
	"fmt"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/computers"
	"github.com/barbell-math/util/src/argparse/translators"
)

func Example_SubParserSimple() {
	vals := struct {
		I int
		B bool
	}{}

	b1 := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&vals.I, &b1, "int",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('i').
			SetDescription("This is an integer"),
	)
	subParser, err := b1.ToParser("", "")
	fmt.Println("Sub Parser error:", err)

	b2 := argparse.ArgBuilder{}
	argparse.AddFlag(
		&vals.B, &b2, "bool",
		argparse.NewOpts[translators.Flag]().
			SetShortName('b').
			SetDescription("This is a flag argument"),
	)
	parser, err := b2.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)
	err = parser.AddSubParsers(subParser)
	fmt.Println("Add Sub Parser error:", err)

	args := []string{"--bool", "-i=3"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.B)
	fmt.Println(vals.I)

	// Output:
	//Sub Parser error: <nil>
	//Parser error: <nil>
	//Add Sub Parser error: <nil>
	//Parsing [--bool -i=3]
	//<nil>
	//true
	//3
}

func Example_SubParserComplex() {
	// This addmitidly contrived example creates a set of subparsers that when
	// executed evauluate the expression: (A+B)*(C-D)
	// In reality something like this should probably never be used, but it
	// demonstrates that the subparsers can be combined to evaulate arbitrary
	// expressions. This means computed arguments can rely on values that are
	// the result of other computed arguments as long as the subparser tree is
	// properly constructed.
	res := struct {
		A int
		B int
		C int
		D int
		E int
		F int
		G int
	}{}

	b1 := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&res.A,
		&b1,
		"aa",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('a').
			SetRequired(true),
	)
	p1, _ := b1.ToParser("", "")

	b2 := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&res.B,
		&b2,
		"bb",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('b').
			SetRequired(true),
	)
	p2, _ := b2.ToParser("", "")

	b3 := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&res.C,
		&b3,
		"cc",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('c').
			SetRequired(true),
	)
	p3, _ := b3.ToParser("", "")

	b4 := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinInt](
		&res.D,
		&b4,
		"dd",
		argparse.NewOpts[translators.BuiltinInt]().
			SetShortName('d').
			SetRequired(true),
	)
	p4, _ := b4.ToParser("", "")

	b5 := argparse.ArgBuilder{}
	argparse.AddComputedArg[computers.Add[int]](
		&res.E,
		&b5,
		computers.Add[int]{L: &res.A, R: &res.B},
	)
	p5, _ := b5.ToParser("", "")

	b6 := argparse.ArgBuilder{}
	argparse.AddComputedArg[computers.Sub[int]](
		&res.F,
		&b6,
		computers.Sub[int]{L: &res.C, R: &res.D},
	)
	p6, _ := b6.ToParser("", "")

	b7 := argparse.ArgBuilder{}
	argparse.AddComputedArg[computers.Mul[int]](
		&res.G,
		&b7,
		computers.Mul[int]{L: &res.E, R: &res.F},
	)
	p7, _ := b7.ToParser("", "")

	p5.AddSubParsers(p1, p2)
	p6.AddSubParsers(p3, p4)
	p7.AddSubParsers(p5, p6)

	args := []string{
		"-a=3",
		"-b=5",
		"-c=7",
		"-d=9",
	}
	err := p7.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(res.A)
	fmt.Println(res.B)
	fmt.Println(res.C)
	fmt.Println(res.D)
	fmt.Println(res.E)
	fmt.Println(res.F)
	fmt.Println(res.G)

	// Output:
	//Parsing [-a=3 -b=5 -c=7 -d=9]
	//<nil>
	//3
	//5
	//7
	//9
	//8
	//-2
	//-16
}
