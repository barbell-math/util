package reflect

// Code generated by ../bin/structDefaultInit - DO NOT EDIT.
import ()

// Returns a new structHashOpts struct initialized with the default values.
func NewStructHashOpts() *structHashOpts {
	return &structHashOpts{
		optionsFlag: NewOptionsFlag(),
	}
}

func (o *structHashOpts) SetOptionsFlag(v optionsFlag) *structHashOpts {
	o.optionsFlag = v
	return o
}
