package reflect

import (
	"reflect"

	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/widgets"
)

//go:generate ../bin/enum -type=optionsFlag -package=reflect
//go:generate ../bin/flags -type=optionsFlag -package=reflect
//go:generate ../bin/structDefaultInit -struct=structHashOpts

type (
	//gen:enum unknownValue unknownOptionsFlag
	//gen:enum default includeMapVals | includeArrayVals | includeSliceVals | followPntrs
	optionsFlag    int
	structHashOpts struct {
		optionsFlag `default:"NewOptionsFlag()" setter:"f" getter:"f"`
	}
)

const (
	// Description: set to true if the hash value should be calculated by
	// following pointer values rather than using the pointers value itself
	//
	// Used by: [ToStructs]
	//
	// Default: true
	//gen:enum string followPntrs
	followPntrs optionsFlag = 1 << iota
	// Description: set to true if the hash value should be calculated by
	// following interface values rather than using the interface value itself
	//
	// Used by: [ToStructs]
	//
	// Default: true
	//gen:enum string followInterface
	followInterface
	// Description: set to true to include map key value pairs in the hash
	// calculation
	//
	// Used by: [ToStructs]
	//
	// Default: true
	//gen:enum string includeMapVals
	includeMapVals
	// Description: set to true to include slice values in the hash calculation
	//
	// Used by: [ToStructs]
	//
	// Default: true
	//gen:enum string includeSliceVals
	includeSliceVals
	// Description: set to true to include array values in the hash
	// calculation
	//
	// Used by: [ToStructs]
	//
	// Default: true
	//gen:enum string includeArrayVals
	includeArrayVals
	//gen:flags noSetter
	//gen:enum string unknownOptionsFlag
	unknownOptionsFlag
)

func getHash(val FieldInfo, opts *structHashOpts) hash.Hash {
	var rv hash.Hash
	switch val.Kind {
	case reflect.Bool:
		bw := widgets.BuiltinBool{}
		if iterV, err := val.Pntr(); err != nil {
			rv = bw.Hash(iterV.(*bool)).Combine(rv)
		}
	case reflect.Int:
		iw := widgets.BuiltinInt{}
		if iterV, err := val.Pntr(); err != nil {
			rv = iw.Hash(iterV.(*int)).Combine(rv)
		}
	case reflect.Int8:
		i8w := widgets.BuiltinInt8{}
		if iterV, err := val.Pntr(); err != nil {
			rv = i8w.Hash(iterV.(*int8)).Combine(rv)
		}
	case reflect.Int16:
		i16w := widgets.BuiltinInt16{}
		if iterV, err := val.Pntr(); err != nil {
			rv = i16w.Hash(iterV.(*int16)).Combine(rv)
		}
	case reflect.Int32:
		i32w := widgets.BuiltinInt32{}
		if iterV, err := val.Pntr(); err != nil {
			rv = i32w.Hash(iterV.(*int32)).Combine(rv)
		}
	case reflect.Int64:
		i64w := widgets.BuiltinInt64{}
		if iterV, err := val.Pntr(); err != nil {
			rv = i64w.Hash(iterV.(*int64)).Combine(rv)
		}
	case reflect.Uint:
		uw := widgets.BuiltinUint{}
		if iterV, err := val.Pntr(); err != nil {
			rv = uw.Hash(iterV.(*uint)).Combine(rv)
		}
	case reflect.Uint8:
		u8w := widgets.BuiltinUint8{}
		if iterV, err := val.Pntr(); err != nil {
			rv = u8w.Hash(iterV.(*uint8)).Combine(rv)
		}
	case reflect.Uint16:
		u16w := widgets.BuiltinUint16{}
		if iterV, err := val.Pntr(); err != nil {
			rv = u16w.Hash(iterV.(*uint16)).Combine(rv)
		}
	case reflect.Uint32:
		u32w := widgets.BuiltinUint32{}
		if iterV, err := val.Pntr(); err != nil {
			rv = u32w.Hash(iterV.(*uint32)).Combine(rv)
		}
	case reflect.Uint64:
		u64w := widgets.BuiltinUint64{}
		if iterV, err := val.Pntr(); err != nil {
			rv = u64w.Hash(iterV.(*uint64)).Combine(rv)
		}
	case reflect.Float32:
		f32w := widgets.BuiltinFloat32{}
		if iterV, err := val.Pntr(); err != nil {
			rv = f32w.Hash(iterV.(*float32)).Combine(rv)
		}
	case reflect.Float64:
		f64w := widgets.BuiltinFloat64{}
		if iterV, err := val.Pntr(); err != nil {
			rv = f64w.Hash(iterV.(*float64)).Combine(rv)
		}
	case reflect.Complex64:
		c64w := widgets.BuiltinComplex64{}
		if iterV, err := val.Pntr(); err != nil {
			rv = c64w.Hash(iterV.(*complex64)).Combine(rv)
		}
	case reflect.Complex128:
		c128w := widgets.BuiltinComplex128{}
		if iterV, err := val.Pntr(); err != nil {
			rv = c128w.Hash(iterV.(*complex128)).Combine(rv)
		}
	case reflect.String:
		sw := widgets.BuiltinString{}
		if iterV, err := val.Pntr(); err != nil {
			rv = sw.Hash(iterV.(*string)).Combine(rv)
		}
	case reflect.Uintptr:
		uw := widgets.BuiltinUintptr{}
		if iterV, err := val.Pntr(); err != nil {
			rv = uw.Hash(iterV.(*uintptr)).Combine(rv)
		}
	case reflect.Pointer:
		if !opts.GetFlag(followPntrs) {
			uw := widgets.BuiltinUintptr{}
			if iterV, err := val.ReflectPntr(); err != nil {
				v := uintptr(iterV.Pointer())
				rv = uw.Hash(&v)
			}
		} else {
			// get hash of underlying value
		}
	case reflect.Chan:
	case reflect.Array:
	case reflect.Slice:
	case reflect.Map:
	case reflect.Func:
	case reflect.Interface:
	case reflect.UnsafePointer:
	case reflect.Struct:
		fallthrough // iteration will be handled by the calling func
	case reflect.Invalid:
		fallthrough // Ignore "bad" values
	default:
	}
	return 0
}

func StructHash[T any, S reflect.Value | *T](
	s S,
	opts *structHashOpts,
) hash.Hash {
	var rv hash.Hash
	RecursiveStructFieldInfo[T, S](s, false).ForEach(
		func(index int, val FieldInfo) (iter.IteratorFeedback, error) {
			rv.Combine(getHash(val, opts))
			return iter.Continue, nil
		},
	)
	return rv
}
