package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestIsArrayVal(t *testing.T){
    v:=[]int{}
    test.BasicTest(false,IsArrayVal[[]int](&v),
        "IsArrayVal returned false positive.",t,
    )
    v2:=reflect.ValueOf(v)
    test.BasicTest(false,IsArrayVal[[]int](v2),
        "IsArrayVal returned false positive.",t,
    )
    v2=reflect.ValueOf(&v)
    test.BasicTest(false,IsArrayVal[[]int](v2),
        "IsArrayVal returned false positive.",t,
    )
    a:=[3]int{}
    test.BasicTest(true,IsArrayVal[[3]int](&a),
        "IsArrayVal returned false negative.",t,
    )
    a2:=reflect.ValueOf(a)
    test.BasicTest(true,IsArrayVal[[3]int](a2),
        "IsArrayVal returned false negative.",t,
    )
    a2=reflect.ValueOf(&a)
    test.BasicTest(true,IsArrayVal[[3]int](a2),
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

func TestArrayElemInfoReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    info,err:=ArrayElemInfo[[3]int](v2).Collect()
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
        test.BasicTest(nil,pntr,
            "ArrayElemInfo returned an incorrect pntr.",t,
        )
        if !IsInAddressableField(err) {
            test.FormatError(InAddressableField(""),err,
                "Getting a pointer to an inadressable field returned the wrong error.",t,
            )
        }
    }
}

func TestArrayElemInfoReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    info,err:=ArrayElemInfo[[3]int](v2).Collect()
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

func TestNonArrayRecursiveElemInfo(t *testing.T){
    v:=0
    _,err:=RecursiveArrayElemInfo[int](&v).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayRecursiveElemInfoReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=RecursiveArrayElemInfo[reflect.Value](v2).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonArrayRecursiveElemInfoReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=RecursiveArrayElemInfo[reflect.Value](v2).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "ArrayElemVals returned an incorrect error.",t,
        )
    }
}

func TestRecursiveArrayElemInfo(t *testing.T){
    v:=[3][3]int{[3]int{1,2,3},[3]int{5,6,7},[3]int{9,10,11}}
    vals,err:=RecursiveArrayElemInfo[[3][3]int](&v).Collect()
    test.BasicTest(nil,err,
        "RecursiveArrayElemInfo returned an error when it should not have.",t,
    )
    test.BasicTest(12,len(vals),
        "The wrong number of values were returned.",t,
    )
    for i,iterV:=range(vals) {
        if i==0 || i==4 || i==8 {
            test.BasicTest([3]int{i+1,i+2,i+3},iterV.Val.([3]int),
                "Recurse did not return the proper root node.",t,
            )
            test.BasicTest(reflect.TypeOf([3]int{}),iterV.Type,
                "Recurse did not return the proper type of a root node.",t,
            )
            test.BasicTest(reflect.Array,iterV.Kind,
                "Recurse did not return the proper kind of a root node.",t,
            )
            p,err:=iterV.Pntr()
            test.BasicTest(nil,err,
                "Pointer to a root node did not return the correct value.",t,
            )
            test.BasicTest(&v[i/4],p,
                "Pointer to a root node did not return the correct value.",t,
            )
        } else {
            test.BasicTest(i,iterV.Val.(int),
                "Recurse did not return the proper root node.",t,
            )
            test.BasicTest(reflect.TypeOf(int(0)),iterV.Type,
                "Recurse did not return the proper type of a root node.",t,
            )
            test.BasicTest(reflect.Int,iterV.Kind,
                "Recurse did not return the proper kind of a root node.",t,
            )
            p,err:=iterV.Pntr()
            test.BasicTest(nil,err,
                "Pointer to a root node did not return the correct value.",t,
            )
            test.BasicTest(&v[i/4][i%4-1],p,
                "Pointer to a root node did not return the correct value.",t,
            )
        }
    }
}

func TestRecursiveArrayElemInfoReflectVal(t *testing.T){
    v:=[3][3]int{[3]int{1,2,3},[3]int{5,6,7},[3]int{9,10,11}}
    v2:=reflect.ValueOf(v)
    vals,err:=RecursiveArrayElemInfo[[3][3]int](v2).Collect()
    test.BasicTest(1,len(vals),
        "The wrong number of values were returned.",t,
    )
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "Field was not inadressable when it should have been.",t,
        )
    }
    for i,iterV:=range(vals) {
        test.BasicTest([3]int{4*i+1,4*i+2,4*i+3},iterV.Val.([3]int),
            "Recurse did not return the proper root node.",t,
        )
        test.BasicTest(reflect.TypeOf([3]int{}),iterV.Type,
            "Recurse did not return the proper type of a root node.",t,
        )
        test.BasicTest(reflect.Array,iterV.Kind,
            "Recurse did not return the proper kind of a root node.",t,
        )
        p,err:=iterV.Pntr()
        test.BasicTest(nil,p,
            "Pointer to a root node did not return the correct value.",t,
        )
        if !IsInAddressableField(err) {
            test.FormatError(InAddressableField(""),err,
                "Field was not inadressable when it should have been.",t,
            )
        }
    }
}

func TestRecursiveArrayElemInfoReflectValPntr(t *testing.T){
    v:=[3][3]int{[3]int{1,2,3},[3]int{5,6,7},[3]int{9,10,11}}
    v2:=reflect.ValueOf(&v)
    vals,err:=RecursiveArrayElemInfo[[3][3]int](v2).Collect()
    test.BasicTest(nil,err,
        "RecursiveArrayElemInfo returned an error when it should not have.",t,
    )
    test.BasicTest(12,len(vals),
        "The wrong number of values were returned.",t,
    )
    for i,iterV:=range(vals) {
        if i==0 || i==4 || i==8 {
            test.BasicTest([3]int{i+1,i+2,i+3},iterV.Val.([3]int),
                "Recurse did not return the proper root node.",t,
            )
            test.BasicTest(reflect.TypeOf([3]int{}),iterV.Type,
                "Recurse did not return the proper type of a root node.",t,
            )
            test.BasicTest(reflect.Array,iterV.Kind,
                "Recurse did not return the proper kind of a root node.",t,
            )
            p,err:=iterV.Pntr()
            test.BasicTest(nil,err,
                "Pointer to a root node did not return the correct value.",t,
            )
            test.BasicTest(&v[i/4],p,
                "Pointer to a root node did not return the correct value.",t,
            )
        } else {
            test.BasicTest(i,iterV.Val.(int),
                "Recurse did not return the proper root node.",t,
            )
            test.BasicTest(reflect.TypeOf(int(0)),iterV.Type,
                "Recurse did not return the proper type of a root node.",t,
            )
            test.BasicTest(reflect.Int,iterV.Kind,
                "Recurse did not return the proper kind of a root node.",t,
            )
            p,err:=iterV.Pntr()
            test.BasicTest(nil,err,
                "Pointer to a root node did not return the correct value.",t,
            )
            test.BasicTest(&v[i/4][i%4-1],p,
                "Pointer to a root node did not return the correct value.",t,
            )
        }
    }
}
