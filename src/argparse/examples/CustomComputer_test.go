package examples

import (
	"errors"
	"fmt"
	"strings"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
)

// A computer that interleaves two strings that is has pointers to.
// In this example, S1 and S2 will be pointers to the values received from the
// cmd line interface that were translated and saved.
type StrWeaver struct {
	S1 *string
	S2 *string
}

// The compute function. This is where all the work is done.
// The resulting computed value should be returned from this function.
func (s StrWeaver) ComputeVals() (string, error) {
	if len(*s.S1) != len(*s.S2) {
		return "", errors.New("Strings must be the same length.")
	}

	var sb strings.Builder
	for i := 0; i < len(*s.S1); i++ {
		sb.WriteByte((*s.S1)[i])
		sb.WriteByte((*s.S2)[i])
	}
	return sb.String(), nil
}

// The reset method. If the computer had any state that needs to be reset then
// the actions to do that would be placed here. The method has to exist to
// satisfy the Computer interface.
func (s StrWeaver) Reset() {
	// Intentional noop - reverser has no state to reset!
}

func Example_CustomComputer() {
	vals := struct {
		S1  string
		S2  string
		Res string
	}{}

	b := argparse.ArgBuilder{}
	argparse.AddArg[translators.BuiltinString](
		&vals.S1, &b, "str1",
		argparse.NewOpts[translators.BuiltinString]().
			SetShortName('1').
			SetRequired(true),
	)
	argparse.AddArg[translators.BuiltinString](
		&vals.S2, &b, "str2",
		argparse.NewOpts[translators.BuiltinString]().
			SetShortName('2').
			SetRequired(true),
	)
	argparse.AddComputedArg[StrWeaver](
		&vals.Res, &b, StrWeaver{
			S1: &vals.S1,
			S2: &vals.S2,
		},
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--str1=abc", "--str2=def"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.S1)
	fmt.Println(vals.S2)
	fmt.Println(vals.Res)

	// Output:
	//Parser error: <nil>
	//Parsing [--str1=abc --str2=def]
	//<nil>
	//abc
	//def
	//adbecf
}
