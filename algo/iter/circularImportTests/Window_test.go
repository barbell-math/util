package circularImportTests;

import (
    "testing"
	"github.com/barbell-math/util/dataStruct/types"
	"github.com/barbell-math/util/dataStruct"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/test"
)


func TestWindowEmpty(t *testing.T){
    q,_:=dataStruct.NewCircularQueue[int](2)
    cnt,err:=iter.Window[int](iter.SliceElems([]int{}),&q,true).Count();
    test.BasicTest(0,cnt,"Window returned values from empty parent iterator.",t);
    test.BasicTest(nil,err,
        "Window returned an error when it was not supposed to.",t,
    );
}

func TestWindowNoPartialsNoWindowValues(t *testing.T){
    cntr:=0;
    q,_:=dataStruct.NewCircularQueue[int](101);
    vals:=make([]int,100);
    for i:=0; i<len(vals); i++ {
        vals[i]=i;
    }
    cntr,err:=iter.Window[int](iter.SliceElems(vals),&q,false).Count();
    test.BasicTest(0,cntr,
        "Window returned values from empty parent iterator.",t,
    );
    test.BasicTest(nil,err,
        "Window returned an error when it was not supposed to.",t,
    );
}

func TestWindowNoPartials(t *testing.T){
    cntr:=0;
    q,_:=dataStruct.NewCircularQueue[int](2);
    vals:=make([]int,100);
    for i:=0; i<len(vals); i++ {
        vals[i]=i;
    }
    err:=iter.Window[int](iter.SliceElems(vals),&q,false).ForEach(
    func(index int, val types.Queue[int]) (iter.IteratorFeedback, error) {
        cntr++;
        test.BasicTest(2,q.Length(),
            "Partials were returned when they should not have.",t,
        );
        if v,err:=q.Peek(0); err==nil {
            test.BasicTest(index,v,"Window values were out of order.",t);
        }
        if v,err:=q.Peek(1); err==nil {
            test.BasicTest(index+1,v,"Window values were out of order.",t);
        }
        return iter.Continue,nil;
    });
    test.BasicTest(99,cntr,
        "Window returned wrong number of values.",t,
    );
    test.BasicTest(nil,err,
        "Window returned an error when it was not supposed to.",t,
    );
}

func TestWindowPartials(t *testing.T){
    cntr:=0;
    q,_:=dataStruct.NewCircularQueue[int](2);
    vals:=make([]int,100);
    for i:=0; i<len(vals); i++ {
        vals[i]=i;
    }
    err:=iter.Window[int](iter.SliceElems(vals),&q,true).ForEach(
    func(index int, val types.Queue[int]) (iter.IteratorFeedback, error) {
        cntr++;
        if index==0 || index==100 {
            test.BasicTest(1,q.Length(),
                "Partials were returned when they should not have.",t,
            );
            if v,err:=q.Peek(0); err==nil {
                test.BasicTest(index,v,"Window values were out of order.",t);
            }
        } else {
            test.BasicTest(2,q.Length(),
                "Partials were returned when they should not have.",t,
            );
            if v,err:=q.Peek(0); err==nil {
                test.BasicTest(index-1,v,"Window values were out of order.",t);
            }
            if v,err:=q.Peek(1); err==nil {
                test.BasicTest(index,v,"Window values were out of order.",t);
            }
        }
        return iter.Continue,nil;
    });
    test.BasicTest(100,cntr,
        "Window returned wrong number of values.",t,
    );
    test.BasicTest(nil,err,
        "Window returned an error when it was not supposed to.",t,
    );
}
