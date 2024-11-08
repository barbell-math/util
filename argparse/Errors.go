package argparse

import "errors"

var (
	ParserConfigErr = errors.New("An error occurred setting up the parser")
	ParsingErr = errors.New("An error occurred parsing the supplied arguments")
	UnrecognizedFlag = errors.New("Unrecognized argument")
)
