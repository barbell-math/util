package argparse

import "errors"

var (
	ParserConfigErr                         = errors.New("An error occurred setting up the parser")
	ReservedShortNameErr                    = errors.New("Reserved short name used")
	ReservedLongNameErr                     = errors.New("Reserved long name used")
	DuplicateShortNameErr                   = errors.New("Duplicate short name")
	DuplicateLongNameErr                    = errors.New("Duplicate long name")
	LongNameToShortErr                      = errors.New("Long name must be more than one char")
	UnrecognizedConditionallyRequiredArgErr = errors.New("Unrecognized conditionally required argument")

	ParserCombinationErr = errors.New("Could not combine parsers")

	ParsingErr                     = errors.New("An error occurred parsing the supplied arguments")
	ExpectedArgumentErr            = errors.New("Expected an argument (short or long)")
	ExpectedValueErr               = errors.New("Expected a value")
	UnrecognizedShortArgErr        = errors.New("Unrecognized short argument")
	UnrecognizedLongArgErr         = errors.New("Unrecognized long argument")
	EndOfTokenStreamErr            = errors.New("The end of the token stream was reached")
	ArgumentPassedMultipleTimesErr = errors.New("Argument was passed multiple times but was expected only once")

	ParserConfigFileErr       = errors.New("An error occurred parsing a parser config file")
	ParserConfigFileSyntaxErr = errors.New("Syntax error")

	ArgumentTranslationErr             = errors.New("An error occurred translating the supplied argument")
	MissingRequiredArgErr              = errors.New("Required argument(s) missing")
	MissingConditionallyRequiredArgErr = errors.New("Conditionally required argument(s) missing")
	ComputedArgumentErr                = errors.New("An error occurred calculating a computed argument")

	// The error returned when the help menu is displayed, indicating that the
	// parsing the arguments did not end in a "true" error but also did not
	// completely finish.
	HelpErr = errors.New("Help flag specified. Stopping.")
)
