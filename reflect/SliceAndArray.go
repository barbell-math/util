package reflect

import (
	"fmt"
	"reflect"

	"github.com/barbell-math/util/algo/iter"
)

// This struct has fields that contain all the possible relevant information
// about a value.
type ValInfo struct {
    // The concreete value of the field. This is a copy of the value, not the
    // original value contained in the struct. To access the original value
    // use the Pntr field of this struct.
    Val any
    // The type of the field.
    Type reflect.Type
    // The kind of the field.
    Kind reflect.Kind
    // Returns a pointer to the struct field, if possible. Note that the Pntr 
    // field of this struct is a function that may return an error. This is 
    // because, depending on the value that is passed to the iterator function, 
    // not all struct fields will be addressable.
    Pntr func() (any,error)
}

// Returns true if the supplied value is a pointer to an array. As a special
// case, if a reflect.Value is passed to this function it will return true if
// that reflect value either contains an array or contains a pointer to an array.
func IsArrayVal[A any](a *A) bool {
    return isKindOrReflectValKind[A](a,reflect.Array)
}

func arrayValError[A any](a *A) error {
    return valError[A](a, reflect.Array, IsArrayVal[A])
}

// Returns true if the supplied value is a pointer to a slice. As a special
// case, if a reflect.Value is passed to this function it will return true if 
// that reflect value either contains a slice or contains a pointer to a slice.
func IsSliceVal[S any](s *S) bool {
    return isKindOrReflectValKind[S](s,reflect.Slice)
}

func sliceValError[S any](s *S) error {
    return valError[S](s, reflect.Array, IsSliceVal[S])
}

func getArrayVal[T any, A reflect.Value | *T](a A) (reflect.Value,error) {
    switch reflect.TypeOf(a) {
        case reflect.TypeOf(reflect.Value{}):
            if err:=arrayValError(&a); err!=nil {
                return reflect.Value{},err
            }
            if refVal:=any(a).(reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem(),nil
            } else {
                return refVal,nil
            }
        case reflect.TypeOf(&reflect.Value{}):
            if err:=arrayValError(any(a).(*reflect.Value)); err!=nil {
                return reflect.Value{},err
            }
            if refVal:=any(a).(*reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem(),nil
            } else {
                return *refVal,nil
            }
        default:
            if err:=arrayValError(any(a).(*T)); err!=nil {
                return reflect.Value{},err
            }
            return reflect.ValueOf(any(a).(*T)).Elem(),nil
    }
}

func getSliceVal[T any, S reflect.Value | *T](s S) (reflect.Value,error) {
    switch reflect.TypeOf(s) {
        case reflect.TypeOf(reflect.Value{}):
            if err:=sliceValError(&s); err!=nil {
                return reflect.Value{},err
            }
            if refVal:=any(s).(reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem(),nil
            } else {
                return refVal,nil
            }
        case reflect.TypeOf(&reflect.Value{}):
            if err:=sliceValError(any(s).(*reflect.Value)); err!=nil {
                return reflect.Value{},err
            }
            if refVal:=any(s).(*reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem(),nil
            } else {
                return *refVal,nil
            }
        default:
            if err:=sliceValError(any(s).(*T)); err!=nil {
                return reflect.Value{},err
            }
            return reflect.ValueOf(any(s).(*T)).Elem(),nil
    }
}

// Returns an iterator that will iterate over the array elements returning the
// value at each index if an array is supplied as an argument, returns an error
// otherwise. As a special case, if a reflect.Value is passed to this function 
// it will return an iterator over the array it contains or an iterator over 
// the array it points to if the reflect.Value contains a pointer to an array.
func ArrayElemVals[T any, A reflect.Value | *T](a A) iter.Iter[any] {
    arrayVal,err:=getArrayVal[T,A](a)
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
    arrayVal,err:=getArrayVal[T,A](a)
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
    arrayVal,err:=getArrayVal[T,A](a)
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
    arrayVal,err:=getArrayVal[T,A](a)
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
func ArrayElemInfo[T any, A reflect.Value | *T](a A) iter.Iter[ValInfo] {
    arrayVal,err:=getArrayVal[T,A](a)
    if err!=nil {
        return iter.ValElem[ValInfo](ValInfo{},err,1)
    }
    return iter.SequentialElems[ValInfo](
        arrayVal.Len(),
        func(i int) (ValInfo, error) {
            return ValInfo{
                Val: arrayVal.Index(i).Interface(),
                Type: arrayVal.Index(i).Type(),
                Kind: arrayVal.Index(i).Kind(),
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

// func RecursiveArrayElemInfo[T any, A reflect.Value | *T](a A) iter.Iter[ValInfo] {
// 
// }
