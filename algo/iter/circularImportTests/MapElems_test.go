package circularImportTests

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/dataStruct"
	"github.com/barbell-math/util/dataStruct/types/static"
	"github.com/barbell-math/util/test"
)

func TestMapElemsStopEarly(t *testing.T){
    f:=func() static.Pair[string,int] { return &dataStruct.Pair[string,int]{} }
    caseHit:=false
    for i:=0; i<10; i++ {
        cntr:=0
        _,_,found:=iter.MapElems(map[string]int{"test": 1, "test2": 2, "test3": 3},f).Find(
        func(val static.Pair[string, int]) (bool, error) {
            cntr+=1
            return val.GetA()=="test2",nil
        })
        test.BasicTest(true,found,"Iteration did not find the value.",t)
        caseHit=(caseHit || cntr<3)
    }
    test.BasicTest(true,caseHit,
        "The case being tested was not hit. Run tests again.",t,
    )
}

func mapElemsHelper[K comparable, V any](m map[K]V, t *testing.T){
    f:=func() static.Pair[K, V] { return &dataStruct.Pair[K,V]{} }
    mIter:=iter.MapElems[K,V](m,f);
    for i:=0; i<len(m); i++ {
        mV,mErr,mBool:=mIter(iter.Continue);
        test.BasicTest(mV.GetB(),m[mV.GetA()],
            "An incorrect pair was returned while iterating over the map.",t,
        )
        test.BasicTest(nil,mErr,
            "MapElems iteration produced an error when it shouldn't have.",t,
        )
        test.BasicTest(true,mBool,
            "MapElems iteration stoped when it should not have.",t,
        )
    }
    _,mErr,mBool:=mIter(iter.Continue)
    test.BasicTest(nil,mErr,
        "MapElems iteration produced an error when it shouldn't have.",t,
    )
    test.BasicTest(false,mBool,
        "MapElems iterations did not stop when it should have.",t,
    )
}
func TestMapElems(t *testing.T){
    mapElemsHelper(map[string]int{},t);
    mapElemsHelper(map[string]int{"test": 1},t);
    mapElemsHelper(map[string]int{"test": 1, "test2": 2, "test3": 3},t);
    mapElemsHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0},t);
}

func TestMapKeysStopEarly(t *testing.T){
    caseHit:=false
    for i:=0; i<10; i++ {
        cntr:=0
        _,_,found:=iter.MapKeys(map[string]int{"test": 1, "test2": 2, "test3": 3}).Find(
        func(val string) (bool, error) {
            cntr+=1
            return val=="test2",nil
        })
        test.BasicTest(true,found,"Iteration did not find the value.",t)
        caseHit=(caseHit || cntr<3)
    }
    test.BasicTest(true,caseHit,
        "The case being tested was not hit. Run tests again.",t,
    )
}

func mapKeysHelper[K comparable, V any](m map[K]V, t *testing.T){
    mIter:=iter.MapKeys[K,V](m);
    for i:=0; i<len(m); i++ {
        mV,mErr,mBool:=mIter(iter.Continue);
        _,ok:=m[mV]
        test.BasicTest(true,ok,
            "An incorrect key was returned while iterating over the map.",t,
        )
        test.BasicTest(nil,mErr,
            "MapElems iteration produced an error when it shouldn't have.",t,
        )
        test.BasicTest(true,mBool,
            "MapElems iteration stoped when it should not have.",t,
        )
    }
    _,mErr,mBool:=mIter(iter.Continue)
    test.BasicTest(nil,mErr,
        "MapElems iteration produced an error when it shouldn't have.",t,
    )
    test.BasicTest(false,mBool,
        "MapElems iterations did not stop when it should have.",t,
    )
}
func TestMapKeys(t *testing.T){
    mapKeysHelper(map[string]int{},t);
    mapKeysHelper(map[string]int{"test": 1},t);
    mapKeysHelper(map[string]int{"test": 1, "test2": 2, "test3": 3},t);
    mapKeysHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0},t);
}

func mapValsHelper[K comparable, V comparable](m map[K]V, t *testing.T){
    mIter:=iter.MapVals[K,V](m);
    for i:=0; i<len(m); i++ {
        mV,mErr,mBool:=mIter(iter.Continue);
        found:=false;
        for _,v:=range(m) {
            found=(found || v==mV)
        }
        test.BasicTest(true,found,
            "An incorrect value was returned while iterating over the map.",t,
        )
        test.BasicTest(nil,mErr,
            "MapElems iteration produced an error when it shouldn't have.",t,
        )
        test.BasicTest(true,mBool,
            "MapElems iteration stoped when it should not have.",t,
        )
    }
    _,mErr,mBool:=mIter(iter.Continue)
    test.BasicTest(nil,mErr,
        "MapElems iteration produced an error when it shouldn't have.",t,
    )
    test.BasicTest(false,mBool,
        "MapElems iterations did not stop when it should have.",t,
    )
}
func TestMapVals(t *testing.T){
    mapValsHelper(map[string]int{},t);
    mapValsHelper(map[string]int{"test": 1},t);
    mapValsHelper(map[string]int{"test": 1, "test2": 2, "test3": 3},t);
    mapValsHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0},t);
}

func TestMapValsStopEarly(t *testing.T){
    caseHit:=false
    for i:=0; i<10; i++ {
        cntr:=0
        _,_,found:=iter.MapVals(map[string]int{"test": 1, "test2": 2, "test3": 3}).Find(
        func(val int) (bool, error) {
            cntr+=1
            return val==2,nil
        })
        test.BasicTest(true,found,"Iteration did not find the value.",t)
        caseHit=(caseHit || cntr<3)
    }
    test.BasicTest(true,caseHit,
        "The case being tested was not hit. Run tests again.",t,
    )
}
