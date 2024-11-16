package reflect

import (
	"reflect"

	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/iter"
)

var (
	InaddressableMapErr = customerr.Wrap(
		InAddressableField, "Maps are not addressable.",
	)
)

// A simple type to hold key value pairs from a map.
type KeyValue basic.Pair[any, any]

// This type has fields that contain all the possible relevant information
// about a maps key value pair.
type KeyValInfo basic.Pair[ValInfo, ValInfo]

// Returns true if the supplied value is a pointer to a map. As a special
// case, if a reflect.Value is passed to this function it will return true if
// that reflect value either contains a map or contains a pointer to a map.
func IsMapVal[T any, M reflect.Value | *T](m M) bool {
	return isKindOrReflectValKind[T, M](m, reflect.Map)
}

func mapValError[T any, M reflect.Value | *T](m M) error {
	return valError[T, M](m, reflect.Map, IsMapVal[T, M])
}

func homogonizeMapVal[T any, M reflect.Value | *T](m M) (reflect.Value, error) {
	return homogonizeValue[T, M](m, mapValError[T, M])
}

// Returns the kind of the keys in the map that is passed to the function.
// If a map is not passed to the function an error will be returned. As a
// special case if a reflect.Value is passed to this function then the kind of
// the keys in the map it either contains or points to will be returned.
func MapKeyKind[T any, M reflect.Value | *T](m M) (reflect.Kind, error) {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return 0, err
	}
	return mapVal.Type().Key().Kind(), nil
}

// Returns the type of the keys in the map that is passed to the function.
// If a map is not passed to the function an error will be returned. As a
// special case if a reflect.Value is passed to this function then the type of
// the keys in the map it either contains or points to will be returned.
func MapKeyType[T any, M reflect.Value | *T](m M) (reflect.Type, error) {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return nil, err
	}
	return mapVal.Type().Key(), nil
}

// Returns the kind of the values in the map that is passed to the function.
// If a map is not passed to the function an error will be returned. As a
// special case if a reflect.Value is passed to this function then the kind of
// the values in the map it either contains or points to will be returned.
func MapValKind[T any, M reflect.Value | *T](m M) (reflect.Kind, error) {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return 0, err
	}
	return mapVal.Type().Elem().Kind(), nil
}

// Returns the type of the values in the map that is passed to the function.
// If a map is not passed to the function an error will be returned. As a
// special case if a reflect.Value is passed to this function then the type of
// the values in the map it either contains or points to will be returned.
func MapValType[T any, M reflect.Value | *T](m M) (reflect.Type, error) {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return nil, err
	}
	return mapVal.Type().Elem(), nil
}

// Returns an iterator that will iterate over the supplied maps keys if a map is
// supplied as an argument, returns an error otherwise. As a special case, if a
// reflect.Value is passed to this function it will return an iterator over the
// the map keys it contains or an iterator over the map keys it points to if the
// reflect.Value contains a pointer to a map.
func MapElemKeys[T any, M reflect.Value | *T](m M) iter.Iter[any] {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return iter.ValElem[any](nil, err, 1)
	}
	return iter.Map[reflect.Value, any](
		iter.SliceElems[reflect.Value](mapVal.MapKeys()),
		func(index int, val reflect.Value) (any, error) {
			return val.Interface(), nil
		},
	)
}

// Returns an iterator that will iterate over the supplied maps values if a map is
// supplied as an argument, returns an error otherwise. As a special case, if a
// reflect.Value is passed to this function it will return an iterator over the
// the map values it contains or an iterator over the map values it points to if the
// reflect.Value contains a pointer to a map.
func MapElemVals[T any, M reflect.Value | *T](m M) iter.Iter[any] {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return iter.ValElem[any](nil, err, 1)
	}
	return iter.Map[reflect.Value, any](
		iter.SliceElems[reflect.Value](mapVal.MapKeys()),
		func(index int, val reflect.Value) (any, error) {
			return mapVal.MapIndex(val).Interface(), nil
		},
	)
}

// Returns an iterator that will iterate over the supplied maps key value pairs
// if a map is supplied as an argument, returns an error otherwise. As a special
// case, if a reflect.Value is passed to this function it will return an iterator
// over the the map key value pairs it contains or an iterator over the map
// values it points to if the reflect.Value contains a pointer to a map.
func MapElems[T any, M reflect.Value | *T](m M) iter.Iter[KeyValue] {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return iter.ValElem[KeyValue](KeyValue{}, err, 1)
	}
	return iter.Map[reflect.Value, KeyValue](
		iter.SliceElems[reflect.Value](mapVal.MapKeys()),
		func(index int, val reflect.Value) (KeyValue, error) {
			return KeyValue{
				A: val.Interface(),
				B: mapVal.MapIndex(val).Interface(),
			}, nil
		},
	)
}

// Returns an iterator that provides the key information for each element
// in the supplied map if a map is supplied as an argument, returns an error
// otherwise. As a special case if a reflect.Value is passed to this function
// then the iterator will iterate over the key elements in the map it either
// contains or points to.
func MapElemKeyInfo[T any, M reflect.Value | *T](
	m M,
	keepVal bool,
) iter.Iter[ValInfo] {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return iter.ValElem[ValInfo](ValInfo{}, err, 1)
	}
	return iter.Map[reflect.Value, ValInfo](
		iter.SliceElems[reflect.Value](mapVal.MapKeys()),
		func(index int, val reflect.Value) (ValInfo, error) {
			return NewValInfo(mapVal, keepVal, InaddressableMapErr), nil
		},
	)
}

// Returns an iterator that provides the value information for each element
// in the supplied map if a map is supplied as an argument, returns an error
// otherwise. As a special case if a reflect.Value is passed to this function
// then the iterator will iterate over the value elements in the map it either
// contains or points to.
func MapElemValInfo[T any, M reflect.Value | *T](
	m M,
	keepVal bool,
) iter.Iter[ValInfo] {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return iter.ValElem[ValInfo](ValInfo{}, err, 1)
	}
	return iter.Map[reflect.Value, ValInfo](
		iter.SliceElems[reflect.Value](mapVal.MapKeys()),
		func(index int, val reflect.Value) (ValInfo, error) {
			return ValInfo{
				Type: mapVal.Type().Elem(),
				Kind: mapVal.Type().Elem().Kind(),
				Val: func() (any, bool) {
					if keepVal {
						return mapVal.MapIndex(val).Interface(), true
					}
					return nil, false
				},
				Pntr: func() (any, error) {
					return nil, InaddressableMapErr
				},
			}, nil
		},
	)
}

// Returns an iterator that provides the key value information for each element
// in the supplied map if a map is supplied as an argument, returns an error
// otherwise. As a special case if a reflect.Value is passed to this function
// then the iterator will iterate over the elements in the map it either
// contains or points to.
func MapElemInfo[T any, M reflect.Value | *T](
	m M,
	keepVal bool,
) iter.Iter[KeyValInfo] {
	mapVal, err := homogonizeMapVal[T, M](m)
	if err != nil {
		return iter.ValElem[KeyValInfo](KeyValInfo{}, err, 1)
	}
	return iter.Map[reflect.Value, KeyValInfo](
		iter.SliceElems[reflect.Value](mapVal.MapKeys()),
		func(index int, val reflect.Value) (KeyValInfo, error) {
			return KeyValInfo{
				A: ValInfo{
					v:    val,
					Type: mapVal.Type().Key(),
					Kind: mapVal.Type().Key().Kind(),
					Val: func() (any, bool) {
						if keepVal {
							return val.Interface(), true
						}
						return nil, false
					},
					Pntr: func() (any, error) {
						return nil, InaddressableMapErr
					},
				},
				B: ValInfo{
					v:    mapVal.MapIndex(val),
					Type: mapVal.Type().Elem(),
					Kind: mapVal.Type().Elem().Kind(),
					Val: func() (any, bool) {
						if keepVal {
							return mapVal.MapIndex(val).Interface(), true
						}
						return nil, false
					},
					Pntr: func() (any, error) {
						return nil, InaddressableMapErr
					},
				},
			}, nil
		},
	)
}

// Returns an iterator that recursively provides the key value info if a map is
// supplied as an argument, returns an error otherwise. As a special case, if a
// reflect.Value is passed to this function it will return the recursively found
// val info of the maps it contains or the recursively found val info of the
// map that it points to if the reflect.Value contains a pointer to a map.
// Any value that is a map value will be recursed on, pointers to maps will not
// be recursed on.
func RecursiveMapElemInfo[T any, M reflect.Value | *T](
	m M,
	keepVal bool,
) iter.Iter[KeyValInfo] {
	if err := mapValError[T, M](m); err != nil {
		return iter.ValElem[KeyValInfo](KeyValInfo{}, err, 1)
	}
	return iter.Recurse[KeyValInfo](
		MapElemInfo[T, M](m, keepVal),
		func(v KeyValInfo) bool { return v.B.Kind == reflect.Map },
		func(v KeyValInfo) iter.Iter[KeyValInfo] {
			if v, ok := v.B.Val(); ok {
				return MapElemInfo[T, reflect.Value](reflect.ValueOf(&v), keepVal)
			} else {
				return iter.ValElem[KeyValInfo](
					KeyValInfo{},
					InaddressableMapErr,
					1,
				)
			}
		},
	)
}
