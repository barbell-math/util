package testenum

// Code generated by ../../../bin/enum - DO NOT EDIT.
import (
	"errors"
	"fmt"
)

var (
	InvalidTestEnum            = errors.New("Invalid TestEnum")
	TEST_ENUM       []TestEnum = []TestEnum{
		UnknownAppAction,
		AppActionOne,
		AppActionTwo,
	}
)

func NewTestEnum() TestEnum {
	return UnknownAppAction
}

func (o TestEnum) Value() TestEnum {
	return o
}

func (o TestEnum) Valid() error {
	switch o {

	case UnknownAppAction:
		return nil

	case AppActionOne:
		return nil

	case AppActionTwo:
		return nil

	default:
		return InvalidTestEnum
	}
}

func (o TestEnum) String() string {
	switch o {
	case UnknownAppAction:
		return "UnknownAppAction"
	case AppActionOne:
		return "AppActionOne"
	case AppActionTwo:
		return "AppActionTwo"

	default:
		return "UnknownAppAction"
	}
}

func (o TestEnum) MarshalJSON() ([]byte, error) {
	switch o {

	case UnknownAppAction:
		return []byte("UnknownAppAction"), nil

	case AppActionOne:
		return []byte("AppActionOne"), nil

	case AppActionTwo:
		return []byte("AppActionTwo"), nil

	default:
		return []byte("UnknownAppAction"), InvalidTestEnum
	}
}

func (o *TestEnum) FromString(s string) error {
	switch s {

	case "UnknownAppAction":
		*o = UnknownAppAction
		return nil

	case "AppActionOne":
		*o = AppActionOne
		return nil

	case "AppActionTwo":
		*o = AppActionTwo
		return nil

	default:
		*o = UnknownAppAction
		return fmt.Errorf("%w: %s", InvalidTestEnum, s)
	}
}

func (o *TestEnum) UnmarshalJSON(b []byte) error {
	switch string(b) {

	case "UnknownAppAction":
		*o = UnknownAppAction
		return nil

	case "AppActionOne":
		*o = AppActionOne
		return nil

	case "AppActionTwo":
		*o = AppActionTwo
		return nil

	default:
		*o = UnknownAppAction
		return fmt.Errorf("%w: %s", InvalidTestEnum, string(b))
	}
}
