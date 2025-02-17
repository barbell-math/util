package examples

import (
	"fmt"
	"strings"

	"github.com/barbell-math/util/src/argparse"
)

// A translator that removes a given sequence from the supplied strings and
// counts how many times that sequence has been removed.
type SequenceCounter struct {
	Sequence string
	cntr     int
	Vals     []string
}

// The translate method. This is where all the work is done. (Ignore the poor
// time complexity, this is an example.)
// Note how the sequence counters internal state is updated here.
func (s *SequenceCounter) Translate(arg string) ([]string, error) {
	s.cntr += strings.Count(arg, s.Sequence)
	if newS := strings.ReplaceAll(arg, s.Sequence, ""); len(newS) > 0 {
		s.Vals = append(s.Vals, newS)
	}
	return s.Vals, nil
}

// The reset method. This translator actually has state that needs to be reset.
//
// The sequence field is not reset because the sequence counter still needs to
// look for the same sequence the next time the argparser uses the translator.
// This is an important detail: make sure that you know the difference between
// the state that needs to live for the entire length of the translator
// (potentially multiple argument sets being parsed) and the state that needs to
// live for the length of the current argument set (a single argument set being
// parsed).
func (s *SequenceCounter) Reset() {
	s.cntr = 0
	s.Vals = []string{}
}

func (s SequenceCounter) GetNumReplacements() int {
	return s.cntr
}

func Example_CustomTranslator_Stateful() {
	vals := struct {
		S []string
	}{}

	s := SequenceCounter{Sequence: "abc"}

	b := argparse.ArgBuilder{}
	// The call to SetTranslator is mandatory in this case because the sequence
	// counter translator has state that needs to be initialized.
	argparse.AddArg[*SequenceCounter](
		&vals.S, &b, "str",
		argparse.NewOpts[*SequenceCounter]().
			SetShortName('s').
			SetRequired(true).
			SetArgType(argparse.MultiValueArgType).
			SetTranslator(&s),
	)

	parser, err := b.ToParser("Prog name", "Prog description")
	fmt.Println("Parser error:", err)

	args := []string{"--str=abcdef", "ghiabcjkl", "mnopqrabc", "abc"}
	err = parser.Parse(argparse.ArgvIterFromSlice(args).ToTokens())
	fmt.Println("Parsing", args)
	fmt.Println(err)
	fmt.Println(vals.S)
	fmt.Println(s.GetNumReplacements())

	// Output:
	//Parser error: <nil>
	//Parsing [--str=abcdef ghiabcjkl mnopqrabc abc]
	//<nil>
	//[def ghijkl mnopqr]
	//4
}
