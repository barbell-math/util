package iter

import (
	"testing"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/test"
)

// All of these tests could technically go in the producer test file but there
// has been difficulty getting the channels to work so these tests are kept
// separate.

func TestMapElemsStopEarly(t *testing.T){
    f:=func() basic.Pair[string,int] { return basic.Pair[string,int]{} }
    caseHit:=false
    for i:=0; i<10; i++ {
        cntr:=0
        _,_,found:=MapElems(map[string]int{"test": 1, "test2": 2, "test3": 3},f).Find(
        func(val basic.Pair[string, int]) (bool, error) {
            cntr+=1
            return val.GetA()=="test2",nil
        })
        test.BasicTest(true,found,"ation did not find the value.",t)
        caseHit=(caseHit || cntr<3)
    }
    test.BasicTest(true,caseHit,
        "The case being tested was not hit. Run tests again.",t,
    )
}

func mapElemsHelper[K comparable, V any](m map[K]V, t *testing.T){
    f:=func() basic.Pair[K, V] { return basic.Pair[K,V]{} }
    mIter:=MapElems[K,V](m,f);
    for i:=0; i<len(m); i++ {
        mV,mErr,mBool:=mIter(Continue);
        test.BasicTest(mV.GetB(),m[mV.GetA()],
            "An incorrect pair was returned while ating over the map.",t,
        )
        test.BasicTest(nil,mErr,
            "MapElems ation produced an error when it shouldn't have.",t,
        )
        test.BasicTest(true,mBool,
            "MapElems ation stoped when it should not have.",t,
        )
    }
    _,mErr,mBool:=mIter(Continue)
    test.BasicTest(nil,mErr,
        "MapElems ation produced an error when it shouldn't have.",t,
    )
    test.BasicTest(false,mBool,
        "MapElems ations did not stop when it should have.",t,
    )
}
func TestMapElems(t *testing.T){
    mapElemsHelper(map[string]int{},t);
    mapElemsHelper(map[string]int{"test": 1},t);
    mapElemsHelper(map[string]int{"test": 1, "test2": 2, "test3": 3},t);
    mapElemsHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0},t);
}

func TestMapElemsConsume(t *testing.T){
    f:=func() basic.Pair[string,int] { return basic.Pair[string,int]{} }
    test.BasicTest(nil,MapElems[string,int](map[string]int{},f).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapElems[string,int](map[string]int{"test": 1},f).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapElems[string,int](map[string]int{"test": 1, "test2": 2},f).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapElems[string,int](map[string]int{"test": 1, "test2": 2, "test3": 3},f).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
}

func TestMapKeysStopEarly(t *testing.T){
    caseHit:=false
    for i:=0; i<10; i++ {
        cntr:=0
        _,_,found:=MapKeys(map[string]int{"test": 1, "test2": 2, "test3": 3}).Find(
        func(val string) (bool, error) {
            cntr+=1
            return val=="test2",nil
        })
        test.BasicTest(true,found,"ation did not find the value.",t)
        caseHit=(caseHit || cntr<3)
    }
    test.BasicTest(true,caseHit,
        "The case being tested was not hit. Run tests again.",t,
    )
}

func mapKeysHelper[K comparable, V any](m map[K]V, t *testing.T){
    mIter:=MapKeys[K,V](m);
    for i:=0; i<len(m); i++ {
        mV,mErr,mBool:=mIter(Continue);
        _,ok:=m[mV]
        test.BasicTest(true,ok,
            "An incorrect key was returned while ating over the map.",t,
        )
        test.BasicTest(nil,mErr,
            "MapElems ation produced an error when it shouldn't have.",t,
        )
        test.BasicTest(true,mBool,
            "MapElems ation stoped when it should not have.",t,
        )
    }
    _,mErr,mBool:=mIter(Continue)
    test.BasicTest(nil,mErr,
        "MapElems ation produced an error when it shouldn't have.",t,
    )
    test.BasicTest(false,mBool,
        "MapElems ations did not stop when it should have.",t,
    )
}
func TestMapKeys(t *testing.T){
    mapKeysHelper(map[string]int{},t);
    mapKeysHelper(map[string]int{"test": 1},t);
    mapKeysHelper(map[string]int{"test": 1, "test2": 2, "test3": 3},t);
    mapKeysHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0},t);
}

func TestMapKeysConsume(t *testing.T){
    test.BasicTest(nil,MapKeys[string,int](map[string]int{}).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapKeys[string,int](map[string]int{"test": 1}).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapKeys[string,int](map[string]int{"test": 1, "test2": 2}).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapKeys[string,int](map[string]int{"test": 1, "test2": 2, "test3": 3}).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
}

func TestMapValsStopEarly(t *testing.T){
    caseHit:=false
    for i:=0; i<10; i++ {
        cntr:=0
        _,_,found:=MapVals(map[string]int{"test": 1, "test2": 2, "test3": 3}).Find(
        func(val int) (bool, error) {
            cntr+=1
            return val==2,nil
        })
        test.BasicTest(true,found,"ation did not find the value.",t)
        caseHit=(caseHit || cntr<3)
    }
    test.BasicTest(true,caseHit,
        "The case being tested was not hit. Run tests again.",t,
    )
}

func mapValsHelper[K comparable, V comparable](m map[K]V, t *testing.T){
    mIter:=MapVals[K,V](m);
    for i:=0; i<len(m); i++ {
        mV,mErr,mBool:=mIter(Continue);
        found:=false;
        for _,v:=range(m) {
            found=(found || v==mV)
        }
        test.BasicTest(true,found,
            "An incorrect value was returned while ating over the map.",t,
        )
        test.BasicTest(nil,mErr,
            "MapElems ation produced an error when it shouldn't have.",t,
        )
        test.BasicTest(true,mBool,
            "MapElems ation stoped when it should not have.",t,
        )
    }
    _,mErr,mBool:=mIter(Continue)
    test.BasicTest(nil,mErr,
        "MapElems ation produced an error when it shouldn't have.",t,
    )
    test.BasicTest(false,mBool,
        "MapElems ations did not stop when it should have.",t,
    )
}
func TestMapVals(t *testing.T){
    mapValsHelper(map[string]int{},t);
    mapValsHelper(map[string]int{"test": 1},t);
    mapValsHelper(map[string]int{"test": 1, "test2": 2, "test3": 3},t);
    mapValsHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0},t);
}

func TestMapValsConsume(t *testing.T){
    test.BasicTest(nil,MapVals[string,int](map[string]int{}).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapVals[string,int](map[string]int{"test": 1}).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapVals[string,int](map[string]int{"test": 1, "test2": 2}).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
    test.BasicTest(
        nil,
        MapVals[string,int](map[string]int{"test": 1, "test2": 2, "test3": 3}).Consume(),
        "Consuming the val iterator returned an error when it should not have.",t,
    )
}

func TestMapElemsChanClosing(t *testing.T) {
    test.NoPanic(
        func() {
            m:=map[string]int{"test": 1, "test2": 2, "test3": 3}
            for i:=0; i<10000; i++ {
                MapVals[string,int](m).Consume()
                m["four"]=i
            }
        },
        "Writing to the map after iteration finished paniced. (Means channel did not immediately close.)",t,
    )
}
