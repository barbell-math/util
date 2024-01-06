package circularImportTests;

import (
    "testing"
    "github.com/barbell-math/util/dataStruct"
    staticType "github.com/barbell-math/util/dataStruct/types/static"
    "github.com/barbell-math/util/algo/iter"
    "github.com/barbell-math/util/test"
)


func TestJoinEmptyLeftAndRight(t *testing.T){
    cnt,err:=iter.Join[int,int](iter.SliceElems([]int{}),iter.SliceElems([]int{}),
        func() staticType.Variant[int, int] { return dataStruct.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).Count();
    test.BasicTest(0,cnt,
        "Join on two empty iterators returned the wrong count.",t,
    );
    test.BasicTest(nil,err,
        "Join on two empty iterators returned an error when it shouldn't have.",t,
    );
}

func TestJoinEmptyLeftAndNonEmptyRight(t *testing.T){
    cntr:=0;
    err:=iter.JoinSame[int](iter.SliceElems([]int{}),
        iter.SliceElems([]int{1,2,3,4}),
        func() staticType.Variant[int, int] { return dataStruct.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        cntr++;
        test.BasicTest(index+1,val,
            "Join did not return correct values.",t,
        );
        return iter.Continue,nil;
    });
    test.BasicTest(4,cntr,
        "Join on two empty iterators returned the wrong count.",t,
    );
    test.BasicTest(nil,err,
        "Join on two empty iterators returned an error when it shouldn't have.",t,
    );
}

func TestJoinEmptyRightAndNonEmptyLeft(t *testing.T){
    cntr:=0;
    err:=iter.JoinSame[int](iter.SliceElems([]int{1,2,3,4}),
        iter.SliceElems([]int{}),
        func() staticType.Variant[int, int] { return dataStruct.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        cntr++;
        test.BasicTest(index+1,val,
            "Join did not return correct values.",t,
        );
        return iter.Continue,nil;
    });
    test.BasicTest(4,cntr,
        "Join on two empty iterators returned the wrong count.",t,
    );
    test.BasicTest(nil,err,
        "Join on two empty iterators returned an error when it shouldn't have.",t,
    );
}

func TestJoinRightLessThanLeft(t *testing.T){
    cntr:=0;
    err:=iter.JoinSame[int](iter.SliceElems([]int{1,3,5,7}),
        iter.SliceElems([]int{2,4,6}),
        func() staticType.Variant[int, int] { return dataStruct.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        cntr++;
        test.BasicTest(index+1,val,
            "Join did not return correct values.",t,
        );
        return iter.Continue,nil;
    });
    test.BasicTest(7,cntr,
        "Join on two empty iterators returned the wrong count.",t,
    );
    test.BasicTest(nil,err,
        "Join on two empty iterators returned an error when it shouldn't have.",t,
    );
}

func TestJoinLeftLessThanRight(t *testing.T){
    cntr:=0;
    err:=iter.JoinSame[int](iter.SliceElems([]int{2,4,6}),
        iter.SliceElems([]int{1,3,5,7}),
        func() staticType.Variant[int, int] { return dataStruct.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        cntr++;
        test.BasicTest(index+1,val,
            "Join did not return correct values.",t,
        );
        return iter.Continue,nil;
    });
    test.BasicTest(7,cntr,
        "Join on two empty iterators returned the wrong count.",t,
    );
    test.BasicTest(nil,err,
        "Join on two empty iterators returned an error when it shouldn't have.",t,
    );
}

func TestJoinLeftEqualsRight(t *testing.T){
    cntr:=0;
    err:=iter.JoinSame[int](iter.SliceElems([]int{2,4,6}),
        iter.SliceElems([]int{1,3,5}),
        func() staticType.Variant[int, int] { return dataStruct.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        cntr++;
        test.BasicTest(index+1,val,
            "Join did not return correct values.",t,
        );
        return iter.Continue,nil;
    });
    test.BasicTest(6,cntr,
        "Join on two empty iterators returned the wrong count.",t,
    );
    test.BasicTest(nil,err,
        "Join on two empty iterators returned an error when it shouldn't have.",t,
    );
}
