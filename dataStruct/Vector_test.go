package dataStruct

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/dataStruct/types/static"
	customerr "github.com/barbell-math/util/err"
	"github.com/barbell-math/util/test"
)

func TestVectorWriteInterface(t *testing.T){
    v:=Vector[string](make(Vector[string],0))
    v2,_:=NewSyncedVector[string](5)
    writeInterfaceTypeCeck[int,string](&v);
    writeInterfaceTypeCeck[int,string](&v2);
}

func TestVectorReadInterface(t *testing.T){
    v:=Vector[string](make(Vector[string],0))
    v2,_:=NewSyncedVector[string](5)
    readInterfaceTypeCeck[int,string](&v);
    readInterfaceTypeCeck[int,string](&v2);
}

func TestVectorDynVectorTypeInterface(t *testing.T) {
    v:=Vector[int](make(Vector[int], 0))
    v2,_:=NewSyncedVector[int](5)
    dynVectorInterfaceTypeCheck[int](&v);
    dynVectorInterfaceTypeCheck[int](&v2);
}

func TestVectorDynStackTypeInterface(t *testing.T) {
    v:=Vector[int](make(Vector[int], 0))
    v2,_:=NewSyncedVector[int](5)
    dynStackInterfaceTypeCheck[int](&v);
    dynStackInterfaceTypeCheck[int](&v2);
}

func TestVectorDynQueueTypeInterface(t *testing.T) {
    v:=Vector[int](make(Vector[int], 0))
    v2,_:=NewSyncedVector[int](5)
    dynQueueInterfaceTypeCheck[int](&v);
    dynQueueInterfaceTypeCheck[int](&v2);
}

func TestVectorDynDequeTypeInterface(t *testing.T) {
    v:=Vector[int](make(Vector[int], 0))
    v2,_:=NewSyncedVector[int](5)
    dynDequeInterfaceTypeCheck[int](&v);
    dynDequeInterfaceTypeCheck[int](&v2);
}

func TestVectorEqualsTypeInterface(t *testing.T) {
    v:=Vector[int](make(Vector[int], 0))
    v2,_:=NewSyncedVector[int](5)
    equalsInterfaceTypeCheck[Vector[int]](&v);
    equalsInterfaceTypeCheck[Vector[int]](&v2);
}

func TestVectorStaticTypeInterface(t *testing.T){
    test.Panics(
        func () {
            var c any
            c,_=NewVector[int](5)
            c2:=c.(static.Vector[int])
            _=c2
        }, 
        "Code did not panic when casting a dynamic vector to a static vector.",t,
    )
    test.Panics(
        func () {
            var c any
            c,_=NewVector[int](5)
            c2:=c.(static.Queue[int])
            _=c2
        }, 
        "Code did not panic when casting a dynamic vector to a static queue.",t,
    )
    test.Panics(
        func () {
            var c any
            c,_=NewVector[int](5)
            c2:=c.(static.Stack[int])
            _=c2
        }, 
        "Code did not panic when casting a dynamic vector to a static stack.",t,
    )
    test.Panics(
        func () {
            var c any
            c,_=NewVector[int](5)
            c2:=c.(static.Deque[int])
            _=c2
        }, 
        "Code did not panic when casting a dynamic vector to a static deque.",t,
    )
}

func TestVectorGet(t *testing.T){
    v:=Vector[int]{0,1,2,3,4,5}
    for i:=0; i<5; i++ {
	_v,err:=v.Get(i)
	test.BasicTest(i,_v,
	    "Get did not return the correct value.",t,
	)
	test.BasicTest(nil,err,
	    "Get returned an error when it shouldn't have.",t,
	)
    }
    _,err:=v.Get(-1)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
    _,err=v.Get(6)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
    v=Vector[int]{}
    _,err=v.Get(0)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
}

func TestVectorGetPntr(t *testing.T){
    v:=Vector[int]{0,1,2,3,4,5}
    for i:=0; i<5; i++ {
	_v,err:=v.GetPntr(i)
	test.BasicTest(i,*_v,
	    "Get did not return the correct value.",t,
	)
	test.BasicTest(nil,err,
	    "Get returned an error when it shouldn't have.",t,
	)
    }
    _,err:=v.GetPntr(-1)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
    _,err=v.GetPntr(6)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
    v=Vector[int]{}
    _,err=v.GetPntr(0)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
}

func TestVectorSet(t *testing.T){
    v:=Vector[int]{0,1,2,3,4,5}
    for i:=0; i<5; i++ {
	err:=v.Set(i,i+1)
	test.BasicTest(nil,err,
	    "Get returned an error when it shouldn't have.",t,
	)
    }
    for i:=0; i<5; i++ {
	test.BasicTest(i+1,v[i],
	    "Set did not set the value correctly.",t,
	)
    }
    err:=v.Set(-1,6)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
    err=v.Set(6,6)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
    v=Vector[int]{}
    err=v.Set(0,6)
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Get did not return the correct error with invalid index.",t,
	)
    }
}

func TestVectorAppend(t *testing.T) {
    v,_:=NewVector[int](0)
    for i:=0; i<5; i++ { v.Append(i) }
    for i:=0; i<5; i++ {
	test.BasicTest(i,v[i],
	    "Append did not add the value correctly.",t,
	)
    }
    v.Append(5,6,7)
    for i:=0; i<8; i++ {
	test.BasicTest(i,v[i],
	    "Append did not add the value correctly.",t,
	)
    }
}

func vectorInsertHelper(idx int, l int, t *testing.T){
    tmp:=make(Vector[int],l-1)
    for i:=0; i<l-1; i++ { tmp[i]=i }
    err:=tmp.Insert(idx,l-1)
    test.BasicTest(nil,err,
        "Insert returned an error when it shouldn't have.",t,
    )
    test.BasicTest(l,tmp.Length(),
        "Insert did not increment the number of elements.",t,
    )
    for i:=0; i<len(tmp); i++ {
        var exp int
        v,_:=tmp.Get(i)
        if i<idx {
            exp=i
        } else if i==idx {
            exp=l-1
        } else {
            exp=i-1
        }
        test.BasicTest(exp,v,
            "Insert did not put the value in the correct place.",t,
        )
    }
}
func TestVectorInsert(t *testing.T){
    v:=make(Vector[int],0)
    for i:=2; i>=0; i-- { v.Insert(0,i) }
    for i:=3; i<5; i++ { v.Insert(len(v),i) }
    for i:=0; i<5; i++ {
	test.BasicTest(i,v[i],
	    "Insert did not put the values in the correct place.",t,
	)
    }
    for i:=0; i<5; i++ {
	vectorInsertHelper(i,5,t)
    }
}

func TestVectorDelete(t *testing.T){
    v:=Vector[int]{0,1,2,3,4,5}
    for i:=len(v)-1; i>=0; i-- {
	v.Delete(i)
	test.BasicTest(i,len(v),"Delete removed to many values.",t)
	for j:=0; j<i; j++ {
	    test.BasicTest(j,v[j],"Delete changed the wrong value.",t)
	}
    }
    err:=v.Delete(0)
    test.BasicTest(nil,err,
	"Delete returned an error when it should not have.",t,
    )
}

func TestVectorClear(t *testing.T){
    v:=Vector[int]{0,1,2,3,4,5}
    v.Clear()
    test.BasicTest(0,len(v),"Clear did not reset the underlying vector.",t)
    test.BasicTest(0,cap(v),"Clear did not reset the underlying vector.",t)
}

func TestVectorSetCapacity(t *testing.T){
    v:=Vector[int]{0,1,2,3,4,5}
    test.BasicTest(6,len(v),"Initial length is not correct.",t)
    test.BasicTest(6,cap(v),"Initial cap is not correct.",t)
    v.SetCapacity(10)
    test.BasicTest(10,len(v),"Larger length is not correct.",t)
    test.BasicTest(10,cap(v),"Larger cap is not correct.",t)
    for i:=0; i<len(v); i++ {
	exp:=0
	if i<6 {
	    exp=i
	}
	test.BasicTest(exp,v[i],"Changing capacity changed the values.",t)
    }
    v.SetCapacity(3)
    test.BasicTest(3,len(v),"Smaller length is not correct.",t)
    test.BasicTest(3,cap(v),"Smaller cap is not correct.",t)
    for i:=0; i<len(v); i++ {
	test.BasicTest(i,v[i],"Changing capacity changed the values.",t)
    }
}

func TestVectorPeekPntrFront(t *testing.T){
    v:=Vector[int]{}
    _v,err:=v.PeekPntrFront()
    test.BasicTest((*int)(nil),_v,
	"Peek pntr front did not return the correct value.",t,	
    )
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Peek pntr front returned an incorrect error.",t,
	)
    }
    v.Append(1)
    _v,err=v.PeekPntrFront()
    test.BasicTest(1,*_v,
	"Peek pntr front did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek pntr front returned an error when it shouldn't have.",t,
    )
    v.Append(2)
    _v,err=v.PeekPntrFront()
    test.BasicTest(1,*_v,
	"Peek pntr front did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek pntr front returned an error when it shouldn't have.",t,
    )
}

func TestVectorPeekFront(t *testing.T){
    v:=Vector[int]{}
    _,err:=v.PeekFront()
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Peek front returned an incorrect error.",t,
	)
    }
    v.Append(1)
    _v,err:=v.PeekFront()
    test.BasicTest(1,_v,
	"Peek front did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek front returned an error when it shouldn't have.",t,
    )
    v.Append(2)
    _v,err=v.PeekFront()
    test.BasicTest(1,_v,
	"Peek front did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek front returned an error when it shouldn't have.",t,
    )
}

func TestVectorPeekPntrBack(t *testing.T){
    v:=Vector[int]{}
    _v,err:=v.PeekPntrBack()
    test.BasicTest((*int)(nil),_v,
	"Peek pntr back did not return the correct value.",t,	
    )
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Peek pntr back returned an incorrect error.",t,
	)
    }
    v.Append(1)
    _v,err=v.PeekPntrBack()
    test.BasicTest(1,*_v,
	"Peek pntr back did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek pntr back returned an error when it shouldn't have.",t,
    )
    v.Append(2)
    _v,err=v.PeekPntrBack()
    test.BasicTest(2,*_v,
	"Peek pntr back did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek pntr back returned an error when it shouldn't have.",t,
    )
}

func TestVectorPeekBack(t *testing.T){
    v:=Vector[int]{}
    _,err:=v.PeekBack()
    if !customerr.IsValOutsideRange(err) {
	test.FormatError(customerr.ValOutsideRange(""),err,
	    "Peek back returned an incorrect error.",t,
	)
    }
    v.Append(1)
    _v,err:=v.PeekBack()
    test.BasicTest(1,_v,
	"Peek back did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek back returned an error when it shouldn't have.",t,
    )
    v.Append(2)
    _v,err=v.PeekBack()
    test.BasicTest(2,_v,
	"Peek back did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek back returned an error when it shouldn't have.",t,
    )
}

func TestVectorPopFront(t *testing.T) {
    v:=Vector[int]{0,1,2,3}
    for i:=0; i<4; i++ {
	f,err:=v.PopFront()
	test.BasicTest(i,f,
	    "Pop front returned the incorrect value.",t,
	)
	test.BasicTest(nil,err,
	    "Pop front returned an error when it shoudn't have.",t,
	)
    }
    _,err:=v.PopFront()
    if !IsEmpty(err) {
	test.FormatError(Empty(""),err,
	    "Pop front returned an incorrect error.",t,
	)
    }
}

func TestVectorPopBack(t *testing.T) {
    v:=Vector[int]{0,1,2,3}
    for i:=3; i>=0; i-- {
	f,err:=v.PopBack()
	test.BasicTest(i,f,
	    "Pop front returned the incorrect value.",t,
	)
	test.BasicTest(nil,err,
	    "Pop front returned an error when it shoudn't have.",t,
	)
    }
    _,err:=v.PopBack()
    if !IsEmpty(err) {
	test.FormatError(Empty(""),err,
	    "Pop front returned an incorrect error.",t,
	)
    }
}

func TestVectorPushFront(t *testing.T){
    v:=Vector[int]{}
    for i:=0; i<4; i++ {
	v.PushFront(i)
	test.BasicTest(i+1,len(v),
	    "Push front did not add the value correctly.",t,
	)
	for j:=0; j<i+1; j++ {
	    test.BasicTest(i-j,v[j],
		"Push front did not put the value in the correct place.",t,
	    )
	}
    }
}

func TestVectorPushBack(t *testing.T){
    v:=Vector[int]{}
    for i:=0; i<4; i++ {
	v.PushBack(i)
	test.BasicTest(i+1,len(v),
	    "Push back did not add the value correctly.",t,
	)
	for j:=0; j<i+1; j++ {
	    test.BasicTest(j,v[j],
		"Push back did not put the value in the correct place.",t,
	    )
	}
    }
}

func testVectorElemsHelper(v SyncedVector[int], l int, t *testing.T){
    for i:=0; i<l; i++ {
        v.PushBack(i);
    }
    cnt:=0
    v.Elems().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        cnt++
        test.BasicTest(index,val,"Element was skipped while iterating.",t);
        return iter.Continue,nil;
    });
    test.BasicTest(l,cnt,
        "All the elements were not iterated over.",t,
    )
}
func TestVectorElems(t *testing.T){
    tmp,err:=NewSyncedVector[int](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,0,t);
    tmp,err=NewSyncedVector[int](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,1,t);
    tmp,err=NewSyncedVector[int](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,2,t);
}

func testVectorPntrElemsHelper(v Vector[int], l int, t *testing.T){
    for i:=0; i<l; i++ {
        v.PushBack(i);
    }
    cnt:=0
    v.PntrElems().ForEach(func(index int, val *int) (iter.IteratorFeedback, error) {
        cnt++
        test.BasicTest(index,*val,"Element was skipped while iterating.",t);
        *val=100;
        return iter.Continue,nil;
    });
    v.Elems().ForEach(func(index int, val int) (iter.IteratorFeedback, error) {
        test.BasicTest(100,val,"Element was not updated while iterating.",t);
        return iter.Continue,nil;
    });
    test.BasicTest(l,cnt,
        "All the elements were not iterated over.",t,
    )
}
func TestVectorElemPntrs(t *testing.T){
    tmp,err:=NewSyncedVector[int](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,0,t);
    tmp,err=NewSyncedVector[int](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,1,t);
    tmp,err=NewSyncedVector[int](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,2,t);
}

func TestVectorEq(t *testing.T){
    v:=Vector[int]{0,1,2,3}
    v2:=Vector[int]{0,1,2,3}
    test.BasicTest(true,v.Eq(v2),
	"Eq returned a false negative.",t,
    )
    test.BasicTest(true,v2.Eq(v),
	"Eq returned a false negative.",t,
    )
    v.Delete(3)
    test.BasicTest(false,v.Eq(v2),
	"Eq returned a false positive.",t,
    )
    test.BasicTest(false,v2.Eq(v),
	"Eq returned a false positive.",t,
    )
    v=Vector[int]{0}
    v2=Vector[int]{0}
    test.BasicTest(true,v.Eq(v2),
	"Eq returned a false negative.",t,
    )
    test.BasicTest(true,v2.Eq(v),
	"Eq returned a false negative.",t,
    )
    v.Delete(0)
    test.BasicTest(false,v.Eq(v2),
	"Eq returned a false positive.",t,
    )
    test.BasicTest(false,v2.Eq(v),
	"Eq returned a false positive.",t,
    )
    v=Vector[int]{}
    v2=Vector[int]{}
    test.BasicTest(true,v.Eq(v2),
	"Eq returned a false negative.",t,
    )
    test.BasicTest(true,v2.Eq(v),
	"Eq returned a false negative.",t,
    )
}

func TestVectorNeq(t *testing.T){
    v:=Vector[int]{0,1,2,3}
    v2:=Vector[int]{0,1,2,3}
    test.BasicTest(false,v.Neq(v2),
	"Neq returned a false positive.",t,
    )
    test.BasicTest(false,v2.Neq(v),
	"Neq returned a false positive.",t,
    )
    v.Delete(3)
    test.BasicTest(true,v.Neq(v2),
	"Neq returned a false negative.",t,
    )
    test.BasicTest(true,v2.Neq(v),
	"Neq returned a false negative.",t,
    )
    v=Vector[int]{0}
    v2=Vector[int]{0}
    test.BasicTest(false,v.Neq(v2),
	"Neq returned a false positive.",t,
    )
    test.BasicTest(false,v2.Neq(v),
	"Neq returned a false positive.",t,
    )
    v.Delete(0)
    test.BasicTest(true,v.Neq(v2),
	"Neq returned a false negative.",t,
    )
    test.BasicTest(true,v2.Neq(v),
	"Neq returned a false negative.",t,
    )
    v=Vector[int]{}
    v2=Vector[int]{}
    test.BasicTest(false,v.Neq(v2),
	"Neq returned a false positive.",t,
    )
    test.BasicTest(false,v2.Neq(v),
	"Neq returned a false positive.",t,
    )
}
