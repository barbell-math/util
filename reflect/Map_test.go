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
