package examples

import (
	"fmt"

	"github.com/barbell-math/util/src/argparse"
)

// A translator that reverses any strings it receives.
type Reverser struct{}

// The translate method. This is where all the work is done.
// Note the method receiver is '_' denoting that the value is stateless.
func (_ Reverser) Translate(arg string) (string, error) {
	rv := []byte(arg)
	for i := 0; i < len(arg); i++ {
		rv[i] = arg[len(arg)-i-1]
	}
	return string(rv), nil
}

// The reset method. If this were a stateful translator then the translators
// state would be reset here, but it is not so it there is nothing to do. The
// method still has to exist to satisfy the Translator interface.
func (_ Reverser) Reset() {
	// Intentional noop - reverser has no state to reset!
}

func Example_CustomTranslator_Stateless() {
	vals := struct {
		S string
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[Reverser](
		&vals.S, &b, "str",
		argparse.NewOpts[Reverser]().
			SetShortName('s').
			SetRequired(true),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--str=abcdefghijklmnopqrstuvwxyz"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.S)

	// Output:
	//Parser error: <nil>
	//Parsing [--str=abcdefghijklmnopqrstuvwxyz]
	//<nil>
	//zyxwvutsrqponmlkjihgfedcba
}
