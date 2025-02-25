package reflect

// Code generated by ../../bin/enum - DO NOT EDIT.
import (
	"errors"
	"fmt"
)

var (
	InvalidOptionsFlag               = errors.New("Invalid optionsFlag")
	OPTIONS_FLAG       []optionsFlag = []optionsFlag{
		followPntrs,
		followInterfaces,
		recurseStructs,
		includeMapVals,
		includeSliceVals,
		includeArrayVals,
		unknownOptionsFlag,
	}
)

func NewOptionsFlag() optionsFlag {
	return includeMapVals | includeArrayVals | includeSliceVals | followPntrs | followInterfaces | recurseStructs
}

func (o optionsFlag) Value() optionsFlag {
	return o
}

func (o optionsFlag) Valid() error {
	switch o {

	case followPntrs:
		return nil

	case followInterfaces:
		return nil

	case recurseStructs:
		return nil

	case includeMapVals:
		return nil

	case includeSliceVals:
		return nil

	case includeArrayVals:
		return nil

	case unknownOptionsFlag:
		return nil

	default:
		return InvalidOptionsFlag
	}
}

func (o optionsFlag) String() string {
	switch o {
	case followPntrs:
		return "followPntrs"
	case followInterfaces:
		return "followInterfaces"
	case recurseStructs:
		return "recurseStructs"
	case includeMapVals:
		return "includeMapVals"
	case includeSliceVals:
		return "includeSliceVals"
	case includeArrayVals:
		return "includeArrayVals"
	case unknownOptionsFlag:
		return "unknownOptionsFlag"

	default:
		return "unknownOptionsFlag"
	}
}

func (o optionsFlag) MarshalJSON() ([]byte, error) {
	switch o {

	case followPntrs:
		return []byte("followPntrs"), nil

	case followInterfaces:
		return []byte("followInterfaces"), nil

	case recurseStructs:
		return []byte("recurseStructs"), nil

	case includeMapVals:
		return []byte("includeMapVals"), nil

	case includeSliceVals:
		return []byte("includeSliceVals"), nil

	case includeArrayVals:
		return []byte("includeArrayVals"), nil

	case unknownOptionsFlag:
		return []byte("unknownOptionsFlag"), nil

	default:
		return []byte("unknownOptionsFlag"), InvalidOptionsFlag
	}
}

func (o *optionsFlag) FromString(s string) error {
	switch s {

	case "followPntrs":
		*o = followPntrs
		return nil

	case "followInterfaces":
		*o = followInterfaces
		return nil

	case "recurseStructs":
		*o = recurseStructs
		return nil

	case "includeMapVals":
		*o = includeMapVals
		return nil

	case "includeSliceVals":
		*o = includeSliceVals
		return nil

	case "includeArrayVals":
		*o = includeArrayVals
		return nil

	case "unknownOptionsFlag":
		*o = unknownOptionsFlag
		return nil

	default:
		*o = unknownOptionsFlag
		return fmt.Errorf("%w: %s", InvalidOptionsFlag, s)
	}
}

func (o *optionsFlag) UnmarshalJSON(b []byte) error {
	switch string(b) {

	case "followPntrs":
		*o = followPntrs
		return nil

	case "followInterfaces":
		*o = followInterfaces
		return nil

	case "recurseStructs":
		*o = recurseStructs
		return nil

	case "includeMapVals":
		*o = includeMapVals
		return nil

	case "includeSliceVals":
		*o = includeSliceVals
		return nil

	case "includeArrayVals":
		*o = includeArrayVals
		return nil

	case "unknownOptionsFlag":
		*o = unknownOptionsFlag
		return nil

	default:
		*o = unknownOptionsFlag
		return fmt.Errorf("%w: %s", InvalidOptionsFlag, string(b))
	}
}
