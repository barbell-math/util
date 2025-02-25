package generatortests

// Code generated by ../../bin/enum - DO NOT EDIT.
import (
	"errors"
	"fmt"
)

var (
	InvalidTokenType             = errors.New("Invalid tokenType")
	TOKEN_TYPE       []tokenType = []tokenType{
		unknownTokenType,
		shortFlagToken,
		longFlagToken,
		valueToken,
	}
)

func NewTokenType() tokenType {
	return unknownTokenType
}

func (o tokenType) Value() tokenType {
	return o
}

func (o tokenType) Valid() error {
	switch o {

	case unknownTokenType:
		return nil

	case shortFlagToken:
		return nil

	case longFlagToken:
		return nil

	case valueToken:
		return nil

	default:
		return InvalidTokenType
	}
}

func (o tokenType) String() string {
	switch o {
	case unknownTokenType:
		return "unknownTokenType"
	case shortFlagToken:
		return "shortFlagToken"
	case longFlagToken:
		return "longFlagToken"
	case valueToken:
		return "valueToken"

	default:
		return "unknownTokenType"
	}
}

func (o tokenType) MarshalJSON() ([]byte, error) {
	switch o {

	case unknownTokenType:
		return []byte("unknownTokenType"), nil

	case shortFlagToken:
		return []byte("shortFlagToken"), nil

	case longFlagToken:
		return []byte("longFlagToken"), nil

	case valueToken:
		return []byte("valueToken"), nil

	default:
		return []byte("unknownTokenType"), InvalidTokenType
	}
}

func (o *tokenType) FromString(s string) error {
	switch s {

	case "unknownTokenType":
		*o = unknownTokenType
		return nil

	case "shortFlagToken":
		*o = shortFlagToken
		return nil

	case "longFlagToken":
		*o = longFlagToken
		return nil

	case "valueToken":
		*o = valueToken
		return nil

	default:
		*o = unknownTokenType
		return fmt.Errorf("%w: %s", InvalidTokenType, s)
	}
}

func (o *tokenType) UnmarshalJSON(b []byte) error {
	switch string(b) {

	case "unknownTokenType":
		*o = unknownTokenType
		return nil

	case "shortFlagToken":
		*o = shortFlagToken
		return nil

	case "longFlagToken":
		*o = longFlagToken
		return nil

	case "valueToken":
		*o = valueToken
		return nil

	default:
		*o = unknownTokenType
		return fmt.Errorf("%w: %s", InvalidTokenType, string(b))
	}
}
