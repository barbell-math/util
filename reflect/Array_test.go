package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestIsArrayVal(t *testing.T){
    v:=[]int{}
    test.BasicTest(false,IsArrayVal(&v),
        "IsArrayVal returned false positive.",t,
    )
    v2:=reflect.ValueOf(v)
    test.BasicTest(false,IsArrayVal(&v2),
        "IsArrayVal returned false positive.",t,
    )
    v2=reflect.ValueOf(&v)
    test.BasicTest(false,IsArrayVal(&v2),
        "IsArrayVal returned false positive.",t,
    )
    a:=[3]int{}
    test.BasicTest(true,IsArrayVal(&a),
        "IsArrayVal returned false negative.",t,
    )
    a2:=reflect.ValueOf(a)
    test.BasicTest(true,IsArrayVal(&a2),
        "IsArrayVal returned false negative.",t,
    )
    a2=reflect.ValueOf(&a)
    test.BasicTest(true,IsArrayVal(&a2),
        "IsArrayVal returned false negative.",t,
    )
}

func TestNonArrayElemVals(t *testing.T){
    v:=0
    vals,err:=ArrayElemVals[int](&v).Collect()
    test.BasicTest(0,len(vals),
        "ArrayElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemValsReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=ArrayElemVals[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "ArrayElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemValsReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=ArrayElemVals[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "ArrayElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestArrayElemVals(t *testing.T){
    v:=[3]int{0,1,2}
    vals,err:=ArrayElemVals[[3]int](&v).Collect()
    test.BasicTest(3,len(vals),
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
    for i,v:=range(vals) {
        test.BasicTest(i,v.(int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestArrayElemValsReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    vals,err:=ArrayElemVals[reflect.Value](&v2).Collect()
    test.BasicTest(3,len(vals),
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
    for i,v:=range(vals) {
        test.BasicTest(i,v.(int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestArrayElemValsReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    vals,err:=ArrayElemVals[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
    for i,v:=range(vals) {
        test.BasicTest(i,v.(int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestNonArrayElemPntrs(t *testing.T){
    v:=0
    vals,err:=ArrayElemPntrs[int](&v).Collect()
    test.BasicTest(0,len(vals),
        "ArrayElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemPntrsReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=ArrayElemPntrs[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "ArrayElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemPntrsReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=ArrayElemPntrs[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "ArrayElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestArrayElemPntrs(t *testing.T){
    v:=[3]int{0,1,2}
    vals,err:=ArrayElemPntrs[[3]int](&v).Collect()
    test.BasicTest(3,len(vals),
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
    for i,iterV:=range(vals) {
        test.BasicTest(&v[i],iterV.(*int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestArrayElemPntrsReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    vals,err:=ArrayElemPntrs[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "ArrayElemVals the incorrect number of values.",t,
    )
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "Field was addressable when it shouldn't have been.",t,
        )
    }
}

func TestArrayElemPntrsReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    vals,err:=ArrayElemPntrs[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
    for i,iterV:=range(vals) {
        test.BasicTest(&v[i],iterV.(*int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestNonArrayElemType(t *testing.T){
    v:=0
    _,err:=ArrayElemType[int](&v)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemTypeReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=ArrayElemType[reflect.Value](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemTypeReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=ArrayElemType[reflect.Value](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestArrayElemType(t *testing.T){
    v:=[3]int{0,1,2}
    _type,err:=ArrayElemType[[3]int](&v)
    test.BasicTest(reflect.TypeOf(int(0)),_type,
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
}

func TestArrayElemTypeReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    _type,err:=ArrayElemType[[3]int](v2)
    test.BasicTest(reflect.TypeOf(int(0)),_type,
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
}

func TestArrayElemTypeReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    _type,err:=ArrayElemType[[3]int](v2)
    test.BasicTest(reflect.TypeOf(int(0)),_type,
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
}

func TestNonArrayElemKind(t *testing.T){
    v:=0
    _,err:=ArrayElemKind[int](&v)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemKindReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=ArrayElemKind[reflect.Value](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemKindReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=ArrayElemKind[reflect.Value](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestArrayElemKind(t *testing.T){
    v:=[3]int{0,1,2}
    _type,err:=ArrayElemKind[[3]int](&v)
    test.BasicTest(reflect.Int,_type,
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
}

func TestArrayElemKindReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    _type,err:=ArrayElemKind[[3]int](v2)
    test.BasicTest(reflect.Int,_type,
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
}

func TestArrayElemKindReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    _type,err:=ArrayElemKind[[3]int](v2)
    test.BasicTest(reflect.Int,_type,
        "ArrayElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
}

func TestNonArrayElemInfo(t *testing.T){
    v:=0
    _,err:=ArrayElemInfo[int](&v).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemInfoReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=ArrayElemInfo[reflect.Value](v2).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayElemInfoReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=ArrayElemInfo[reflect.Value](v2).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestArrayElemInfo(t *testing.T){
    v:=[3]int{0,1,2}
    info,err:=ArrayElemInfo[[3]int](&v).Collect()
    test.BasicTest(nil,err,
        "ArrayElemVals returned an error when it should not have.",t,    
    )
    test.BasicTest(3,len(info),
        "ArrayElemVals returned the wrong number of elements.",t,
    )
    for i,iterV:=range(v) {
        test.BasicTest(iterV,info[i].Val.(int),
            "ArrayElemInfo returned an incorrect value.",t,
        )
        test.BasicTest(reflect.TypeOf(int(0)),info[i].Type,
            "ArrayElemInfo returned an incorrect type.",t,
        )
        test.BasicTest(reflect.Int,info[i].Kind,
            "ArrayElemInfo returned an incorrect type.",t,
        )
        pntr,err:=info[i].Pntr()
        test.BasicTest(&v[i],pntr.(*int),
            "ArrayElemInfo returned an incorrect pntr.",t,
        )
        test.BasicTest(nil,err,
            "ArrayElemInfo returned an error when accessing a pointer value when it should not have.",t,    
        )
    }
}

// TODO -complete elem info testing
// func TestArrayElemKindReflectVal(t *testing.T){
//     v:=[3]int{0,1,2}
//     v2:=reflect.ValueOf(v)
//     _type,err:=ArrayElemKind[[3]int](v2)
//     test.BasicTest(reflect.Int,_type,
//         "ArrayElemVals the incorrect number of values.",t,
//     )
//     test.BasicTest(nil,err,
//         "ArrayElemVals returned an error when it should not have.",t,    
//     )
// }
// 
// func TestArrayElemKindReflectValPntr(t *testing.T){
//     v:=[3]int{0,1,2}
//     v2:=reflect.ValueOf(&v)
//     _type,err:=ArrayElemKind[[3]int](v2)
//     test.BasicTest(reflect.Int,_type,
//         "ArrayElemVals the incorrect number of values.",t,
//     )
//     test.BasicTest(nil,err,
//         "ArrayElemVals returned an error when it should not have.",t,    
//     )
// }
