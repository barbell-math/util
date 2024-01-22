package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/test"
	customerr "github.com/barbell-math/util/err"
)

func TestIsSliceVal(t *testing.T){
    v:=[3]int{}
    test.BasicTest(false,IsSliceVal[[3]int](&v),
        "IsSliceVal returned false positive.",t,
    )
    v2:=reflect.ValueOf(v)
    test.BasicTest(false,IsSliceVal[int](v2),
        "IsSliceVal returned false positive.",t,
    )
    v2=reflect.ValueOf(&v)
    test.BasicTest(false,IsSliceVal[int](v2),
        "IsSliceVal returned false positive.",t,
    )
    a:=[]int{}
    test.BasicTest(true,IsSliceVal[[]int](&a),
        "IsSliceVal returned false negative.",t,
    )
    a2:=reflect.ValueOf(a)
    test.BasicTest(true,IsSliceVal[[]int](a2),
        "IsSliceVal returned false negative.",t,
    )
    a2=reflect.ValueOf(&a)
    test.BasicTest(true,IsSliceVal[[]int](a2),
        "IsSliceVal returned false negative.",t,
    )
}

func TestNonSliceElemVals(t *testing.T){
    v:=0
    vals,err:=SliceElemVals[int](&v).Collect()
    test.BasicTest(0,len(vals),
        "SliceElemVals returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemValsReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=SliceElemVals[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "SliceElemVals returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemValsReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=SliceElemVals[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "SliceElemVals returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestSliceElemVals(t *testing.T){
    v:=[]int{0,1,2}
    vals,err:=SliceElemVals[[]int](&v).Collect()
    test.BasicTest(3,len(vals),
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    for i,v:=range(vals) {
        test.BasicTest(i,v.(int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestSliceElemValsReflectVal(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(v)
    vals,err:=SliceElemVals[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    for i,v:=range(vals) {
        test.BasicTest(i,v.(int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestSliceElemValsReflectValPntr(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    vals,err:=SliceElemVals[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    for i,v:=range(vals) {
        test.BasicTest(i,v.(int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestNonSliceElemPntrs(t *testing.T){
    v:=0
    vals,err:=SliceElemPntrs[int](&v).Collect()
    test.BasicTest(0,len(vals),
        "SliceElemVals returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemPntrsReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=SliceElemPntrs[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "SliceElemVals returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemPntrsReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=SliceElemPntrs[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "SliceElemVals returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestSliceElemPntrs(t *testing.T){
    v:=[]int{0,1,2}
    vals,err:=SliceElemPntrs[[]int](&v).Collect()
    test.BasicTest(3,len(vals),
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    for i,iterV:=range(vals) {
        test.BasicTest(&v[i],iterV.(*int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestSliceElemPntrsReflectVal(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(v)
    vals,err:=SliceElemPntrs[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    for i,iterV:=range(vals) {
        test.BasicTest(&v[i],iterV.(*int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestSliceElemPntrsReflectValPntr(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    vals,err:=SliceElemPntrs[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    for i,iterV:=range(vals) {
        test.BasicTest(&v[i],iterV.(*int),
            "A value in the array was incorrect.",t,
        )
    }
}

func TestNonSliceElemType(t *testing.T){
    v:=0
    _,err:=SliceElemType[int](&v)
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemTypeReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=SliceElemType[reflect.Value](v2)
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemTypeReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=SliceElemType[reflect.Value](v2)
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestSliceElemType(t *testing.T){
    v:=[]int{0,1,2}
    _type,err:=SliceElemType[[]int](&v)
    test.BasicTest(reflect.TypeOf(int(0)),_type,
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
}

func TestSliceElemTypeReflectVal(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(v)
    _type,err:=SliceElemType[[]int](v2)
    test.BasicTest(reflect.TypeOf(int(0)),_type,
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
}

func TestSliceElemTypeReflectValPntr(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    _type,err:=SliceElemType[[]int](v2)
    test.BasicTest(reflect.TypeOf(int(0)),_type,
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
}

func TestNonSliceElemKind(t *testing.T){
    v:=0
    _,err:=SliceElemKind[int](&v)
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemKindReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=SliceElemKind[reflect.Value](v2)
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemKindReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=SliceElemKind[reflect.Value](v2)
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestSliceElemKind(t *testing.T){
    v:=[]int{0,1,2}
    _type,err:=SliceElemKind[[]int](&v)
    test.BasicTest(reflect.Int,_type,
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
}

func TestSliceElemKindReflectVal(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(v)
    _type,err:=SliceElemKind[[]int](v2)
    test.BasicTest(reflect.Int,_type,
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
}

func TestSliceElemKindReflectValPntr(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    _type,err:=SliceElemKind[[]int](v2)
    test.BasicTest(reflect.Int,_type,
        "SliceElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
}

func TestNonSliceElemInfo(t *testing.T){
    v:=0
    _,err:=SliceElemInfo[int](&v,true).Collect()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemInfoReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=SliceElemInfo[reflect.Value](v2,true).Collect()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceElemInfoReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=SliceElemInfo[reflect.Value](v2,true).Collect()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestSliceElemInfo(t *testing.T){
    v:=[]int{0,1,2}
    info,err:=SliceElemInfo[[]int](&v,true).Collect()
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    test.BasicTest(3,len(info),
        "SliceElemVals returned the wrong number of elements.",t,
    )
    for i,iterV:=range(v) {
        _v,ok:=info[i].Val()
        test.BasicTest(iterV,_v.(int),
            "SliceElemInfo returned an incorrect value.",t,
        )
        test.BasicTest(true,ok,
            "SliceElemInfo returned an incorrect value.",t,
        )
        test.BasicTest(reflect.TypeOf(int(0)),info[i].Type,
            "SliceElemInfo returned an incorrect type.",t,
        )
        test.BasicTest(reflect.Int,info[i].Kind,
            "SliceElemInfo returned an incorrect type.",t,
        )
        pntr,err:=info[i].Pntr()
        test.BasicTest(&v[i],pntr.(*int),
            "SliceElemInfo returned an incorrect pntr.",t,
        )
        test.BasicTest(nil,err,
            "SliceElemInfo returned an error when accessing a pointer value when it should not have.",t,    
        )
    }
}

func TestSliceElemInfoReflectVal(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(v)
    info,err:=SliceElemInfo[[]int](v2,true).Collect()
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    test.BasicTest(3,len(info),
        "SliceElemVals returned the wrong number of elements.",t,
    )
    for i,iterV:=range(v) {
        _v,ok:=info[i].Val()
        test.BasicTest(iterV,_v.(int),
            "SliceElemInfo returned an incorrect value.",t,
        )
        test.BasicTest(true,ok,
            "SliceElemInfo returned an incorrect value.",t,
        )
        test.BasicTest(reflect.TypeOf(int(0)),info[i].Type,
            "SliceElemInfo returned an incorrect type.",t,
        )
        test.BasicTest(reflect.Int,info[i].Kind,
            "SliceElemInfo returned an incorrect type.",t,
        )
        pntr,err:=info[i].Pntr()
        test.BasicTest(&v[i],pntr.(*int),
            "SliceElemInfo returned an incorrect pntr.",t,
        )
        test.BasicTest(nil,err,
            "SliceElemInfo returned an error when accessing a pointer value when it should not have.",t,    
        )
    }
}

func TestSliceElemInfoReflectValPntr(t *testing.T){
    v:=[]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    info,err:=SliceElemInfo[[]int](v2,true).Collect()
    test.BasicTest(nil,err,
        "SliceElemVals returned an error when it should not have.",t,    
    )
    test.BasicTest(3,len(info),
        "SliceElemVals returned the wrong number of elements.",t,
    )
    for i,iterV:=range(v) {
        _v,ok:=info[i].Val()
        test.BasicTest(iterV,_v.(int),
            "SliceElemInfo returned an incorrect value.",t,
        )
        test.BasicTest(true,ok,
            "SliceElemInfo returned an incorrect value.",t,
        )
        test.BasicTest(reflect.TypeOf(int(0)),info[i].Type,
            "SliceElemInfo returned an incorrect type.",t,
        )
        test.BasicTest(reflect.Int,info[i].Kind,
            "SliceElemInfo returned an incorrect type.",t,
        )
        pntr,err:=info[i].Pntr()
        test.BasicTest(&v[i],pntr.(*int),
            "SliceElemInfo returned an incorrect pntr.",t,
        )
        test.BasicTest(nil,err,
            "SliceElemInfo returned an error when accessing a pointer value when it should not have.",t,    
        )
    }
}

func TestNonSliceRecursiveElemInfo(t *testing.T){
    v:=0
    _,err:=RecursiveSliceElemInfo[int](&v,true).Collect()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceRecursiveElemInfoReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=RecursiveSliceElemInfo[reflect.Value](v2,true).Collect()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonSliceRecursiveElemInfoReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=RecursiveSliceElemInfo[reflect.Value](v2,true).Collect()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "SliceElemVals returned an incorrect error.",t,
        )
    }
}

func TestRecursiveSliceElemInfo(t *testing.T){
    v:=[][]int{[]int{1,2,3},[]int{5,6,7},[]int{9,10,11}}
    vals,err:=RecursiveSliceElemInfo[[][]int](&v,true).Collect()
    test.BasicTest(nil,err,
        "RecursiveSliceElemInfo returned an error when it should not have.",t,
    )
    test.BasicTest(12,len(vals),
        "The wrong number of values were returned.",t,
    )
    for i,iterV:=range(vals) {
        if i==0 || i==4 || i==8 {
            tmp:=[]int{i+1,i+2,i+3}
            for j,tmpV:=range(tmp) {
                _v,ok:=iterV.Val()
                test.BasicTest(tmpV,_v.([]int)[j],
                    "Recurse did not return the proper root node.",t,
                )
                test.BasicTest(true,ok,
                    "Recurse did not return the proper root node.",t,
                )
            }
            test.BasicTest(reflect.TypeOf([]int{}),iterV.Type,
                "Recurse did not return the proper type of a root node.",t,
            )
            test.BasicTest(reflect.Slice,iterV.Kind,
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
            _v,ok:=iterV.Val()
            test.BasicTest(i,_v.(int),
                "Recurse did not return the proper root node.",t,
            )
            test.BasicTest(true,ok,
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

func TestRecursiveSliceElemInfoReflectVal(t *testing.T){
    v:=[][]int{[]int{1,2,3},[]int{5,6,7},[]int{9,10,11}}
    v2:=reflect.ValueOf(v)
    vals,err:=RecursiveSliceElemInfo[[][]int](v2,true).Collect()
    test.BasicTest(nil,err,
        "RecursiveSliceElemInfo returned an error when it should not have.",t,
    )
    test.BasicTest(12,len(vals),
        "The wrong number of values were returned.",t,
    )
    for i,iterV:=range(vals) {
        if i==0 || i==4 || i==8 {
            tmp:=[]int{i+1,i+2,i+3}
            for j,tmpV:=range(tmp) {
                _v,ok:=iterV.Val()
                test.BasicTest(tmpV,_v.([]int)[j],
                    "Recurse did not return the proper root node.",t,
                )
                test.BasicTest(true,ok,
                    "Recurse did not return the proper root node.",t,
                )
            }
            test.BasicTest(reflect.TypeOf([]int{}),iterV.Type,
                "Recurse did not return the proper type of a root node.",t,
            )
            test.BasicTest(reflect.Slice,iterV.Kind,
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
            _v,ok:=iterV.Val()
            test.BasicTest(i,_v.(int),
                "Recurse did not return the proper root node.",t,
            )
            test.BasicTest(true,ok,
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

func TestRecursiveSliceElemInfoReflectValPntr(t *testing.T){
    v:=[][]int{[]int{1,2,3},[]int{5,6,7},[]int{9,10,11}}
    v2:=reflect.ValueOf(&v)
    vals,err:=RecursiveSliceElemInfo[[][]int](v2,true).Collect()
    test.BasicTest(nil,err,
        "RecursiveSliceElemInfo returned an error when it should not have.",t,
    )
    test.BasicTest(12,len(vals),
        "The wrong number of values were returned.",t,
    )
    for i,iterV:=range(vals) {
        if i==0 || i==4 || i==8 {
            tmp:=[]int{i+1,i+2,i+3}
            for j,tmpV:=range(tmp) {
                _v,ok:=iterV.Val()
                test.BasicTest(tmpV,_v.([]int)[j],
                    "Recurse did not return the proper root node.",t,
                )
                test.BasicTest(true,ok,
                    "Recurse did not return the proper root node.",t,
                )
            }
            test.BasicTest(reflect.TypeOf([]int{}),iterV.Type,
                "Recurse did not return the proper type of a root node.",t,
            )
            test.BasicTest(reflect.Slice,iterV.Kind,
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
            _v,ok:=iterV.Val()
            test.BasicTest(i,_v.(int),
                "Recurse did not return the proper root node.",t,
            )
            test.BasicTest(true,ok,
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
