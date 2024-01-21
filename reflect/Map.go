package reflect

import (
	"reflect"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/dataStruct"
)

// A simple type to hold key value pairs from a map.
type KeyValue dataStruct.Pair[any,any]

// This type has fields that contain all the possible relevant information
// about a maps key value pair.
type KeyValInfo dataStruct.Pair[ValInfo,ValInfo]

// Returns true if the supplied value is a pointer to a map. As a special
// case, if a reflect.Value is passed to this function it will return true if
// that reflect value either contains a map or contains a pointer to a map.
func IsMapVal[T any, M reflect.Value | *T](m M) bool {
    return isKindOrReflectValKind[T,M](m,reflect.Map)
}

func mapValError[T any, M reflect.Value | *T](m M) error {
    return valError[T,M](m,reflect.Map,IsMapVal[T,M])
}

func homogonizeMapVal[T any, M reflect.Value | *T](m M) (reflect.Value,error) {
    return homogonizeValue[T,M](m,mapValError[T,M])
}

// Returns an iterator that will iterate over the supplied maps keys if a map is 
// supplied as an argument, returns an error otherwise. As a special case, if a 
// reflect.Value is passed to this function it will return an iterator over the
// the map keys it contains or an iterator over the map keys it points to if the 
// reflect.Value contains a pointer to a map.
func MapElemKeys[T any, M reflect.Value | *T](m M) iter.Iter[any] {
    mapVal,err:=homogonizeMapVal[T,M](m)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    return iter.Map[reflect.Value,any](
        iter.SliceElems[reflect.Value](mapVal.MapKeys()),
        func(index int, val reflect.Value) (any, error) {
            return val.Interface(),nil
        },
    )
}

// Returns an iterator that will iterate over the supplied maps values if a map is 
// supplied as an argument, returns an error otherwise. As a special case, if a 
// reflect.Value is passed to this function it will return an iterator over the
// the map values it contains or an iterator over the map values it points to if the 
// reflect.Value contains a pointer to a map.
func MapElemVals[T any, M reflect.Value | *T](m M) iter.Iter[any] {
    mapVal,err:=homogonizeMapVal[T,M](m)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    return iter.Map[reflect.Value,any](
        iter.SliceElems[reflect.Value](mapVal.MapKeys()),
        func(index int, val reflect.Value) (any, error) {
            return mapVal.MapIndex(val).Interface(),nil
        },
    )
}

// Returns an iterator that will iterate over the supplied maps key value pairs 
// if a map is supplied as an argument, returns an error otherwise. As a special 
// case, if a reflect.Value is passed to this function it will return an iterator 
// over the the map key value pairs it contains or an iterator over the map 
// values it points to if the reflect.Value contains a pointer to a map.
func MapElems[T any, M reflect.Value | *T](m M) iter.Iter[KeyValue] {
    mapVal,err:=homogonizeMapVal[T,M](m)
    if err!=nil {
        return iter.ValElem[KeyValue](KeyValue{},err,1)
    }
    return iter.Map[reflect.Value,KeyValue](
        iter.SliceElems[reflect.Value](mapVal.MapKeys()),
        func(index int, val reflect.Value) (KeyValue, error) {
            return KeyValue{
                A: val.Interface(), 
                B: mapVal.MapIndex(val).Interface(),
            },nil
        },
    )
}

// TODO 
// MapElemKeyInfo
// MapElemValInfo
// MapElemInfo
// RecusriveMapElemInfo

// TODO - test
func MapKeyType[T any, M reflect.Value | *T](m M) reflect.Type {
    mapVal,err:=homogonizeMapVal[T,M](m)
    if err!=nil {
        return nil
    }
    return mapVal.Type().Key()
}

// TODO - test
func MapValType[T any, M reflect.Value | *T](m M) reflect.Type {
    mapVal,err:=homogonizeMapVal[T,M](m)
    if err!=nil {
        return nil
    }
    return mapVal.Type().Elem()
}
