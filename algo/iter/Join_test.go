package iter

import (
	"testing"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/test"
)


func TestJoinEmptyLeftAndRight(t *testing.T){
    cnt,err:=Join[int,int](SliceElems([]int{}),SliceElems([]int{}),
        func() basic.Variant[int, int] { return basic.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).Count();
    test.Eq(0,cnt,t)
    test.Nil(err,t)
}

func TestJoinEmptyLeftAndNonEmptyRight(t *testing.T){
    cntr:=0;
    err:=JoinSame[int](SliceElems([]int{}),
        SliceElems([]int{1,2,3,4}),
        func() basic.Variant[int, int] { return basic.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (IteratorFeedback, error) {
        cntr++;
        test.Eq(index+1,val,t)
        return Continue,nil;
    });
    test.Eq(4,cntr,t)
    test.Nil(err,t)
}

func TestJoinEmptyRightAndNonEmptyLeft(t *testing.T){
    cntr:=0;
    err:=JoinSame[int](SliceElems([]int{1,2,3,4}),
        SliceElems([]int{}),
        func() basic.Variant[int, int] { return basic.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (IteratorFeedback, error) {
        cntr++;
        test.Eq(index+1,val,t)
        return Continue,nil;
    });
    test.Eq(4,cntr,t)
    test.Nil(err,t)
}

func TestJoinRightLessThanLeft(t *testing.T){
    cntr:=0;
    err:=JoinSame[int](SliceElems([]int{1,3,5,7}),
        SliceElems([]int{2,4,6}),
        func() basic.Variant[int, int] { return basic.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (IteratorFeedback, error) {
        cntr++;
        test.Eq(index+1,val,t)
        return Continue,nil;
    });
    test.Eq(7,cntr,t)
    test.Nil(err,t)
}

func TestJoinLeftLessThanRight(t *testing.T){
    cntr:=0;
    err:=JoinSame[int](SliceElems([]int{2,4,6}),
        SliceElems([]int{1,3,5,7}),
        func() basic.Variant[int, int] { return basic.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (IteratorFeedback, error) {
        cntr++;
        test.Eq(index+1,val,t)
        return Continue,nil;
    });
    test.Eq(7,cntr,t)
    test.Nil(err,t)
}

func TestJoinLeftEqualsRight(t *testing.T){
    cntr:=0;
    err:=JoinSame[int](SliceElems([]int{2,4,6}),
        SliceElems([]int{1,3,5}),
        func() basic.Variant[int, int] { return basic.Variant[int,int]{} },
        func(left, right int) bool { return left<right; },
    ).ForEach(func(index, val int) (IteratorFeedback, error) {
        cntr++;
        test.Eq(index+1,val,t)
        return Continue,nil;
    });
    test.Eq(6,cntr,t)
    test.Nil(err,t)
}
