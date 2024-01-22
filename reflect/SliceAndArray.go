package reflect

import (
	"fmt"
	"reflect"

	"github.com/barbell-math/util/algo/iter"
)


// Returns true if the supplied value is a pointer to an array. As a special
// case, if a reflect.Value is passed to this function it will return true if
// that reflect value either contains an array or contains a pointer to an array.
func IsArrayVal[T any, A reflect.Value | *T](a A) bool {
    return isKindOrReflectValKind[T,A](a,reflect.Array)
}

func arrayValError[T any, A reflect.Value | *T](a A) error {
    return valError[T,A](a, reflect.Array, IsArrayVal[T,A])
}

func homogonizeArrayVal[T any, A reflect.Value | *T](a A) (reflect.Value,error) {
    return homogonizeValue[T,A](a,arrayValError[T,A])
}

// Returns true if the supplied value is a pointer to a slice. As a special
// case, if a reflect.Value is passed to this function it will return true if 
// that reflect value either contains a slice or contains a pointer to a slice.
func IsSliceVal[T any, S reflect.Value | *T](s S) bool {
    return isKindOrReflectValKind[T,S](s,reflect.Slice)
}

func sliceValError[T any, S reflect.Value | *T](s S) error {
    return valError[T,S](s, reflect.Slice, IsSliceVal[T,S])
}

func homogonizeSliceVal[T any, S reflect.Value | *T](s S) (reflect.Value,error) {
    return homogonizeValue[T,S](s,sliceValError[T,S])
}

// Returns an iterator that will iterate over the array elements returning the
// value at each index if an array is supplied as an argument, returns an error
// otherwise. As a special case, if a reflect.Value is passed to this function 
// it will return an iterator over the array it contains or an iterator over 
// the array it points to if the reflect.Value contains a pointer to an array.
func ArrayElemVals[T any, A reflect.Value | *T](a A) iter.Iter[any] {
    return elemVals[T,A](a,homogonizeArrayVal[T,A])
}

// Returns an iterator that will iterate over the slice elements returning the
// value at each index if a slice is supplied as an argument, returns an error
// otherwise. As a special case, if a reflect.Value is passed to this function 
// it will return an iterator over the slice it contains or an iterator over 
// the slice it points to if the reflect.Value contains a pointer to a slice.
func SliceElemVals[T any, S reflect.Value | *T](a S) iter.Iter[any] {
    return elemVals[T,S](a,homogonizeSliceVal[T,S])
}

func elemVals[T any, U reflect.Value |*T](
    a U, 
    homoginizer func(a U) (reflect.Value,error),
) iter.Iter[any] {
    arrayVal,err:=homoginizer(a)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    return iter.SequentialElems[any](
        arrayVal.Len(),
        func(i int) (any, error) { return arrayVal.Index(i).Interface(),nil },
    )
}

// Returns an iterator that will iterate over the array elements returning a
// pointer at each index if an array is supplied as an argument, returns an error
// otherwise. As a special case, if a reflect.Value is passed to this function 
// it will return an iterator over the array it contains or an iterator over 
// the array it points to if the reflect.Value contains a pointer to an array.
// Note that this function requires any reflect.Value handed to it to be 
// addressable. This means that in most scenarios any reflect.Value passed to
// this function will need to contain a pointer to an array.
func ArrayElemPntrs[T any, A reflect.Value | *T](a A) iter.Iter[any] {
    return elemPntrs[T,A](a,homogonizeArrayVal[T,A])
}

// Returns an iterator that will iterate over the slice elements returning a
// pointer at each index if a slice is supplied as an argument, returns an error
// otherwise. As a special case, if a reflect.Value is passed to this function 
// it will return an iterator over the slice it contains or an iterator over 
// the slice it points to if the reflect.Value contains a pointer to a slice.
func SliceElemPntrs[T any, A reflect.Value | *T](a A) iter.Iter[any] {
    return elemPntrs[T,A](a,homogonizeSliceVal[T,A])
}

func elemPntrs[T any, U reflect.Value | *T](
    u U,
    homoginizer func(a U) (reflect.Value,error),
) iter.Iter[any] {
    arrayVal,err:=homoginizer(u)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    return iter.SequentialElems[any](
        arrayVal.Len(),
        func(i int) (any, error) { 
            if arrayVal.Index(i).CanAddr() {
                return arrayVal.Index(i).Addr().Interface(),nil
            }
            return nil,InAddressableField(fmt.Sprintf("Index: %d",i))
        },
    )
}

// Returns the type of the elements in the array that is passed to the function.
// If an array is not passed to the function an error will be returned. As a 
// special case if a reflect.Value is passed to this function then the type of 
// the elements in the array it either contains or points to will be returned.
func ArrayElemType[T any, A reflect.Value | *T](a A) (reflect.Type,error) {
    return elemType[T,A](a,homogonizeArrayVal[T,A])
}

// Returns the type of the elements in the slice that is passed to the function.
// If a slice is not passed to the function an error will be returned. As a 
// special case if a reflect.Value is passed to this function then the type of 
// the elements in the slice it either contains or points to will be returned.
func SliceElemType[T any, A reflect.Value | *T](a A) (reflect.Type,error) {
    return elemType[T,A](a,homogonizeSliceVal[T,A])
}

func elemType[T any, U reflect.Value | *T](
    u U,
    homoginizer func(a U) (reflect.Value,error),
) (reflect.Type,error) {
    arrayVal,err:=homoginizer(u)
    if err!=nil {
        return nil,err
    }
    return arrayVal.Type().Elem(),nil
}

// Returns the kind of the elements in the array that is passed to the function.
// If an array is not passed to the function an error will be returned. As a 
// special case if a reflect.Value is passed to this function then the kind of 
// the elements in the array it either contains or points to will be returned.
func ArrayElemKind[T any, A reflect.Value | *T](a A) (reflect.Kind,error) {
    return elemKind[T,A](a,homogonizeArrayVal[T,A])
}

// Returns the kind of the elements in the slice that is passed to the function.
// If a slice is not passed to the function an error will be returned. As a 
// special case if a reflect.Value is passed to this function then the kind of 
// the elements in the slice it either contains or points to will be returned.
func SliceElemKind[T any, A reflect.Value | *T](a A) (reflect.Kind,error) {
    return elemKind[T,A](a,homogonizeSliceVal[T,A])
}

func elemKind[T any, U reflect.Value | *T](
    u U,
    homoginizer func(a U) (reflect.Value,error),
) (reflect.Kind,error) {
    arrayVal,err:=homoginizer(u)
    if err!=nil {
        return 0,err
    }
    return arrayVal.Type().Elem().Kind(),nil
}

// Returns an iterator that provides the value information for each element in
// the supplied array if an array is supplied as an argument, returns an error
// otherwise. As a special case if a reflect.Value is passed to this function 
// then the iterator will iterate over the elements in the array it either 
// contains or points to.
// Note that the field info Pntr field may not be able to be populated if the
// passed in value is not addressable. If you need the pointers to the array 
// elements then make sure you either pass a pointer to an array or a
// reflect.Value that contains a pointer to an array.
func ArrayElemInfo[T any, A reflect.Value | *T](a A, keepVal bool) iter.Iter[ValInfo] {
    return elemInfo[T,A](a,homogonizeArrayVal[T,A],keepVal)
}

// Returns an iterator that provides the value information for each element in
// the supplied slice if a slice is supplied as an argument, returns an error
// otherwise. As a special case if a reflect.Value is passed to this function 
// then the iterator will iterate over the elements in the slice it either 
// contains or points to.
func SliceElemInfo[T any, A reflect.Value | *T](a A, keepVal bool) iter.Iter[ValInfo] {
    return elemInfo[T,A](a,homogonizeSliceVal[T,A],keepVal)
}

func elemInfo[T any, U reflect.Value | *T](
    u U,
    homoginizer func(a U) (reflect.Value,error),
    keepVal bool,
) iter.Iter[ValInfo] {
    arrayVal,err:=homoginizer(u)
    if err!=nil {
        return iter.ValElem[ValInfo](ValInfo{},err,1)
    }
    return iter.SequentialElems[ValInfo](
        arrayVal.Len(),
        func(i int) (ValInfo, error) {
            return ValInfo{
                Type: arrayVal.Index(i).Type(),
                Kind: arrayVal.Index(i).Kind(),
                Val: func() (any, bool) {
                    if keepVal {
                        return arrayVal.Index(i).Interface(),true
                    }
                    return nil,false
                },
                Pntr: func() (any, error) {
                    if arrayVal.Index(i).CanAddr() {
                        return arrayVal.Index(i).Addr().Interface(),nil
                    }
                    return nil,InAddressableField(fmt.Sprintf("Index: %d",i))
                },
            },nil
        },
    )
}

// Returns an iterator that recursively provides the array elements val info if 
// an array is supplied as an argument, returns an error otherwise. As a 
// special case, if a reflect.Value is passed to this function it will return 
// the recursively found val info of the arrays it contains or the recursively
// found val info of the array that it points to if the reflect.Value 
// contains a pointer to an array. Any field that is an array value will be 
// recursed on, pointers to arrays will not be recursed on.
// Note that in order to recursively access the fields the array needs to be 
// addressable, as the fields that are arrays will be referenced through 
// pointers. This is done to prevent excess memory use that would be caused by
// copying all sub-arrays by value.
func RecursiveArrayElemInfo[T any, A reflect.Value | *T](
    a A,
    keepVal bool,
) iter.Iter[ValInfo] {
    if err:=arrayValError[T,A](a); err!=nil {
        return iter.ValElem[ValInfo](ValInfo{},err,1)
    }
    return iter.Recurse[ValInfo](
        ArrayElemInfo[T,A](a,keepVal),
        func(v ValInfo) bool { return v.Kind==reflect.Array },
        func(v ValInfo) iter.Iter[ValInfo] {
            if v,err:=v.Pntr(); err==nil {
                return ArrayElemInfo[T,reflect.Value](reflect.ValueOf(v),keepVal)
            } else {
                return iter.ValElem[ValInfo](ValInfo{},err,1)
            }
        },
    )
}

// Returns an iterator that recursively provides the slice elements val info if 
// a slice is supplied as an argument, returns an error otherwise. As a 
// special case, if a reflect.Value is passed to this function it will return 
// the recursively found val info of the slices it contains or the recursively
// found val info of the slice that it points to if the reflect.Value 
// contains a pointer to a slice. Any field that is a slice value will be 
// recursed on, pointers to slices will not be recursed on.
func RecursiveSliceElemInfo[T any, S reflect.Value | *T](
    s S, 
    keepVal bool,
) iter.Iter[ValInfo] {
    if err:=sliceValError[T,S](s); err!=nil {
        return iter.ValElem[ValInfo](ValInfo{},err,1)
    }
    return iter.Recurse[ValInfo](
        SliceElemInfo[T,S](s,keepVal),
        func(v ValInfo) bool { return v.Kind==reflect.Slice },
        func(v ValInfo) iter.Iter[ValInfo] {
            if v,err:=v.Pntr(); err==nil {
                return SliceElemInfo[T,reflect.Value](reflect.ValueOf(v),keepVal)
            } else {
                return iter.ValElem[ValInfo](ValInfo{},err,1)
            }
        },
    )
}
