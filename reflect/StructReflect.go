package reflect

import (
	"fmt"
	"reflect"

	"github.com/barbell-math/util/algo/iter"
)

// This struct has fields that contain all the possible relevant information 
// about a structs field.
type FieldInfo struct {
    // The name of the field.
    Name string
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

// Returns true if the supplied value is a pointer to a struct. As a special
// case, if a reflect.Value is passed to this function it will return true if 
// that reflect value either contains a struct or contains a pointer to a struct.
func IsStructVal[S any](s *S) bool {
    switch reflect.TypeOf(s) {
        case reflect.TypeOf(&reflect.Value{}):
            if refVal:=any(s).(*reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem().Kind()==reflect.Struct
            } else {
                return refVal.Kind()==reflect.Struct
            }
        default: return reflect.ValueOf(s).Elem().Kind()==reflect.Struct
    }
}

func structValError[S any](s *S) error {
    if IsStructVal(s) {
        return nil;
    }
    var fString string
    switch reflect.TypeOf(s) {
        case reflect.TypeOf(&reflect.Value{}):
            if refVal:=any(s).(*reflect.Value); refVal.Kind()==reflect.Ptr {
                fString=fmt.Sprintf(
                    "Got a reflect.Value pointer to: %s",
                    refVal.Elem().Kind().String(),
                )
            } else {
                fString=fmt.Sprintf(
                    "Got a reflect.Value containing: %s",
                    refVal.Kind().String(),
                )
            }
        default:
            fString=fmt.Sprintf(
                "Got a pointer to: %s",
                reflect.ValueOf(s).Elem().Kind().String(),
            )
    }
    return NonStructValue(fmt.Sprintf(
        "Function requires a struct as target. | %s",
        fString,
    ));
}

// // Returns true if the supplied value is a pointer to a pointer to a struct.
// func IsStructPntr[S any](s **S) bool {
//     switch reflect.TypeOf(s) {
//         case reflect.TypeOf(&reflect.Value{}):
//             return any(s).(*reflect.Value).Elem().Kind()==reflect.Struct
//         default: return reflect.ValueOf(s).Elem().Elem().Kind()==reflect.Struct
//     }
// }
// 
// // func IsStructPntr[S any](s **S) bool {
// //     return reflect.ValueOf(s).Elem().Elem().Kind()==reflect.Struct
// // }
// 
// func structPntrError[S any](s **S) error {
//     if !IsStructPntr(s) {
//         return NonStructValue(fmt.Sprintf(
//             "Function requires a struct as target. | Got: %s",
//             reflect.ValueOf(s).Kind().String(),
//         ));
//     }
//     return nil;
// }

func getStructVal[T any, S reflect.Value | *T](s S) (reflect.Value,error) {
    switch reflect.TypeOf(s) {
        case reflect.TypeOf(reflect.Value{}):
            if err:=structValError(&s); err!=nil {
                return reflect.Value{},err
            }
            if refVal:=any(s).(reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem(),nil
            } else {
                return refVal,nil
            }
        default:
            if err:=structValError(any(s).(*T)); err!=nil {
                return reflect.Value{},err
            }
            return reflect.ValueOf(any(s).(*T)).Elem(),nil
    }
}

// Retrieves the struct name if a struct is supplied as an argument, returns 
// an error otherwise. As a special case, if a reflect.Value is passed to this 
// function it will return get the name of the struct it contains or the name of
// the struct that it points to if the reflect.Value contains a pointer to a 
// struct.
func GetStructName[T any, S reflect.Value | *T](s S) (string,error) {
    if structVal,err:=getStructVal[T,S](s); err==nil {
        return structVal.Type().Name(),nil
    } else {
        return "",err
    }
}

// Returns an iterator that provides the struct field names if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// names of the struct it contains or the field names of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
func StructFieldNames[T any, S reflect.Value | *T](s S) iter.Iter[string] {
    structVal,err:=getStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[string]("",err,1)
    }
    i:=-1
    return func(f iter.IteratorFeedback) (string, error, bool) {
        i++
        if f!=iter.Break && i<structVal.NumField() {
            return structVal.Type().Field(i).Name, nil, true
        }
        return "", nil, false
    }
}

// Returns an iterator that provides the struct field values if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// values of the struct it contains or the field values of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
func StructFieldVals[T any, S reflect.Value | *T](s S) iter.Iter[any] {
    structVal,err:=getStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    i:=-1
    return func(f iter.IteratorFeedback) (any, error, bool) {
        i++
        if f!=iter.Break && i<structVal.NumField() {
            return reflect.ValueOf(structVal.Field(i).Interface()).Interface(), nil, true
        }
        return reflect.Value{}, nil, false
    }
}

// Returns an iterator that provides pointers to the struct field values if a
// struct is is supplied as an argument, returns an error otherwise. As a 
// special case, if a reflect.Value is passed to this function it will return 
// pointers to the fields of the struct it contains or pointers to the fields 
// of the struct that it points to if the reflect.Value contains a pointer to a 
// struct.
// Note that this function requires any reflect.Value handed to it to be 
// addressable. This means that in most scenarios any reflect.Value passed to
// this function will need to contain a pointer to a struct.
func StructFieldPntrs[T any, S reflect.Value | *T](s S) iter.Iter[any] {
    structVal,err:=getStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    i:=-1
    return func(f iter.IteratorFeedback) (any, error, bool) {
        i++
        if f!=iter.Break && i<structVal.NumField() {
            f:=structVal.Field(i)
            if f.CanAddr() {
                return f.Addr().Interface(), nil, true
            }
            return reflect.Value{},InAddressableField(fmt.Sprintf(
                "Field Name: %s",structVal.Type().Field(i).Name,
            )),false
        }
        return reflect.Value{}, nil, false
    }
}

// Returns an iterator that provides the struct field types if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// types of the struct it contains or the field types of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
func StructFieldTypes[T any, S reflect.Value | *T](s S) iter.Iter[reflect.Type] {
    structVal,err:=getStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[reflect.Type](nil,err,1)
    }
    i:=-1
    return func(f iter.IteratorFeedback) (reflect.Type, error, bool) {
        i++
        if f!=iter.Break && i<structVal.NumField() {
            return structVal.Field(i).Type(),nil,true
        }
        return nil, nil, false
    }
}

// Returns an iterator that provides the struct field kinds if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// kinds of the struct it contains or the field kinds of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
func StructFieldKinds[T any, S reflect.Value | *T](s S) iter.Iter[reflect.Kind] {
    structVal,err:=getStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[reflect.Kind](0,err,1)
    }
    i:=-1
    return func(f iter.IteratorFeedback) (reflect.Kind, error, bool) {
        i++
        if f!=iter.Break && i<structVal.NumField() {
            return structVal.Field(i).Kind(),nil,true
        }
        return 0, nil, false
    }
}

// Returns an iterator that provides the struct field info if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// info of the struct it contains or the field info of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
// Note that the field info Pntr field may not be able to be populated if the
// passed in value is not addressable. If you need the pointers to the struct 
// fields then make sure you either pass a pointer to a struct or a 
// reflect.Value that contains a pointer to a struct.
func StructFieldInfo[T any, S reflect.Value | *T](s S) iter.Iter[FieldInfo] {
    structVal,err:=getStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[FieldInfo](FieldInfo{},err,1)
    }
    i:=-1
    return func(f iter.IteratorFeedback) (FieldInfo, error, bool) {
        i++
        if f!=iter.Break && i<structVal.NumField() {
            n:=structVal.Type().Field(i).Name
            f:=structVal.Field(i)
            return FieldInfo{
                Name: n,
                Val: reflect.ValueOf(f.Interface()).Interface(),
                Type: f.Type(),
                Kind: f.Kind(),
                Pntr: func() (any, error) {
                    if f.CanAddr() {
                        return f.Addr().Interface(),nil
                    }
                    return nil,InAddressableField(fmt.Sprintf(
                        "Field Name: %s",n,
                    ))
                },
            }, nil, true
        }
        return FieldInfo{}, nil, false
    }
}

// Returns an iterator that recursively provides the struct field info if a 
// struct is is supplied as an argument, returns an error otherwise. As a 
// special case, if a reflect.Value is passed to this function it will return 
// the recursively found field info of the struct it contains or the recursively
// found field info of the struct that it points to if the reflect.Value 
// contains a pointer to a struct. Any field that is a struct value will be 
// recursed on, pointers to structs will not be recursed on.
// Note that in order to recursively access the fields the struct needs to be 
// addressable, as the fields that are structs will be referenced through 
// pointers. This is done to prevent excess memory use that would be caused by
// copying all sub-structs by value.
func RecursiveStructFieldInfo[T any, S reflect.Value | *T](
    s S,
) iter.Iter[FieldInfo] {
    if _,err:=getStructVal[T,S](s); err!=nil {
        return iter.ValElem[FieldInfo](FieldInfo{},err,1)
    }
    return iter.Recurse[FieldInfo](
        StructFieldInfo[T,S](s),
        func(v FieldInfo) bool { return v.Kind==reflect.Struct },
        func(v FieldInfo) iter.Iter[FieldInfo] {
            if p,err:=v.Pntr(); err==nil {
                return StructFieldInfo[reflect.Value](reflect.ValueOf(p))
            } else {
                return iter.ValElem[FieldInfo](FieldInfo{},err,1)
            }
        },
    )
}
