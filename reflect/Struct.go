package reflect

import (
	"reflect"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/customerr"
)

func getInaddressableFieldError(name string) error {
    return customerr.Wrap(InAddressableField,"Field Name: %s",name)
}

// This struct has fields that contain all the possible relevant information 
// about a structs field.
type FieldInfo struct {
    // The name of the field.
    Name string
    // The tag associated with the struct field.
    Tag reflect.StructTag
    // The value information for the field.
    ValInfo
}

// Returns true if the supplied value is a pointer to a struct. As a special
// case, if a reflect.Value is passed to this function it will return true if 
// that reflect value either contains a struct or contains a pointer to a struct.
func IsStructVal[T any, S reflect.Value | *T](s S) bool {
    return isKindOrReflectValKind[T,S](s,reflect.Struct)
}

func structValError[T any, S reflect.Value | *T](s S) error {
    return valError[T,S](s, reflect.Struct, IsStructVal[T,S])
}

func homogonizeStructVal[T any, S reflect.Value | *T](s S) (reflect.Value,error) {
    return homogonizeValue[T,S](s,structValError[T,S])
}

// Retrieves the struct name if a struct is supplied as an argument, returns 
// an error otherwise. As a special case, if a reflect.Value is passed to this 
// function it will return the name of the struct it contains or the name of
// the struct that it points to if the reflect.Value contains a pointer to a 
// struct.
func GetStructName[T any, S reflect.Value | *T](s S) (string,error) {
    if structVal,err:=homogonizeStructVal[T,S](s); err==nil {
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
    structVal,err:=homogonizeStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[string]("",err,1)
    }
    return iter.SequentialElems[string](
        structVal.NumField(),
        func(i int) (string,error) {
            return structVal.Type().Field(i).Name,nil
        },
    )
}

// Returns an iterator that provides the struct field values if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// values of the struct it contains or the field values of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
func StructFieldVals[T any, S reflect.Value | *T](s S) iter.Iter[any] {
    structVal,err:=homogonizeStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    return iter.SequentialElems[any](
        structVal.NumField(),
        func(i int) (any,error) {
            return structVal.Field(i).Interface(),nil
        },
    )
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
    structVal,err:=homogonizeStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[any](nil,err,1)
    }
    return iter.SequentialElems[any](
        structVal.NumField(),
        func(i int) (any,error) {
            f:=structVal.Field(i)
            if f.CanAddr() {
                return f.Addr().Interface(), nil
            }
            return reflect.Value{},getInaddressableFieldError(
                structVal.Type().Field(i).Name,
            )
        },
    )
}

// Returns an iterator that provides the struct field types if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// types of the struct it contains or the field types of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
func StructFieldTypes[T any, S reflect.Value | *T](s S) iter.Iter[reflect.Type] {
    structVal,err:=homogonizeStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[reflect.Type](nil,err,1)
    }
    return iter.SequentialElems[reflect.Type](
        structVal.NumField(),
        func(i int) (reflect.Type,error) {
            return structVal.Field(i).Type(),nil
        },
    )
}

// Returns an iterator that provides the struct field kinds if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// kinds of the struct it contains or the field kinds of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
func StructFieldKinds[T any, S reflect.Value | *T](s S) iter.Iter[reflect.Kind] {
    structVal,err:=homogonizeStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[reflect.Kind](0,err,1)
    }
    return iter.SequentialElems[reflect.Kind](
        structVal.NumField(),
        func(i int) (reflect.Kind,error) {
            return structVal.Field(i).Kind(),nil
        },
    )
}

// TODO -test
// Returns an iterator that provides the struct field tags if a struct is
// is supplied as an argument, returns an error otherwise. As a special case, 
// if a reflect.Value is passed to this function it will return the field 
// tags of the struct it contains or the field tags of the struct that it
// points to if the reflect.Value contains a pointer to a struct.
func StructFieldTags[T any, S reflect.Value | *T](s S) iter.Iter[reflect.StructTag] {
    structVal,err:=homogonizeStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[reflect.StructTag]("",err,1)
    }
    return iter.SequentialElems[reflect.StructTag](
        structVal.NumField(),
        func(i int) (reflect.StructTag,error) {
            return structVal.Type().Field(i).Tag,nil
        },
    )
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
func StructFieldInfo[T any, S reflect.Value | *T](
    s S, 
    keepVal bool,
) iter.Iter[FieldInfo] {
    structVal,err:=homogonizeStructVal[T,S](s)
    if err!=nil {
        return iter.ValElem[FieldInfo](FieldInfo{},err,1)
    }
    return iter.SequentialElems[FieldInfo](
        structVal.NumField(),
        func(i int) (FieldInfo, error) {
            n:=structVal.Type().Field(i).Name
            f:=structVal.Field(i)
            return FieldInfo{
                Name: n,
                Tag: structVal.Type().Field(i).Tag,
                ValInfo: ValInfo{
                    Type: f.Type(),
                    Kind: f.Kind(),
                    Val: func() (any, bool) {
                        if keepVal {
                            return f.Interface(),true
                        }
                        return nil,false
                    },
                    Pntr: func() (any, error) {
                        if f.CanAddr() {
                            return f.Addr().Interface(),nil
                        }
                        return reflect.Value{},getInaddressableFieldError(n)
                    },
                },
            }, nil
        },
    )
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
    keepVal bool,
) iter.Iter[FieldInfo] {
    if err:=structValError[T,S](s); err!=nil {
        return iter.ValElem[FieldInfo](FieldInfo{},err,1)
    }
    return iter.Recurse[FieldInfo](
        StructFieldInfo[T,S](s,keepVal),
        func(v FieldInfo) bool { return v.Kind==reflect.Struct },
        func(v FieldInfo) iter.Iter[FieldInfo] {
            if p,err:=v.Pntr(); err==nil {
                return StructFieldInfo[T,reflect.Value](reflect.ValueOf(p),keepVal)
            } else {
                return iter.ValElem[FieldInfo](FieldInfo{},err,1)
            }
        },
    )
}
