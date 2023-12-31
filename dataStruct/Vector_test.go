package dataStruct

import (
	"testing"

	"github.com/barbell-math/util/dataStruct/types/static"
	customerr "github.com/barbell-math/util/err"
	"github.com/barbell-math/util/test"
)

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
	err:=v.Set(i+1,i)
	test.BasicTest(nil,err,
	    "Get returned an error when it shouldn't have.",t,
	)
    }
    for i:=0; i<5; i++ {
	test.BasicTest(i+1,v[i],
	    "Set did not set the value correctly.",t,
	)
    }
    err:=v.Set(6,-1)
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
    err=v.Set(6,0)
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

func TestPeekPntrFront(t *testing.T){
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

func TestPeekFront(t *testing.T){
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

func TestPeekPntrBack(t *testing.T){
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

func TestPeekBack(t *testing.T){
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

func TestPopFront(t *testing.T) {
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

func TestPopBack(t *testing.T) {
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

func TestPushFront(t *testing.T){
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

func TestPushBack(t *testing.T){
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
