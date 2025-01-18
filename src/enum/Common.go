package enum

import "encoding/json"

type (
	Pntr[E Value] interface {
		*E
		writeable
	}

	writeable interface {
		json.Unmarshaler
		FromString(s string) error
	}

	Value interface {
		json.Marshaler
		Valid() error
		String() string
	}
)
