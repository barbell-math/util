package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/test"
)

func TestIsMapVal(t *testing.T){
    v:=0
    test.BasicTest(false,IsMapVal[int](&v),
        "IsMapVal returned false positive.",t,
    )
    v2:=reflect.ValueOf(v)
    test.BasicTest(false,IsMapVal[int](v2),
        "IsMapVal returned false positive.",t,
    )
    v2=reflect.ValueOf(&v)
    test.BasicTest(false,IsMapVal[int](v2),
        "IsMapVal returned false positive.",t,
    )
    a:=map[int]int{}
    test.BasicTest(true,IsMapVal[map[int]int](&a),
        "IsMapVal returned false negative.",t,
    )
    a2:=reflect.ValueOf(a)
    test.BasicTest(true,IsMapVal[map[int]int](a2),
        "IsMapVal returned false negative.",t,
    )
    a2=reflect.ValueOf(&a)
    test.BasicTest(true,IsMapVal[map[int]int](a2),
        "IsMapVal returned false negative.",t,
    )
}

func TestNonMapElemKeys(t *testing.T){
    v:=0
    vals,err:=MapElemKeys[int](&v).Collect()
    test.BasicTest(0,len(vals),
        "MapElemKeys returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElemKeys returned an incorrect error.",t,
        )
    }
}

func TestNonMapElemKeysReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=MapElemKeys[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "MapElemKeys returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElemKeys returned an incorrect error.",t,
        )
    }
}

func TestNonMapElemKeysReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=MapElemKeys[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "MapElemKeys returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElemKeys returned an incorrect error.",t,
        )
    }
}

func TestMapElemKeys(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    vals,err:=MapElemKeys[map[int]string](&v).Collect()
    test.BasicTest(3,len(vals),
        "MapElemKeys the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElemKeys returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,ok:=v[iterV.(int)]
        test.BasicTest(true,ok,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestMapElemKeysReflectVal(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(v)
    vals,err:=MapElemKeys[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "MapElemKeys the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElemKeys returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,ok:=v[iterV.(int)]
        test.BasicTest(true,ok,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestMapElemKeysReflectValPntr(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(&v)
    vals,err:=MapElemKeys[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "MapElemKeys the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElemKeys returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,ok:=v[iterV.(int)]
        test.BasicTest(true,ok,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestNonMapElemVals(t *testing.T){
    v:=0
    vals,err:=MapElemVals[int](&v).Collect()
    test.BasicTest(0,len(vals),
        "MapElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonMapElemValsReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=MapElemVals[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "MapElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElemVals returned an incorrect error.",t,
        )
    }
}

func TestNonMapElemValsReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=MapElemVals[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "MapElemVals returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElemVals returned an incorrect error.",t,
        )
    }
}

func TestMapElemVals(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    vals,err:=MapElemVals[map[int]string](&v).Collect()
    test.BasicTest(3,len(vals),
        "MapElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElemVals returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,_,found:=iter.MapVals(v).Find(func(val string) (bool, error) {
            return val==iterV.(string),nil
        })
        test.BasicTest(true,found,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestMapElemValsReflectVal(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(v)
    vals,err:=MapElemVals[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "MapElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElemVals returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,_,found:=iter.MapVals(v).Find(func(val string) (bool, error) {
            return val==iterV.(string),nil
        })
        test.BasicTest(true,found,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestMapElemValsReflectValPntr(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(&v)
    vals,err:=MapElemVals[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "MapElemVals the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElemVals returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,_,found:=iter.MapVals(v).Find(func(val string) (bool, error) {
            return val==iterV.(string),nil
        })
        test.BasicTest(true,found,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestNonMapElems(t *testing.T){
    v:=0
    vals,err:=MapElems[int](&v).Collect()
    test.BasicTest(0,len(vals),
        "MapElems returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapElemsReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    vals,err:=MapElemVals[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "MapElems returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapElemsReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    vals,err:=MapElemVals[reflect.Value](v2).Collect()
    test.BasicTest(0,len(vals),
        "MapElems returned values when it should not have.",t,
    )
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestMapElems(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    vals,err:=MapElems[map[int]string](&v).Collect()
    test.BasicTest(3,len(vals),
        "MapElems the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,ok:=v[iterV.A.(int)]
        test.BasicTest(true,ok,
            "A value in the array was incorrect.",t,
        )
    }
    for _,iterV:=range(vals) {
        _,_,found:=iter.MapVals(v).Find(func(val string) (bool, error) {
            return val==iterV.B,nil
        })
        test.BasicTest(true,found,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestMapElemsReflectVal(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(v)
    vals,err:=MapElems[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "MapElems the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,ok:=v[iterV.A.(int)]
        test.BasicTest(true,ok,
            "A value in the array was incorrect.",t,
        )
    }
    for _,iterV:=range(vals) {
        _,_,found:=iter.MapVals(v).Find(func(val string) (bool, error) {
            return val==iterV.B,nil
        })
        test.BasicTest(true,found,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestMapElemsReflectValPntr(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(&v)
    vals,err:=MapElems[reflect.Value](v2).Collect()
    test.BasicTest(3,len(vals),
        "MapElems the incorrect number of values.",t,
    )
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    for _,iterV:=range(vals) {
        _,ok:=v[iterV.A.(int)]
        test.BasicTest(true,ok,
            "A value in the array was incorrect.",t,
        )
    }
    for _,iterV:=range(vals) {
        _,_,found:=iter.MapVals(v).Find(func(val string) (bool, error) {
            return val==iterV.B,nil
        })
        test.BasicTest(true,found,
            "A value in the array was incorrect.",t,
        )
    }
}

func TestNonMapKeyType(t *testing.T){
    v:=0
    _,err:=MapKeyType[int](&v)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapKeyTypeReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=MapKeyType[int](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapKeyTypeReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=MapKeyType[int](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestMapKeyType(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    _t,err:=MapKeyType[map[int]string](&v)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.TypeOf(int(0)),_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestMapKeyTypeReflectVal(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(v)
    _t,err:=MapKeyType[map[int]string](v2)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.TypeOf(int(0)),_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestMapKeyTypeReflectValPntr(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(&v)
    _t,err:=MapKeyType[map[int]string](v2)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.TypeOf(int(0)),_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestNonMapKeyKind(t *testing.T){
    v:=0
    _,err:=MapKeyKind[int](&v)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapKeyKindReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=MapKeyKind[int](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapKeyKindReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=MapKeyKind[int](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestMapKeyKind(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    _t,err:=MapKeyKind[map[int]string](&v)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.Int,_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestMapKeyKindReflectVal(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(v)
    _t,err:=MapKeyKind[map[int]string](v2)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.Int,_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestMapKeyKindReflectValPntr(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(&v)
    _t,err:=MapKeyKind[map[int]string](v2)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.Int,_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestNonMapValType(t *testing.T){
    v:=0
    _,err:=MapValType[int](&v)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapValTypeReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=MapValType[int](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapValTypeReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=MapValType[int](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestMapValType(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    _t,err:=MapValType[map[int]string](&v)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.TypeOf(""),_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestMapValTypeReflectVal(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(v)
    _t,err:=MapValType[map[int]string](v2)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.TypeOf(""),_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestMapValTypeReflectValPntr(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(&v)
    _t,err:=MapValType[map[int]string](v2)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.TypeOf(""),_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestNonMapValKind(t *testing.T){
    v:=0
    _,err:=MapValKind[int](&v)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapValKindReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=MapValKind[int](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapValKindReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=MapValKind[int](v2)
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestMapValKind(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    _t,err:=MapValKind[map[int]string](&v)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.String,_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestMapValKindReflectVal(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(v)
    _t,err:=MapValKind[map[int]string](v2)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.String,_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestMapValKindReflectValPntr(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(&v)
    _t,err:=MapValKind[map[int]string](v2)
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(reflect.String,_t,
        "MapKeyType did not return the correct type.",t,
    )
}

func TestNonMapElemInfo(t *testing.T){
    v:=0
    _,err:=MapElemInfo[int](&v,true).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapElemInfoReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    _,err:=MapElemInfo[int](v2,true).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func TestNonMapElemInfoReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    _,err:=MapElemInfo[int](v2,true).Collect()
    if !IsIncorrectType(err) {
        test.FormatError(IncorrectType(""),err,
            "MapElems returned an incorrect error.",t,
        )
    }
}

func mapElemInfoHelper(v map[int]string, info []KeyValInfo, err error, t *testing.T){
    test.BasicTest(nil,err,
        "MapElems returned an error when it should not have.",t,    
    )
    test.BasicTest(len(v),len(info),
        "The wrong number of elements were returned.",t,
    )
    for _,iterV:=range(info) {
        test.BasicTest(reflect.TypeOf(int(0)),iterV.A.Type,
            "The type of the key was incorrect.",t,
        )
        test.BasicTest(reflect.TypeOf(""),iterV.B.Type,
            "The type of the key was incorrect.",t,
        )
        test.BasicTest(reflect.Int,iterV.A.Kind,
            "The kind of the key was incorrect.",t,
        )
        test.BasicTest(reflect.String,iterV.B.Kind,
            "The kind of the key was incorrect.",t,
        )
        k,ok:=iterV.A.Val()
        test.BasicTest(true,ok,
            "The key was not returned when it should have been.",t,    
        )
        actV,ok:=v[k.(int)]
        test.BasicTest(true,ok,
            "The key that was returned was not present in the map.",t,    
        )
        _v,ok:=iterV.B.Val()
        test.BasicTest(true,ok,
            "The key was not returned when it should have been.",t,    
        )
        test.BasicTest(actV,_v.(string),
            "The value that was returned was not the correct one.",t,
        )
    }
}

func TestMapElemInfo(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    info,err:=MapElemInfo[map[int]string](&v,true).Collect()
    mapElemInfoHelper(v,info,err,t)
}

func TestMapElemInfoReflectVal(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(v)
    info,err:=MapElemInfo[map[int]string](v2,true).Collect()
    mapElemInfoHelper(v,info,err,t)
}

func TestMapElemInfoReflectValPntr(t *testing.T){
    v:=map[int]string{0: "zero", 1: "one", 2: "two"}
    v2:=reflect.ValueOf(&v)
    info,err:=MapElemInfo[map[int]string](v2,true).Collect()
    mapElemInfoHelper(v,info,err,t)
}


