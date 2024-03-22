package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/test"
	"github.com/barbell-math/util/customerr"
)

func TestIsArrayVal(t *testing.T){
    v:=[]int{}
    test.False(IsArrayVal[[]int](&v),t)
    v2:=reflect.ValueOf(v)
    test.False(IsArrayVal[[]int](v2),t)
    v2=reflect.ValueOf(&v)
    test.False(IsArrayVal[[]int](v2),t)
    a:=[3]int{}
    test.True(IsArrayVal[[3]int](&a),t)
    a2:=reflect.ValueOf(a)
    test.True(IsArrayVal[[3]int](a2),t)
    a2=reflect.ValueOf(&a)
    test.True(IsArrayVal[[3]int](a2),t)
}

func TestNonArrayElemVals(t *testing.T){
    v:=0
    vals,err:=ArrayElemVals[int](&v).Collect()
    test.Eq(0,len(vals),t)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemValsReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=ArrayElemVals[reflect.Value](v2).Collect()
    test.Eq(0,len(vals),t)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemValsReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=ArrayElemVals[reflect.Value](v2).Collect()
    test.Eq(0,len(vals),t)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestArrayElemVals(t *testing.T){
    v:=[3]int{0,1,2}
    vals,err:=ArrayElemVals[[3]int](&v).Collect()
    test.Eq(3,len(vals),t)
    test.Nil(err,t)
    for i,v:=range(vals) {
        test.Eq(i,v.(int),t)
    }
}

func TestArrayElemValsReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    vals,err:=ArrayElemVals[reflect.Value](v2).Collect()
    test.Eq(3,len(vals),t)
    test.Nil(err,t)
    for i,v:=range(vals) {
        test.Eq(i,v.(int),t)
    }
}

func TestArrayElemValsReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    vals,err:=ArrayElemVals[reflect.Value](v2).Collect()
    test.Eq(3,len(vals),t)
    test.Nil(err,t)
    for i,v:=range(vals) {
        test.Eq(i,v.(int),t)
    }
}

func TestNonArrayElemPntrs(t *testing.T){
    v:=0
    vals,err:=ArrayElemPntrs[int](&v).Collect()
    test.Eq(0,len(vals),t)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemPntrsReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=ArrayElemPntrs[reflect.Value](v2).Collect()
    test.Eq(0,len(vals),t)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemPntrsReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=ArrayElemPntrs[reflect.Value](v2).Collect()
    test.Eq(0,len(vals),t)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestArrayElemPntrs(t *testing.T){
    v:=[3]int{0,1,2}
    vals,err:=ArrayElemPntrs[[3]int](&v).Collect()
    test.Eq(3,len(vals),t)
    test.Nil(err,t)
    for i,iterV:=range(vals) {
        test.Eq(&v[i],iterV.(*int),t)
    }
}

func TestArrayElemPntrsReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    vals,err:=ArrayElemPntrs[reflect.Value](v2).Collect()
    test.Eq(0,len(vals),t)
    test.ContainsError(InAddressableField,err,t)
}

func TestArrayElemPntrsReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    vals,err:=ArrayElemPntrs[reflect.Value](v2).Collect()
    test.Eq(3,len(vals),t)
    test.Nil(err,t)
    for i,iterV:=range(vals) {
        test.Eq(&v[i],iterV.(*int),t)
    }
}

func TestNonArrayElemType(t *testing.T){
    v:=0
    _,err:=ArrayElemType[int](&v)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemTypeReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=ArrayElemType[reflect.Value](v2)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemTypeReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=ArrayElemType[reflect.Value](v2)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestArrayElemType(t *testing.T){
    v:=[3]int{0,1,2}
    _type,err:=ArrayElemType[[3]int](&v)
    test.Eq(reflect.TypeOf(int(0)),_type,t)
    test.Nil(err,t)
}

func TestArrayElemTypeReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    _type,err:=ArrayElemType[[3]int](v2)
    test.Eq(reflect.TypeOf(int(0)),_type,t)
    test.Nil(err,t)
}

func TestArrayElemTypeReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    _type,err:=ArrayElemType[[3]int](v2)
    test.Eq(reflect.TypeOf(int(0)),_type,t)
    test.Nil(err,t)
}

func TestNonArrayElemKind(t *testing.T){
    v:=0
    _,err:=ArrayElemKind[int](&v)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemKindReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=ArrayElemKind[reflect.Value](v2)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemKindReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=ArrayElemKind[reflect.Value](v2)
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestArrayElemKind(t *testing.T){
    v:=[3]int{0,1,2}
    _type,err:=ArrayElemKind[[3]int](&v)
    test.Eq(reflect.Int,_type,t)
    test.Nil(err,t)
}

func TestArrayElemKindReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    _type,err:=ArrayElemKind[[3]int](v2)
    test.Eq(reflect.Int,_type,t)
    test.Nil(err,t)
}

func TestArrayElemKindReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    _type,err:=ArrayElemKind[[3]int](v2)
    test.Eq(reflect.Int,_type,t)
    test.Nil(err,t)
}

func TestNonArrayElemInfo(t *testing.T){
    v:=0
    _,err:=ArrayElemInfo[int](&v,true).Collect()
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemInfoReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=ArrayElemInfo[reflect.Value](v2,true).Collect()
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayElemInfoReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=ArrayElemInfo[reflect.Value](v2,true).Collect()
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestArrayElemInfo(t *testing.T){
    v:=[3]int{0,1,2}
    info,err:=ArrayElemInfo[[3]int](&v,true).Collect()
    test.Nil(err,t)
    test.Eq(3,len(info),t)
    for i,iterV:=range(v) {
        _v,ok:=info[i].Val()
        test.Eq(iterV,_v.(int),t)
        test.True(ok,t)
        test.Eq(reflect.TypeOf(int(0)),info[i].Type,t)
        test.Eq(reflect.Int,info[i].Kind,t)
        pntr,err:=info[i].Pntr()
        test.Eq(&v[i],pntr.(*int),t)
        test.Nil(err,t)
    }
}

func TestArrayElemInfoReflectVal(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(v)
    info,err:=ArrayElemInfo[[3]int](v2,true).Collect()
    test.Nil(err,t)
    test.Eq(3,len(info),t)
    for i,iterV:=range(v) {
        _v,ok:=info[i].Val()
        test.Eq(iterV,_v.(int),t)
        test.True(ok,t)
        test.Eq(reflect.TypeOf(int(0)),info[i].Type,t)
        test.Eq(reflect.Int,info[i].Kind,t)
        pntr,err:=info[i].Pntr()
        test.Nil(pntr,t)
        test.ContainsError(InAddressableField,err,t)
    }
}

func TestArrayElemInfoReflectValPntr(t *testing.T){
    v:=[3]int{0,1,2}
    v2:=reflect.ValueOf(&v)
    info,err:=ArrayElemInfo[[3]int](v2,true).Collect()
    test.Nil(err,t)
    test.Eq(3,len(info),t)
    for i,iterV:=range(v) {
        _v,ok:=info[i].Val()
        test.Eq(iterV,_v.(int),t)
        test.True(ok,t)
        test.Eq(reflect.TypeOf(int(0)),info[i].Type,t)
        test.Eq(reflect.Int,info[i].Kind,t)
        pntr,err:=info[i].Pntr()
        test.Eq(&v[i],pntr.(*int),t)
        test.Nil(err,t)
    }
}

func TestNonArrayRecursiveElemInfo(t *testing.T){
    v:=0
    _,err:=RecursiveArrayElemInfo[int](&v,true).Collect()
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayRecursiveElemInfoReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=RecursiveArrayElemInfo[reflect.Value](v2,true).Collect()
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestNonArrayRecursiveElemInfoReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=RecursiveArrayElemInfo[reflect.Value](v2,true).Collect()
    test.ContainsError(customerr.IncorrectType,err,t)
}

func TestRecursiveArrayElemInfo(t *testing.T){
    v:=[3][3]int{[3]int{1,2,3},[3]int{5,6,7},[3]int{9,10,11}}
    vals,err:=RecursiveArrayElemInfo[[3][3]int](&v,true).Collect()
    test.Nil(err,t)
    test.Eq(12,len(vals),t)
    for i,iterV:=range(vals) {
        if i==0 || i==4 || i==8 {
            _v,ok:=iterV.Val()
            test.Eq([3]int{i+1,i+2,i+3},_v.([3]int),t)
            test.True(ok,t)
            test.Eq(reflect.TypeOf([3]int{}),iterV.Type,t)
            test.Eq(reflect.Array,iterV.Kind,t)
            p,err:=iterV.Pntr()
            test.Nil(err,t)
            test.Eq(&v[i/4],p,t)
        } else {
            _v,ok:=iterV.Val()
            test.Eq(i,_v.(int),t)
            test.True(ok,t)
            test.Eq(reflect.TypeOf(int(0)),iterV.Type,t)
            test.Eq(reflect.Int,iterV.Kind,t)
            p,err:=iterV.Pntr()
            test.Nil(err,t)
            test.Eq(&v[i/4][i%4-1],p,t)
        }
    }
}

func TestRecursiveArrayElemInfoReflectVal(t *testing.T){
    v:=[3][3]int{[3]int{1,2,3},[3]int{5,6,7},[3]int{9,10,11}}
    v2:=reflect.ValueOf(v)
    vals,err:=RecursiveArrayElemInfo[[3][3]int](v2,true).Collect()
    test.Eq(1,len(vals),t)
    test.ContainsError(InAddressableField,err,t)
    for i,iterV:=range(vals) {
        v,ok:=iterV.Val()
        test.Eq([3]int{4*i+1,4*i+2,4*i+3},v.([3]int),t)
        test.True(ok,t)
        test.Eq(reflect.TypeOf([3]int{}),iterV.Type,t)
        test.Eq(reflect.Array,iterV.Kind,t)
        p,err:=iterV.Pntr()
        test.Nil(p,t)
        test.ContainsError(InAddressableField,err,t)
    }
}

func TestRecursiveArrayElemInfoReflectValPntr(t *testing.T){
    v:=[3][3]int{[3]int{1,2,3},[3]int{5,6,7},[3]int{9,10,11}}
    v2:=reflect.ValueOf(&v)
    vals,err:=RecursiveArrayElemInfo[[3][3]int](v2,true).Collect()
    test.Nil(err,t)
    test.Eq(12,len(vals),t)
    for i,iterV:=range(vals) {
        if i==0 || i==4 || i==8 {
            _v,ok:=iterV.Val()
            test.Eq([3]int{i+1,i+2,i+3},_v.([3]int),t)
            test.True(ok,t)
            test.Eq(reflect.TypeOf([3]int{}),iterV.Type,t)
            test.Eq(reflect.Array,iterV.Kind,t)
            p,err:=iterV.Pntr()
            test.Nil(err,t)
            test.Eq(&v[i/4],p,t)
        } else {
            _v,ok:=iterV.Val()
            test.Eq(i,_v.(int),t)
            test.True(ok,t)
            test.Eq(reflect.TypeOf(int(0)),iterV.Type,t)
            test.Eq(reflect.Int,iterV.Kind,t)
            p,err:=iterV.Pntr()
            test.Nil(err,t)
            test.Eq(&v[i/4][i%4-1],p,t)
        }
    }
}
