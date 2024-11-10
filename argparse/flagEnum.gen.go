package argparse

// Code generated by ../bin/enum - DO NOT EDIT.
import (
	"errors"
	"fmt"
)

var (
	InvalidFlag        = errors.New("Invalid flag")
	FLAG        []flag = []flag{
		unknownFlag,
		shortSpaceFlag,
		shortEqualsFlag,
		longSpaceFlag,
		longEqualsFlag,
	}
)

func NewFlag() flag {
	return unknownFlag
}

func (o flag) Valid() error {
	switch o {

	case unknownFlag:
		return nil

	case shortSpaceFlag:
		return nil

	case shortEqualsFlag:
		return nil

	case longSpaceFlag:
		return nil

	case longEqualsFlag:
		return nil

	default:
		return InvalidFlag
	}
}

func (o flag) String() string {
	switch o {
	case unknownFlag:
		return "unknownFlag"
	case shortSpaceFlag:
		return "shortSpaceFlag"
	case shortEqualsFlag:
		return "shortEqualsFlag"
	case longSpaceFlag:
		return "longSpaceFlag"
	case longEqualsFlag:
		return "longEqualsFlag"

	default:
		return "unknownFlag"
	}
}

func (o flag) MarshalJSON() ([]byte, error) {
	switch o {

	case unknownFlag:
		return []byte("unknownFlag"), nil

	case shortSpaceFlag:
		return []byte("shortSpaceFlag"), nil

	case shortEqualsFlag:
		return []byte("shortEqualsFlag"), nil

	case longSpaceFlag:
		return []byte("longSpaceFlag"), nil

	case longEqualsFlag:
		return []byte("longEqualsFlag"), nil

	default:
		return []byte("unknownFlag"), InvalidFlag
	}
}

func (o *flag) FromString(s string) error {
	switch s {

	case "unknownFlag":
		*o = unknownFlag
		return nil

	case "shortSpaceFlag":
		*o = shortSpaceFlag
		return nil

	case "shortEqualsFlag":
		*o = shortEqualsFlag
		return nil

	case "longSpaceFlag":
		*o = longSpaceFlag
		return nil

	case "longEqualsFlag":
		*o = longEqualsFlag
		return nil

	default:
		*o = unknownFlag
		return fmt.Errorf("%w: %s", InvalidFlag, s)
	}
}

func (o *flag) UnmarshalJSON(b []byte) error {
	switch string(b) {

	case "unknownFlag":
		*o = unknownFlag
		return nil

	case "shortSpaceFlag":
		*o = shortSpaceFlag
		return nil

	case "shortEqualsFlag":
		*o = shortEqualsFlag
		return nil

	case "longSpaceFlag":
		*o = longSpaceFlag
		return nil

	case "longEqualsFlag":
		*o = longEqualsFlag
		return nil

	default:
		*o = unknownFlag
		return fmt.Errorf("%w: %s", InvalidFlag, string(b))
	}
}
