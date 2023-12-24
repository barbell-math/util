package circularImportTests

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/dataStruct"
	"github.com/barbell-math/util/test"
)

func mapElemsHelper[K comparable, V any](m map[K]V, t *testing.T){
    p:=dataStruct.Pair[K,V]{}
    mIter:=iter.MapElems[K,V](m,&p);
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
