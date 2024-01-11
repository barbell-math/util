package reflect

import (
	"reflect"

	"github.com/barbell-math/util/algo/iter"
)

type ValInfo struct {
    Val any
    Type reflect.Type
    Kind reflect.Kind
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
        default:
            if err:=sliceValError(any(s).(*T)); err!=nil {
                return reflect.Value{},err
            }
            return reflect.ValueOf(any(s).(*T)).Elem(),nil
    }
}

// TODO -test
func ArrayElemVals[T any, A reflect.Value | *T](a A) iter.Iter[any] {
    arrayVal,err:=getArrayVal[T,A](a)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    return iter.SequentialElems[any](
        arrayVal.Len(),
        func(i int) (any, error) { return arrayVal.Index(i),nil },
    )
}

// func ArrayElemPntrs[T any, A reflect.Value | *T](a A) iter.Iter[any] {
// 
// }
// 
// func ArrayElemType[T any, A reflect.Value | *T](a A) reflect.Type {
// 
// }
// 
// func ArrayElemKind[T any, A reflect.Value | *T](a A) reflect.Kind {
// 
// }
// 
// func ArrayElemInfo[T any, A reflect.Value | *T](a A) ValInfo {
// 
// }
// 
// func RecursiveArrayElemInfo[T any, A reflect.Value | *T](a A) iter.Iter[ValInfo] {
// 
// }
