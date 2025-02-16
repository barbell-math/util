package reflect

import (
	"reflect"

	"github.com/barbell-math/util/src/hash"
	"github.com/barbell-math/util/src/iter"
	"github.com/barbell-math/util/src/widgets"
)

//go:generate ../../bin/enum -type=optionsFlag -package=reflect
//go:generate ../../bin/flags -type=optionsFlag -package=reflect
//go:generate ../../bin/structDefaultInit -struct=structHashOpts

type (
	//gen:enum unknownValue unknownOptionsFlag
	//gen:enum default includeMapVals | includeArrayVals | includeSliceVals | followPntrs | followInterfaces | recurseStructs
	optionsFlag    int
	structHashOpts struct {
		//gen:structDefaultInit default NewOptionsFlag()
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		optionsFlag
	}
)

const (
	// Description: set to true if the hash value should be calculated by
	// following pointer values rather than using the pointers value itself
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string followPntrs
	followPntrs optionsFlag = 1 << iota
	// Description: set to true if the hash value should be calculated by
	// following interface value
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string followInterfaces
	followInterfaces
	// Description: set to true if the hash value should be calculated by
	// including all sub-struct fields
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string recurseStructs
	recurseStructs
	// Description: set to true to include map key value pairs in the hash
	// calculation. If false the address of the map will be used when
	// calculating the hash.
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string includeMapVals
	includeMapVals
	// Description: set to true to include slice values in the hash calculation.
	// If false the address of the slice will be used when calculating the hash.
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string includeSliceVals
	includeSliceVals
	// Description: set to true to include array values in the hash
	// calculation. If false the address of the slice will be used when
	// calculating the hash.
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string includeArrayVals
	includeArrayVals
	//gen:flags noSetter
	//gen:enum string unknownOptionsFlag
	unknownOptionsFlag
)

func typeCastHashOp[T any, W widgets.BaseInterface[T]](
	val ValInfo,
	h *hash.Hash,
) {
	w := widgets.Base[T, W]{}
	if !val.v.Type().AssignableTo(reflect.TypeOf((*T)(nil)).Elem()) {
		return
	}
	if iterV, err := val.Pntr(); err == nil {
		if v, ok := iterV.(*T); ok {
			*h = h.CombineIgnoreZero(w.Hash(v))
		}
	} else if iterV, ok := val.Val(); ok {
		if v, ok := iterV.(T); ok {
			*h = h.CombineIgnoreZero(w.Hash(&v))
		}
	}
}

// Updates the hash value using the address of the supplied value.
func addrHashOp(val ValInfo, h *hash.Hash) {
	w := widgets.BuiltinUintptr{}
	if CanBeNil(val.v) && val.v.IsNil() {
		return
	}
	if val.v.CanAddr() {
		v := uintptr(val.v.Addr().Pointer())
		*h = h.CombineIgnoreZero(w.Hash(&v))
	}
}

// Updates the hash value using the supplied pointers value. Noop if a pointer
// is not supplied.
func pointerHashOp(val ValInfo, h *hash.Hash) {
	w := widgets.BuiltinUintptr{}
	if CanBeNil(val.v) && val.v.IsNil() {
		return
	}
	if val.v.Kind() != reflect.Pointer {
		return
	}
	v := uintptr(val.v.Pointer())
	*h = h.CombineIgnoreZero(w.Hash(&v))
}

func getHash(val ValInfo, opts *structHashOpts) hash.Hash {
	var rv hash.Hash
	switch val.Kind {
	case reflect.Bool:
		typeCastHashOp[bool, widgets.BuiltinBool](val, &rv)
	case reflect.Int:
		typeCastHashOp[int, widgets.BuiltinInt](val, &rv)
	case reflect.Int8:
		typeCastHashOp[int8, widgets.BuiltinInt8](val, &rv)
	case reflect.Int16:
		typeCastHashOp[int16, widgets.BuiltinInt16](val, &rv)
	case reflect.Int32:
		typeCastHashOp[int32, widgets.BuiltinInt32](val, &rv)
	case reflect.Int64:
		typeCastHashOp[int64, widgets.BuiltinInt64](val, &rv)
	case reflect.Uint:
		typeCastHashOp[uint, widgets.BuiltinUint](val, &rv)
	case reflect.Uint8:
		typeCastHashOp[uint8, widgets.BuiltinUint8](val, &rv)
	case reflect.Uint16:
		typeCastHashOp[uint16, widgets.BuiltinUint16](val, &rv)
	case reflect.Uint32:
		typeCastHashOp[uint32, widgets.BuiltinUint32](val, &rv)
	case reflect.Uint64:
		typeCastHashOp[uint64, widgets.BuiltinUint64](val, &rv)
	case reflect.Float32:
		typeCastHashOp[float32, widgets.BuiltinFloat32](val, &rv)
	case reflect.Float64:
		typeCastHashOp[float64, widgets.BuiltinFloat64](val, &rv)
	case reflect.Complex64:
		typeCastHashOp[complex64, widgets.BuiltinComplex64](val, &rv)
	case reflect.Complex128:
		typeCastHashOp[complex128, widgets.BuiltinComplex128](val, &rv)
	case reflect.String:
		typeCastHashOp[string, widgets.BuiltinString](val, &rv)
	case reflect.Uintptr:
		typeCastHashOp[uintptr, widgets.BuiltinUintptr](val, &rv)
	case reflect.Pointer:
		if !val.v.IsNil() {
			if opts.GetFlag(followPntrs) {
				rv = rv.CombineIgnoreZero(getHash(
					NewValInfo(reflect.Indirect(val.v), false, nil),
					opts,
				))
			} else {
				pointerHashOp(val, &rv)
			}
		}
	case reflect.Array:
		rv = rv.CombineIgnoreZero(hash.Hash(val.v.Len()))
		if opts.GetFlag(includeArrayVals) {
			for i := 0; i < val.v.Len(); i++ {
				rv = rv.CombineIgnoreZero(getHash(
					NewValInfo(val.v.Index(i), false, nil),
					opts,
				))
			}
		}
	case reflect.Slice:
		rv = rv.CombineIgnoreZero(hash.Hash(val.v.Len()))
		if opts.GetFlag(includeSliceVals) {
			for i := 0; i < val.v.Len(); i++ {
				rv = rv.CombineIgnoreZero(getHash(
					NewValInfo(val.v.Index(i), false, nil),
					opts,
				))
			}
		} else if val.v.Len() > 0 {
			addrHashOp(
				NewValInfo(val.v.Index(0), false, nil),
				&rv,
			)
		}
	case reflect.Map:
		rv = rv.CombineIgnoreZero(hash.Hash(val.v.Len()))
		if opts.GetFlag(includeMapVals) {
			mapIter := val.v.MapRange()
			for mapIter.Next() {
				rv = rv.CombineUnorderedIgnoreZero(getHash(
					NewValInfo(mapIter.Key(), true, InaddressableMapErr),
					opts,
				))
				rv = rv.CombineUnorderedIgnoreZero(getHash(
					NewValInfo(mapIter.Value(), true, InaddressableMapErr),
					opts,
				))
			}
		}
	case reflect.Interface:
		if opts.GetFlag(followInterfaces) && !val.v.IsNil() {
			rv = rv.CombineIgnoreZero(getHash(
				NewValInfo(
					reflect.ValueOf(val.v.Interface()),
					true,
					InaddressableValueErr,
				),
				opts,
			))
		}
	case reflect.Struct:
		if opts.GetFlag(recurseStructs) {
			for i := 0; i < val.v.NumField(); i++ {
				t := val.v.Type().Field(i)
				f := val.v.Field(i)
				if t.IsExported() {
					rv = rv.CombineIgnoreZero(getHash(
						NewValInfo(f, false, nil),
						opts,
					))
				}
			}
		}
	case reflect.Chan:
		fallthrough // no good way to get a unique value to represent a chan
	case reflect.Func:
		fallthrough // no good way to get a unique value to represent a func
	case reflect.UnsafePointer:
		fallthrough // do nothing because it's unsafe
	case reflect.Invalid:
		fallthrough // Ignore "bad" values
	default:
	}
	return rv
}

// Returns a hash for the supplied struct. The hash is generated by reflexively
// iterating over exported structs fields. There are a couple rules that dictate
// the hash value that is generated:
//
//  1. The widgets in the [widgets] package are used for generating hashes when
//
// the underlying type is an exact match. Custom types of the built in types
// will not have a hash value generated.
//  2. Only exported fields are used. A struct with no exported fields will have
//
// a hash of 0.
//  3. nil values do not contribute to the hash value. A value of 0 does.
//  4. Arrays, slices, and maps lengths will always be included in the
//
// calculated hash.
//  5. The memory address of a slices data pointer will be included in the
//
// calculated hash only if the underlying values are flagged to not be included.
//  6. Channels and functions will be ignored.
//
// The opts struct determines behaviors about which struct fields are used and
// how they are used.
func StructHash[T any, S reflect.Value | *T](
	s S,
	opts *structHashOpts,
) (hash.Hash, error) {
	var rv hash.Hash
	err := StructFieldInfo[T, S](s, false).ForEach(
		func(index int, val FieldInfo) (iter.IteratorFeedback, error) {
			rv = rv.CombineIgnoreZero(getHash(val.ValInfo, opts))
			return iter.Continue, nil
		},
	)
	return rv, err
}
