package circularImportTests

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/dataStruct"
	"github.com/barbell-math/util/test"
)


func collectIntoIterHelper[T any](
        vals []T,
        expectedNumChange int,
        t *testing.T){
    buff,_:=dataStruct.NewVector[T](0)
    rv,err:=iter.SliceElems(vals).CollectInto(&buff,dataStruct.AppendCollector[int,T])
    test.BasicTest(len(vals),len(buff),
        "Buffer has the wrong number of elements.",t,
    );
    test.BasicTest(expectedNumChange,rv,
        "Total number of elements changed is not correct",t,
    );
    test.BasicTest(nil,err,
        "CollectInto returned and error when it was not supposed to.",t,
    );
    min:=len(buff);
    if len(vals)<len(buff) {
        min=len(vals);
    }
    for i:=0; i<min; i++ {
        test.BasicTest(vals[i],buff[i],fmt.Sprintf(
            "Values do not match | Index: %d",i,
        ),t);
    }
}

func TestCollectInto(t *testing.T){
    collectIntoIterHelper([]int{1,2,3,4},4,t);
    collectIntoIterHelper([]int{1},1,t);
    collectIntoIterHelper([]int{},0,t);
}
