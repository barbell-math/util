package enum

import "encoding/json"

type (
	Enum interface {
		json.Marshaler
		json.Unmarshaler
		Valid() bool
		String() string
		FromString(s string) error
	}
)
