package dataStruct

import (
    "testing"

    "github.com/barbell-math/util/algo/iter"
    customerr "github.com/barbell-math/util/err"
    "github.com/barbell-math/util/test"
    "github.com/barbell-math/util/dataStruct/types/dynamic"
)

func TestNewCircularBuffer(t *testing.T) {
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    test.BasicTest(5,len(tmp.vals),
        "NewCircularBuffer added values to empty queue during initialization.",t,
    );
    test.BasicTest(5,cap(tmp.vals),
        "NewCircularBuffer did not set capacity correctly.",t,
    );
    test.BasicTest(0,tmp.startEnd.A,
        "NewCircularBuffer added values to empty queue during initialization.",t,
    );
    test.BasicTest(4,tmp.startEnd.B,
        "NewCircularBuffer added values to empty queue during initialization.",t,
    );
}

func TestNewCircularBufferBadSize(t *testing.T) {
    tmp,err:=NewCircularBuffer[int](0);
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Invalid queue size did not raise the correct error.",t,
        );
    }
    test.BasicTest(0,len(tmp.vals),
        "NewCircularBuffer added values to empty queue during initialization.",t,
    );
    test.BasicTest(0,cap(tmp.vals),
        "NewCircularBuffer did not set capacity correctly.",t,
    );
    test.BasicTest(0,tmp.startEnd.A,
        "NewCircularBuffer added values to empty queue during initialization.",t,
    );
    test.BasicTest(0,tmp.startEnd.B,
        "NewCircularBuffer added values to empty queue during initialization.",t,
    );
}

func TestCircularBufferPushFront(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    for i:=0; i<5; i++ {
        res:=tmp.PushFront(i);
        test.BasicTest(nil,res,
            "Push returned an error when it should not have.",t,
        );
        test.BasicTest(i+1,tmp.numElems,
            "Push did not increment NumElems after adding value.",t,
        );
        test.BasicTest(i,tmp.vals[4-i],"Push did not save value.",t);
        test.BasicTest(4-i,tmp.startEnd.A,
            "Push did not modify the start index.",t,
        );
        test.BasicTest(4,tmp.startEnd.B,"Push modified the end index.",t);
    }
    res:=tmp.PushBack(5);
    if !IsQueueFull(res) {
        test.FormatError(QueueFull(""),res,
            "Push did not detect the queue was full.",t,
        );
    }
    test.BasicTest(5,tmp.numElems,
        "Push incremented NumElems when queue was full.",t,
    );
    test.BasicTest(0,tmp.startEnd.A,"Push modified the start index.",t);
    test.BasicTest(4,tmp.startEnd.B,
        "Push modified the end index when the queue was full.",t,
    );
}

func TestCircularBufferPushBack(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    for i:=0; i<5; i++ {
        res:=tmp.PushBack(i);
        test.BasicTest(nil,res,
            "Push returned an error when it should not have.",t,
        );
        test.BasicTest(i+1,tmp.numElems,
            "Push did not increment NumElems after adding value.",t,
        );
        test.BasicTest(i,tmp.vals[i],"Push did not save value.",t);
        test.BasicTest(0,tmp.startEnd.A,"Push modified the start index.",t);
        test.BasicTest(i,tmp.startEnd.B,
            "Push did not modify the end index.",t,
        );
    }
    res:=tmp.PushBack(5);
    if !IsQueueFull(res) {
        test.FormatError(QueueFull(""),res,
            "Push did not detect the queue was full.",t,
        );
    }
    test.BasicTest(5,tmp.numElems,
        "Push incremented NumElems when queue was full.",t,
    );
    test.BasicTest(0,tmp.startEnd.A,"Push modified the start index.",t);
    test.BasicTest(len(tmp.vals)-1,tmp.startEnd.B,
        "Push modified the end index when the queue was full.",t,
    );
}

func TestCircularBufferInsertFront(t *testing.T){
    tmp,_:=NewCircularBuffer[int](5)
    err:=tmp.Insert(1,1)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Insert returned the wrong error.",t,
        );
    }
    for i:=0; i<5; i++ {
        err=tmp.Insert(i,i)
        test.BasicTest(nil,err,
            "Insert returned an error when it shouldn't have.",t,
        )
        test.BasicTest(i+1,tmp.Length(),
            "Insert did not increment the number of elements.",t,
        )
    }
    for i:=0; i<5; i++ {
        v,_:=tmp.Get(i)
        test.BasicTest(i,v,
            "Insert did not set the values correctly.",t,
        )
    }
    err=tmp.Insert(5,5)
    if !IsQueueFull(err) {
        test.FormatError(QueueFull(""),err,
            "Insert did not detect the queue was full.",t,
        );
    }
    test.BasicTest(5,tmp.Length(),
        "Insert incremented the number of elements when it shouldn't have.",t,
    )
    err=tmp.Insert(5,6)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Insert returned the wrong error.",t,
        );
    }
    test.BasicTest(5,tmp.Length(),
        "Insert incremented the number of elements when it shouldn't have.",t,
    )
}

func TestCircularBufferInsertBack(t *testing.T) {
    tmp,_:=NewCircularBuffer[int](5)
    for i:=0; i<5; i++ {
        err:=tmp.Insert(i,0)
        test.BasicTest(nil,err,
            "Insert returned an error when it shouldn't have.",t,
        )
        test.BasicTest(i+1,tmp.Length(),
            "Insert did not increment the number of elements.",t,
        )
    }
    for i:=0; i<5; i++ {
        v,_:=tmp.Get(i)
        test.BasicTest(4-i,v,
            "Insert did not set the values correctly.",t,
        )
    }
    err:=tmp.Insert(5,0)
    if !IsQueueFull(err) {
        test.FormatError(QueueFull(""),err,
            "Insert did not detect the queue was full.",t,
        );
    }
    test.BasicTest(5,tmp.Length(),
        "Insert incremented the number of elements when it shouldn't have.",t,
    )
    err=tmp.Insert(5,6)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Insert returned the wrong error.",t,
        );
    }
    test.BasicTest(5,tmp.Length(),
        "Insert incremented the number of elements when it shouldn't have.",t,
    )
}

func circularBufferInsertHelper(idx int, l int, t *testing.T){
    tmp,_:=NewCircularBuffer[int](l)
    for i:=0; i<l-1; i++ {
        tmp.PushBack(i)
    }
    err:=tmp.Insert(l-1,idx)
    test.BasicTest(nil,err,
        "Insert returned an error when it shouldn't have.",t,
    )
    test.BasicTest(l,tmp.Length(),
        "Insert did not increment the number of elements.",t,
    )
    for i:=0; i<l; i++ {
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
func TestCircularBufferRandomInsert(t *testing.T){
    for i:=0; i<5; i++ {
        circularBufferInsertHelper(i,5,t)
    }
}

func TestCircularBufferDeleteFront(t *testing.T){
    tmp,_:=NewCircularBuffer[int](5)
    err:=tmp.Delete(0)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Delete did not return the correct error.",t,
        )
    }
    for i:=0; i<5; i++ {
        tmp.PushBack(i)
    }
    err=tmp.Delete(6)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Delete did not return the correct error.",t,
        )
    }
    for i:=0; i<5; i++ {
        err:=tmp.Delete(0)
        test.BasicTest(nil,err,
            "Delete returned an error when it shouldn't have.",t,
        )
        test.BasicTest(4-i,tmp.Length(),
            "Delete did not update num elements correctly.",t,
        )
        for j:=0; j<5-i-1; j++ {
            v,_:=tmp.Get(j)
            test.BasicTest(i+j+1,v,
                "Delete did not remove the values correctly.",t,
            )
        }
    }
    test.BasicTest(0,tmp.Length(),
        "Delete did not decrement num elements correctly.",t,
    )
}

func TestCircularBufferDeleteBack(t *testing.T) {
    tmp,_:=NewCircularBuffer[int](5)
    err:=tmp.Delete(0)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Delete did not return the correct error.",t,
        )
    }
    for i:=0; i<5; i++ {
        tmp.PushBack(i)
    }
    err=tmp.Delete(6)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Delete did not return the correct error.",t,
        )
    }
    for i:=0; i<5; i++ {
        err:=tmp.Delete(tmp.Length()-1)
        test.BasicTest(nil,err,
            "Delete returned an error when it shouldn't have.",t,
        )
        test.BasicTest(4-i,tmp.Length(),
            "Delete did not update num elements correctly.",t,
        )
        for j:=0; j<5-i-1; j++ {
            v,_:=tmp.Get(j)
            test.BasicTest(j,v,
                "Delete did not remove the values correctly.",t,
            )
        }
    }
    test.BasicTest(0,tmp.Length(),
        "Delete did not decrement num elements correctly.",t,
    )
}

func circularBufferDeleteHelper(idx int, l int, t *testing.T){
    tmp,_:=NewCircularBuffer[int](l)
    for i:=0; i<l-1; i++ {
        tmp.PushBack(i)
    }
    err:=tmp.Delete(idx)
    test.BasicTest(nil,err,
        "Delete returned an error when it shouldn't have.",t,
    )
    test.BasicTest(l-1,tmp.Length(),
        "Delete did not increment the number of elements.",t,
    )
    for i:=0; i<l-1; i++ {
        var exp int
        v,_:=tmp.Get(i)
        if i<idx {
            exp=i
        } else if i>idx {
            exp=i-1
        }
        test.BasicTest(exp,v,
            "Delete did not remove the value in the correct place.",t,
        )
    }
}
func TestCircularBufferRandomDelete(t *testing.T){
    for i:=0; i<5; i++ {
        circularBufferInsertHelper(i,5,t)
    }
}

func TestCircularBufferAppend(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    for i:=0; i<5; i++ {
        res:=tmp.Append(i);
        test.BasicTest(nil,res,
            "Append returned an error when it should not have.",t,
        );
        test.BasicTest(i+1,tmp.numElems,
            "Append did not increment NumElems after adding value.",t,
        );
        test.BasicTest(i,tmp.vals[i],"Push did not save value.",t);
        test.BasicTest(0,tmp.startEnd.A,"Push modified the start index.",t);
        test.BasicTest(i,tmp.startEnd.B,
            "Append did not modify the end index.",t,
        );
    }
    res:=tmp.Append(5);
    if !IsQueueFull(res) {
        test.FormatError(QueueFull(""),res,
            "Append did not detect the queue was full.",t,
        );
    }
    test.BasicTest(5,tmp.numElems,
        "Append incremented NumElems when queue was full.",t,
    );
    test.BasicTest(0,tmp.startEnd.A,"Append modified the start index.",t);
    test.BasicTest(len(tmp.vals)-1,tmp.startEnd.B,
        "Append modified the end index when the queue was full.",t,
    );
}

func TestCircularBufferPushStartFromMiddle(t *testing.T) {
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    tmp.startEnd.A=2;
    tmp.startEnd.B=1;
    for i:=0; i<5; i++ {
        res:=tmp.PushBack(i);
        test.BasicTest(nil,res,
            "Push returned an error when it should not have.",t,
        );
        test.BasicTest(i+1,tmp.numElems,
            "Push did not increment NumElems after adding value.",t,
        );
        test.BasicTest(2,tmp.startEnd.A,"Push modified the start index.",t);
        if i<3 {
            test.BasicTest(i,tmp.vals[i+2],"Push did not save value.",t);
            test.BasicTest(i+2,tmp.startEnd.B,
                "Push did not modify the end index.",t,
            );
        } else {
            test.BasicTest(i,tmp.vals[i-3],"Push did not save value.",t);
            test.BasicTest(i-3,tmp.startEnd.B,
                "Push did not modify the end index.",t,
            );
        }
    }
    res:=tmp.PushBack(5);
    if !IsQueueFull(res) {
        test.FormatError(QueueFull(""),res,
            "Push did not detect the queue was full.",t,
        );
    }
    test.BasicTest(5,tmp.numElems,
        "Push incremented NumElems when queue was full.",t,
    );
    test.BasicTest(2,tmp.startEnd.A,"Push modified the start index.",t);
    test.BasicTest(1,tmp.startEnd.B,
        "Push modified the end index when the queue was full.",t,
    );
}

func TestCircularBufferForcePushFront(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    for i:=0; i<5; i++ {
        tmp.ForcePushFront(i)
    }
    for i:=0; i<5; i++ {
        v,err:=tmp.Get(4-i)
        test.BasicTest(i,v,
            "ForcePushFront did not correctly save values.",t,
        )
        test.BasicTest(nil,err,
            "Get returned an error when it should not have.",t,
        )
    }
    tmp.ForcePushFront(5)
    for i:=1; i<6; i++ {
        v,err:=tmp.Get(i-1)
        test.BasicTest(6-i,v,
            "ForcePushFront did not correctly save values.",t,
        )
        test.BasicTest(nil,err,
            "Get returned an error when it should not have.",t,
        )
    }
}

func TestCircularBufferForcePushBack(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    for i:=0; i<5; i++ {
        tmp.ForcePushBack(i)
    }
    for i:=0; i<5; i++ {
        v,err:=tmp.Get(i)
        test.BasicTest(i,v,
            "ForcePushBack did not correctly save values.",t,
        )
        test.BasicTest(nil,err,
            "Get returned an error when it should not have.",t,
        )
    }
    tmp.ForcePushBack(5)
    for i:=1; i<6; i++ {
        v,err:=tmp.Get(i-1)
        test.BasicTest(i,v,
            "ForcePushBack did not correctly save values.",t,
        )
        test.BasicTest(nil,err,
            "Get returned an error when it should not have.",t,
        )
    }
}

func testCircularBufferGetHelper(c CircularBuffer[int], t *testing.T){
    for i:=0; i<len(c.vals); i++ {
        c.PushBack(i);
    }
    for i:=0; i<len(c.vals); i++ {
        v,err:=c.Get(i);
        test.BasicTest(nil,err,
            "Get returned an error when it should not have.",t,
        );
        test.BasicTest(i,v,"Get did not return correct value.",t);
    }
    _,err:=c.Get(5);
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Get did not return the correct error.",t,
        );
    }
    _,err=c.Get(-1);
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Get did not return the correct error.",t,
        );
    }
}
func TestCircularBufferGet(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    _,err=tmp.Get(0)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Get on a zero length queue did not return the correct error.",t,
        )
    }
    testCircularBufferGetHelper(tmp,t);
    tmp,err=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    tmp.startEnd.A=2;
    tmp.startEnd.B=1;
    testCircularBufferGetHelper(tmp,t);
}

func TestCircularBufferPeekFront(t *testing.T) {
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    _,err=tmp.PeekFront()
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Peek front returned the incorrect error.",t,
        )
    }
    tmp.PushBack(1)
    v,err:=tmp.PeekFront()
    test.BasicTest(1,v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PushBack(2)
    v,err=tmp.PeekFront()
    test.BasicTest(1,v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PushBack(3)
    v,err=tmp.PeekFront()
    test.BasicTest(1,v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PopFront()
    v,err=tmp.PeekFront()
    test.BasicTest(2,v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PopFront()
    v,err=tmp.PeekFront()
    test.BasicTest(3,v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PopFront()
    _,err=tmp.PeekFront()
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Peek front returned the incorrect error.",t,
        )
    }
}

func testCircularBufferGetPntrHelper(c CircularBuffer[int], t *testing.T){
    for i:=0; i<len(c.vals); i++ {
        c.PushBack(i);
    }
    for i:=0; i<len(c.vals); i++ {
        v,err:=c.GetPntr(i);
        test.BasicTest(nil,err,
            "Get returned an error when it should not have.",t,
        );
        test.BasicTest(i,*v,"Get did not return correct value.",t);
        *v=100;
    }
    for i:=0; i<len(c.vals); i++ {
        v,err:=c.GetPntr(i);
        test.BasicTest(nil,err,
            "Get returned an error when it should not have.",t,
        );
        test.BasicTest(100,*v,"Get did not return correct value.",t);
    }
    _,err:=c.GetPntr(5);
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Get did not return the correct error.",t,
        );
    }
    _,err=c.GetPntr(-1);
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Get did not return the correct error.",t,
        );
    }
}
func TestCircularBufferGetPntr(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    v,err:=tmp.GetPntr(0)
    test.BasicTest((*int)(nil),v, "Get pntr returned the incorrect value.",t)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Get on a zero length queue did not return the correct error.",t,
        )
    }
    testCircularBufferGetPntrHelper(tmp,t);
    tmp,err=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    tmp.startEnd.A=2;
    tmp.startEnd.B=1;
    testCircularBufferGetPntrHelper(tmp,t);
}

func TestCircularBufferPeekPntrFront(t *testing.T) {
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    _,err=tmp.PeekPntrFront()
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Peek front returned the incorrect error.",t,
        )
    }
    tmp.PushBack(1)
    v,err:=tmp.PeekPntrFront()
    test.BasicTest(1,*v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PushBack(2)
    v,err=tmp.PeekPntrFront()
    test.BasicTest(1,*v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PushBack(3)
    v,err=tmp.PeekPntrFront()
    test.BasicTest(1,*v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PopFront()
    v,err=tmp.PeekPntrFront()
    test.BasicTest(2,*v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PopFront()
    v,err=tmp.PeekPntrFront()
    test.BasicTest(3,*v,"Peek front returned the incorrect value.",t);
    test.BasicTest(nil,err,
        "Peek front returned an error when it shouldn't have.",t,
    );
    tmp.PopFront()
    _,err=tmp.PeekPntrFront()
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Peek front returned the incorrect error.",t,
        )
    }
}

func TestCircularBufferSet(t *testing.T){
    tmp,_:=NewCircularBuffer[int](5)
    err:=tmp.Set(4,0)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Calling set out of bounds did not return the correct error.",t,
        )
    }
    for i:=0; i<4; i++ {
        tmp.PushBack(i)
    }
    err=tmp.Set(4,4)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Calling set out of bounds did not return the correct error.",t,
        )
    }
    err=tmp.Set(4,-1)
    if !customerr.IsValOutsideRange(err) {
        test.FormatError(customerr.ValOutsideRange(""),err,
            "Calling set out of bounds did not return the correct error.",t,
        )
    }
    tmp.PushBack(4)
    for i:=0; i<5; i++ {
        err=tmp.Set(i+1,i)
        test.BasicTest(nil,err,
            "Set returned an error when it should not have.",t,
        )
    }
    for i:=0; i<5; i++ {
        v,_:=tmp.Get(i);
        test.BasicTest(i+1,v,
            "Set did not set the correct value.",t,
        )
    }
}

func TestCircularBufferPop(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    tmp.PushBack(0);
    for i:=1; i<21; i++ {
        tmp.PushBack(i);
        v,err:=tmp.PopFront();
        test.BasicTest(nil,err,
            "Pop returned an error when it should not have.",t,
        );
        test.BasicTest(i-1,v,"Pop did not return correct value.",t);
    }
    v,err:=tmp.PopFront();
    test.BasicTest(nil,err,
        "Pop returned an error when it should not have.",t,
    );
    test.BasicTest(20,v,"Pop did not return correct value.",t);
    _,err=tmp.PopFront();
    if !IsQueueEmpty(err) {
        test.FormatError(QueueEmpty(""),err,
            "Pop did not return correct error.",t,
        );
    }
}

func testCircularBufferElemsHelper(c CircularBuffer[int], t *testing.T){
    for i:=0; i<len(c.vals); i++ {
        c.PushBack(i);
    }
    c.Elems().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        test.BasicTest(index,val,"Element was skipped while iterating.",t);
        return iter.Continue,nil;
    });
}
func TestCircularBufferElems(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testCircularBufferElemsHelper(tmp,t);
    tmp,err=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    tmp.startEnd.A=2;
    tmp.startEnd.B=1;
    testCircularBufferElemsHelper(tmp,t);
}

func testCircularBufferPntrElemsHelper(c CircularBuffer[int], t *testing.T){
    for i:=0; i<len(c.vals); i++ {
        c.PushBack(i);
    }
    c.PntrElems().ForEach(func(index int, val *int) (iter.IteratorFeedback, error) {
        test.BasicTest(index,*val,"Element was skipped while iterating.",t);
        *val=100;
        return iter.Continue,nil;
    });
    c.Elems().ForEach(func(index int, val int) (iter.IteratorFeedback, error) {
        test.BasicTest(100,val,"Element was not updated while iterating.",t);
        return iter.Continue,nil;
    });
}
func TestCircularBufferPntrElems(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testCircularBufferElemsHelper(tmp,t);
    tmp,err=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    tmp.startEnd.A=2;
    tmp.startEnd.B=1;
    testCircularBufferElemsHelper(tmp,t);
}

func TestCircularBufferStaticQueueTypeInterface(t *testing.T) {
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    staticQueueInterfaceTypeCheck[int](&tmp);
}

func TestCircularBufferStaticStackTypeInterface(t *testing.T){
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    staticStackInterfaceTypeCheck[int](&tmp);
}

func TestCircularBufferStaticVectorTypeInterface(t *testing.T) {
    tmp,err:=NewCircularBuffer[int](5);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    staticVectorInterfaceTypeCheck[int](&tmp);
}

func TestCircularBufferDynQueueTypeInterface(t *testing.T){
    test.Panics(
        func () {
            var c any
            c,_=NewCircularBuffer[int](5)
            c2:=c.(dynamic.Queue[int])
            _=c2
        }, 
        "Code did not panic when casting a circular queue to a dynamic queue.",t,
    )
    test.Panics(
        func () {
            var c any
            c,_=NewCircularBuffer[int](5)
            c2:=c.(dynamic.ReadQueue[int])
            _=c2
        }, 
        "Code did not panic when casting a circular queue to a dynamic queue.",t,
    )
    test.Panics(
        func () {
            var c any
            c,_=NewCircularBuffer[int](5)
            c2:=c.(dynamic.WriteQueue[int])
            _=c2
        }, 
        "Code did not panic when casting a circular queue to a dynamic queue.",t,
    )
}
