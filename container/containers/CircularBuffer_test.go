package containers

import (
	"testing"

	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func TestWrapAroundIntAdd(t *testing.T) {
	v:=wrapingIndex(0)
	v=v.Add(1,5)
	test.Eq(wrapingIndex(1),v,t)
	v=v.Add(1,5)
	test.Eq(wrapingIndex(2),v,t)
	v=v.Add(1,5)
	test.Eq(wrapingIndex(3),v,t)
	v=v.Add(1,5)
	test.Eq(wrapingIndex(4),v,t)
	v=v.Add(1,5)
	test.Eq(wrapingIndex(0),v,t)
	v=v.Add(1,5)
	test.Eq(wrapingIndex(1),v,t)
	v=v.Add(2,5)
	test.Eq(wrapingIndex(3),v,t)
	v=v.Add(2,5)
	test.Eq(wrapingIndex(0),v,t)
	v=v.Add(3,5)
	test.Eq(wrapingIndex(3),v,t)
	v=v.Add(3,5)
	test.Eq(wrapingIndex(1),v,t)
	v=v.Add(4,5)
	test.Eq(wrapingIndex(0),v,t)
	v=v.Add(5,5)
	test.Eq(wrapingIndex(0),v,t)
	v=v.Add(6,5)
	test.Eq(wrapingIndex(1),v,t)
}

func TestWrapAroundIntSub(t *testing.T) {
	v:=wrapingIndex(5)
	v=v.Sub(1,5)
	test.Eq(wrapingIndex(4),v,t)
	v=v.Sub(1,5)
	test.Eq(wrapingIndex(3),v,t)
	v=v.Sub(1,5)
	test.Eq(wrapingIndex(2),v,t)
	v=v.Sub(1,5)
	test.Eq(wrapingIndex(1),v,t)
	v=v.Sub(1,5)
	test.Eq(wrapingIndex(0),v,t)
	v=v.Sub(1,5)
	test.Eq(wrapingIndex(4),v,t)
	v=v.Sub(2,5)
	test.Eq(wrapingIndex(2),v,t)
	v=v.Sub(2,5)
	test.Eq(wrapingIndex(0),v,t)
	v=v.Sub(3,5)
	test.Eq(wrapingIndex(2),v,t)
	v=v.Sub(3,5)
	test.Eq(wrapingIndex(4),v,t)
	v=v.Sub(4,5)
	test.Eq(wrapingIndex(0),v,t)
	v=v.Sub(4,5)
	test.Eq(wrapingIndex(1),v,t)
	v=v.Sub(5,5)
	test.Eq(wrapingIndex(1),v,t)
	v=v.Sub(6,5)
	test.Eq(wrapingIndex(0),v,t)
}

func TestWrapAroundIntGetProperIndex(t *testing.T){
	w:=wrapingIndex(0)
	test.Eq(wrapingIndex(0),w.GetProperIndex(-5,5),t)
	test.Eq(wrapingIndex(1),w.GetProperIndex(-4,5),t)
	test.Eq(wrapingIndex(2),w.GetProperIndex(-3,5),t)
	test.Eq(wrapingIndex(3),w.GetProperIndex(-2,5),t)
	test.Eq(wrapingIndex(4),w.GetProperIndex(-1,5),t)
	test.Eq(wrapingIndex(0),w.GetProperIndex(0,5),t)
	test.Eq(wrapingIndex(1),w.GetProperIndex(1,5),t)
	test.Eq(wrapingIndex(2),w.GetProperIndex(2,5),t)
	test.Eq(wrapingIndex(3),w.GetProperIndex(3,5),t)
	test.Eq(wrapingIndex(4),w.GetProperIndex(4,5),t)
	test.Eq(wrapingIndex(0),w.GetProperIndex(5,5),t)
	test.Eq(wrapingIndex(1),w.GetProperIndex(6,5),t)
	test.Eq(wrapingIndex(2),w.GetProperIndex(7,5),t)
	test.Eq(wrapingIndex(3),w.GetProperIndex(8,5),t)
	test.Eq(wrapingIndex(4),w.GetProperIndex(9,5),t)
}

//go:generate go run interfaceTest.go -type=CircularBuffer -category=static -interface=Vector -genericDecl=[int] -factory=generateCircularBuffer
//go:generate go run interfaceTest.go -type=SyncedCircularBuffer -category=static -interface=Vector -genericDecl=[int] -factory=generateSyncedCircularBuffer

func generateCircularBuffer(capacity int) CircularBuffer[int,widgets.BuiltinInt] {
    c,_:=NewCircularBuffer[int,widgets.BuiltinInt](capacity)
    return c
}

func generateSyncedCircularBuffer(capacity int) SyncedCircularBuffer[int,widgets.BuiltinInt] {
    c,_:=NewSyncedCircularBuffer[int,widgets.BuiltinInt](capacity)
    return c
}
// 
// func TestNewCircularBufferBadSize(t *testing.T) {
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](0);
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Invalid buffer size did not raise the correct error.",t,
//     )
//     test.BasicTest(0,len(tmp.vals),
//         "NewCircularBuffer added values to empty buffer during initialization.",t,
//     );
//     test.BasicTest(0,cap(tmp.vals),
//         "NewCircularBuffer did not set capacity correctly.",t,
//     );
//     test.BasicTest(0,tmp.startEnd.A,
//         "NewCircularBuffer added values to empty buffer during initialization.",t,
//     );
//     test.BasicTest(0,tmp.startEnd.B,
//         "NewCircularBuffer added values to empty buffer during initialization.",t,
//     );
// }
// 
// func TestCircularBufferPushToFront(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     for i:=0; i<5; i++ {
//         res:=tmp.PushFront(i);
//         test.BasicTest(nil,res,
//             "Push returned an error when it should not have.",t,
//         );
//         test.BasicTest(i+1,tmp.numElems,
//             "Push did not increment NumElems after adding value.",t,
//         );
//         test.BasicTest(i,tmp.vals[4-i],"Push did not save value.",t);
//         test.BasicTest(4-i,tmp.startEnd.A,
//             "Push did not modify the start index.",t,
//         );
//         test.BasicTest(4,tmp.startEnd.B,"Push modified the end index.",t);
//     }
//     res:=tmp.PushBack(5);
//     test.ContainsError(containerTypes.Full,res,
//         "Push did not detect the buffer was full.",t,
//     )
//     test.BasicTest(5,tmp.numElems,
//         "Push incremented NumElems when buffer was full.",t,
//     );
//     test.BasicTest(0,tmp.startEnd.A,"Push modified the start index.",t);
//     test.BasicTest(4,tmp.startEnd.B,
//         "Push modified the end index when the buffer was full.",t,
//     );
// }
// 
// func TestCircularBufferPushToBack(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     for i:=0; i<5; i++ {
//         res:=tmp.PushBack(i);
//         test.BasicTest(nil,res,
//             "Push returned an error when it should not have.",t,
//         );
//         test.BasicTest(i+1,tmp.numElems,
//             "Push did not increment NumElems after adding value.",t,
//         );
//         test.BasicTest(i,tmp.vals[i],"Push did not save value.",t);
//         test.BasicTest(0,tmp.startEnd.A,"Push modified the start index.",t);
//         test.BasicTest(i,tmp.startEnd.B,
//             "Push did not modify the end index.",t,
//         );
//     }
//     res:=tmp.PushBack(5);
//     test.ContainsError(containerTypes.Full,res,
//         "Push did not detect the buffer was full.",t,
//     )
//     test.BasicTest(5,tmp.numElems,
//         "Push incremented NumElems when buffer was full.",t,
//     );
//     test.BasicTest(0,tmp.startEnd.A,"Push modified the start index.",t);
//     test.BasicTest(len(tmp.vals)-1,tmp.startEnd.B,
//         "Push modified the end index when the buffer was full.",t,
//     );
// }
// 
// func TestCircularBufferPushFront(t *testing.T){
//     tmp,_:=NewCircularBuffer[int,widgets.BuiltinInt](5)
//     err:=tmp.Push(1,1)
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Push returned the wrong error.",t,
//     )
//     for i:=0; i<5; i++ {
//         err=tmp.Push(i,i)
//         test.BasicTest(nil,err,
//             "Push returned an error when it shouldn't have.",t,
//         )
//         test.BasicTest(i+1,tmp.Length(),
//             "Push did not increment the number of elements.",t,
//         )
//     }
//     for i:=0; i<5; i++ {
//         v,_:=tmp.Get(i)
//         test.BasicTest(i,v,
//             "Push did not set the values correctly.",t,
//         )
//     }
//     err=tmp.Push(5,5)
//     test.ContainsError(containerTypes.Full,err,
//         "Push did not detect the buffer was full.",t,
//     )
//     test.BasicTest(5,tmp.Length(),
//         "Push incremented the number of elements when it shouldn't have.",t,
//     )
//     err=tmp.Push(6,5)
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Push returned the wrong error.",t,
//     )
//     test.BasicTest(5,tmp.Length(),
//         "Push incremented the number of elements when it shouldn't have.",t,
//     )
// }
// 
// func TestCircularBufferPushBack(t *testing.T) {
//     tmp,_:=NewCircularBuffer[int,widgets.BuiltinInt](5)
//     for i:=0; i<5; i++ {
//         err:=tmp.Push(0,i)
//         test.BasicTest(nil,err,
//             "Push returned an error when it shouldn't have.",t,
//         )
//         test.BasicTest(i+1,tmp.Length(),
//             "Push did not increment the number of elements.",t,
//         )
//     }
//     for i:=0; i<5; i++ {
//         v,_:=tmp.Get(i)
//         test.BasicTest(4-i,v,
//             "Push did not set the values correctly.",t,
//         )
//     }
//     err:=tmp.Push(0,5)
//     test.ContainsError(containerTypes.Full,err,
//         "Push did not detect the buffer was full.",t,
//     )
//     test.BasicTest(5,tmp.Length(),
//         "Push incremented the number of elements when it shouldn't have.",t,
//     )
//     err=tmp.Push(6,5)
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Push returned the wrong error.",t,
//     )
//     test.BasicTest(5,tmp.Length(),
//         "Push incremented the number of elements when it shouldn't have.",t,
//     )
// }
// 
// func circularBufferPushHelper(idx int, l int, t *testing.T){
//     tmp,_:=NewCircularBuffer[int,widgets.BuiltinInt](l)
//     for i:=0; i<l-1; i++ {
//         tmp.PushBack(i)
//     }
//     err:=tmp.Push(idx,l-1)
//     test.BasicTest(nil,err,
//         "Push returned an error when it shouldn't have.",t,
//     )
//     test.BasicTest(l,tmp.Length(),
//         "Push did not increment the number of elements.",t,
//     )
//     for i:=0; i<l; i++ {
//         var exp int
//         v,_:=tmp.Get(i)
//         if i<idx {
//             exp=i
//         } else if i==idx {
//             exp=l-1
//         } else {
//             exp=i-1
//         }
//         test.BasicTest(exp,v,
//             "Push did not put the value in the correct place.",t,
//         )
//     }
// }
// func TestCircularBufferRandomPush(t *testing.T){
//     for i:=0; i<5; i++ {
//         circularBufferPushHelper(i,5,t)
//     }
// }
// 
// func TestCircularBufferPush(t *testing.T){
//     tmp,_:=NewCircularBuffer[int,widgets.BuiltinInt](5)
//     err:=tmp.Push(0,1,2,3)
//     test.BasicTest(nil,err,
//         "Push returned an error when it shouldn't have.",t,
//     )
//     for i:=0; i<3; i++ {
//         v,_:=tmp.Get(i)
//         test.BasicTest(i+1,v,
//             "Pushing many values did not work correctly.",t,
//         )
//     }
//     err=tmp.Push(3,4,5)
//     test.BasicTest(nil,err,
//         "Push returned an error when it shouldn't have.",t,
//     )
//     for i:=0; i<5; i++ {
//         v,_:=tmp.Get(i)
//         test.BasicTest(i+1,v,
//             "Pushing many values did not work correctly.",t,
//         )
//     }
//     err=tmp.Push(5,6)
//     test.ContainsError(containerTypes.Full,err,
//         "Push did not detect the buffer was full.",t,
//     )
//     tmp,_=NewCircularBuffer[int,widgets.BuiltinInt](5)
//     tmp.Push(0,1,2)
//     err=tmp.Push(2,3,4,5,6)
//     for i:=0; i<5; i++ {
//         v,_:=tmp.Get(i)
//         test.BasicTest(i+1,v,
//             "Push did not put the values in the right position.",t,
//         )
//     }
//     test.ContainsError(containerTypes.Full,err,
//         "Push did not detect the buffer was full.",t,
//     )
// }
// 
func circularBufferDeleteFrontHelper(l int, startIdx int, t *testing.T){
    tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5)
	tmp.start=wrapingIndex(startIdx)
	test.Nil(err,t)
    err=tmp.Delete(0)
    test.ContainsError(customerr.ValOutsideRange,err,t)
    for i:=0; i<5; i++ {
        tmp.PushBack(i)
    }
    err=tmp.Delete(6)
    test.ContainsError(customerr.ValOutsideRange,err,t)
    for i:=0; i<5; i++ {
        err:=tmp.Delete(0)
        test.Nil(err,t)
        test.Eq(4-i,tmp.Length(),t)
        for j:=0; j<5-i-1; j++ {
            v,err:=tmp.Get(j)
			test.Nil(err,t)
            test.Eq(i+j+1,v,t)
        }
    }
    test.Eq(0,tmp.Length(),t)
}
func TestCircularBufferDeleteFront(t *testing.T){
	for i:=0; i<5; i++ {
		circularBufferDeleteFrontHelper(5,i,t)
	}
}

func circularBufferDeleteBackHelper(l int, startIdx int, t *testing.T){
    tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5)
	test.Nil(err,t)
	tmp.start=wrapingIndex(startIdx)
    err=tmp.Delete(0)
    test.ContainsError(customerr.ValOutsideRange,err,t)
    for i:=0; i<5; i++ {
        tmp.PushBack(i)
    }
    err=tmp.Delete(6)
    test.ContainsError(customerr.ValOutsideRange,err,t)
    for i:=0; i<5; i++ {
        err:=tmp.Delete(tmp.Length()-1)
        test.Nil(err,t)
        test.Eq(4-i,tmp.Length(),t)
        for j:=0; j<5-i-1; j++ {
            v,_:=tmp.Get(j)
            test.Eq(j,v,t)
        }
    }
    test.Eq(0,tmp.Length(),t)
}
func TestCircularBufferDeleteBack(t *testing.T) {
	for i:=0; i<5; i++ {
		circularBufferDeleteBackHelper(5,i,t)
	}
}

func circularBufferDeleteHelper(idx int, l int, startIdx int, t *testing.T){
    tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](l)
	test.Nil(err,t)
	tmp.start=wrapingIndex(startIdx)
    for i:=0; i<l; i++ {
        tmp.PushBack(i)
    }
    err=tmp.Delete(idx)
    test.Nil(err,t)
    test.Eq(l-1,tmp.Length(),t)
    for i:=0; i<l-1; i++ {
        var exp int
        if i<idx {
            exp=i
        } else if i>=idx {
            exp=i+1
        }
        v,err:=tmp.Get(i)
		test.Nil(err,t)
        test.Eq(exp,v,t)
    }
}
func TestCircularBufferRandomDelete(t *testing.T){
    for i:=0; i<5; i++ {
		for j:=0; j<5; j++ {
			circularBufferDeleteHelper(i,5,j,t)
		}
    }
}

// func TestCircularBufferAppend(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     for i:=0; i<5; i++ {
//         res:=tmp.Append(i);
//         test.BasicTest(nil,res,
//             "Append returned an error when it should not have.",t,
//         );
//         test.BasicTest(i+1,tmp.numElems,
//             "Append did not increment NumElems after adding value.",t,
//         );
//         test.BasicTest(i,tmp.vals[i],"Push did not save value.",t);
//         test.BasicTest(0,tmp.startEnd.A,"Push modified the start index.",t);
//         test.BasicTest(i,tmp.startEnd.B,
//             "Append did not modify the end index.",t,
//         );
//     }
//     res:=tmp.Append(5);
//     test.ContainsError(containerTypes.Full,res,
//         "Append did not detect the buffer was full.",t,
//     )
//     test.BasicTest(5,tmp.numElems,
//         "Append incremented NumElems when buffer was full.",t,
//     );
//     test.BasicTest(0,tmp.startEnd.A,"Append modified the start index.",t);
//     test.BasicTest(len(tmp.vals)-1,tmp.startEnd.B,
//         "Append modified the end index when the buffer was full.",t,
//     );
// }
// 
// func TestCircularBufferPushStartFromMiddle(t *testing.T) {
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     tmp.startEnd.A=2;
//     tmp.startEnd.B=1;
//     for i:=0; i<5; i++ {
//         res:=tmp.PushBack(i);
//         test.BasicTest(nil,res,
//             "Push returned an error when it should not have.",t,
//         );
//         test.BasicTest(i+1,tmp.numElems,
//             "Push did not increment NumElems after adding value.",t,
//         );
//         test.BasicTest(2,tmp.startEnd.A,"Push modified the start index.",t);
//         if i<3 {
//             test.BasicTest(i,tmp.vals[i+2],"Push did not save value.",t);
//             test.BasicTest(i+2,tmp.startEnd.B,
//                 "Push did not modify the end index.",t,
//             );
//         } else {
//             test.BasicTest(i,tmp.vals[i-3],"Push did not save value.",t);
//             test.BasicTest(i-3,tmp.startEnd.B,
//                 "Push did not modify the end index.",t,
//             );
//         }
//     }
//     res:=tmp.PushBack(5);
//     test.ContainsError(containerTypes.Full,res,
//         "Append did not detect the buffer was full.",t,
//     )
//     test.BasicTest(5,tmp.numElems,
//         "Push incremented NumElems when buffer was full.",t,
//     );
//     test.BasicTest(2,tmp.startEnd.A,"Push modified the start index.",t);
//     test.BasicTest(1,tmp.startEnd.B,
//         "Push modified the end index when the buffer was full.",t,
//     );
// }
// 
// func TestCircularBufferForcePushFront(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     for i:=0; i<5; i++ {
//         tmp.ForcePushFront(i)
//     }
//     for i:=0; i<5; i++ {
//         v,err:=tmp.Get(4-i)
//         test.BasicTest(i,v,
//             "ForcePushFront did not correctly save values.",t,
//         )
//         test.BasicTest(nil,err,
//             "Get returned an error when it should not have.",t,
//         )
//     }
//     tmp.ForcePushFront(5)
//     for i:=1; i<6; i++ {
//         v,err:=tmp.Get(i-1)
//         test.BasicTest(6-i,v,
//             "ForcePushFront did not correctly save values.",t,
//         )
//         test.BasicTest(nil,err,
//             "Get returned an error when it should not have.",t,
//         )
//     }
// }
// 
// func TestCircularBufferForcePushBack(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     for i:=0; i<5; i++ {
//         tmp.ForcePushBack(i)
//     }
//     for i:=0; i<5; i++ {
//         v,err:=tmp.Get(i)
//         test.BasicTest(i,v,
//             "ForcePushBack did not correctly save values.",t,
//         )
//         test.BasicTest(nil,err,
//             "Get returned an error when it should not have.",t,
//         )
//     }
//     tmp.ForcePushBack(5)
//     for i:=1; i<6; i++ {
//         v,err:=tmp.Get(i-1)
//         test.BasicTest(i,v,
//             "ForcePushBack did not correctly save values.",t,
//         )
//         test.BasicTest(nil,err,
//             "Get returned an error when it should not have.",t,
//         )
//     }
// }
// 
// func testCircularBufferGetHelper(
//     c CircularBuffer[int,widgets.BuiltinInt], 
//     t *testing.T,
// ){
//     for i:=0; i<len(c.vals); i++ {
//         c.PushBack(i);
//     }
//     for i:=0; i<len(c.vals); i++ {
//         v,err:=c.Get(i);
//         test.BasicTest(nil,err,
//             "Get returned an error when it should not have.",t,
//         );
//         test.BasicTest(i,v,"Get did not return correct value.",t);
//     }
//     _,err:=c.Get(5);
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Get did not return the correct error.",t,
//     )
//     _,err=c.Get(-1);
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Get did not return the correct error.",t,
//     )
// }
// func TestCircularBufferGet(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     _,err=tmp.Get(0)
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Get did not return the correct error.",t,
//     )
//     testCircularBufferGetHelper(tmp,t);
//     tmp,err=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     tmp.startEnd.A=2;
//     tmp.startEnd.B=1;
//     testCircularBufferGetHelper(tmp,t);
// }
// 
// func TestCircularBufferPeekFront(t *testing.T) {
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     _,err=tmp.PeekFront()
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Peek front returned the incorrect error.",t,
//     )
//     tmp.PushBack(1)
//     v,err:=tmp.PeekFront()
//     test.BasicTest(1,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PushBack(2)
//     v,err=tmp.PeekFront()
//     test.BasicTest(1,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PushBack(3)
//     v,err=tmp.PeekFront()
//     test.BasicTest(1,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     v,err=tmp.PeekFront()
//     test.BasicTest(2,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     v,err=tmp.PeekFront()
//     test.BasicTest(3,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     _,err=tmp.PeekFront()
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Peek front returned the incorrect error.",t,
//     )
// }
// 
// func TestCircularBufferPeekBack(t *testing.T) {
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     _,err=tmp.PeekBack()
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Peek front returned the incorrect error.",t,
//     )
//     tmp.PushBack(1)
//     v,err:=tmp.PeekBack()
//     test.BasicTest(1,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PushBack(2)
//     v,err=tmp.PeekBack()
//     test.BasicTest(2,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PushBack(3)
//     v,err=tmp.PeekBack()
//     test.BasicTest(3,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     v,err=tmp.PeekBack()
//     test.BasicTest(3,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     v,err=tmp.PeekBack()
//     test.BasicTest(3,v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     _,err=tmp.PeekFront()
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Peek front returned the incorrect error.",t,
//     )
// }
// 
// func testCircularBufferGetPntrHelper(
//     c CircularBuffer[int,widgets.BuiltinInt], 
//     t *testing.T,
// ){
//     for i:=0; i<len(c.vals); i++ {
//         c.PushBack(i);
//     }
//     for i:=0; i<len(c.vals); i++ {
//         v,err:=c.GetPntr(i);
//         test.BasicTest(nil,err,
//             "Get returned an error when it should not have.",t,
//         );
//         test.BasicTest(i,*v,"Get did not return correct value.",t);
//         *v=100;
//     }
//     for i:=0; i<len(c.vals); i++ {
//         v,err:=c.GetPntr(i);
//         test.BasicTest(nil,err,
//             "Get returned an error when it should not have.",t,
//         );
//         test.BasicTest(100,*v,"Get did not return correct value.",t);
//     }
//     _,err:=c.GetPntr(5);
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Get did not return the correct error.",t,
//     )
//     _,err=c.GetPntr(-1);
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Get did not return the correct error.",t,
//     )
// }
// func TestCircularBufferGetPntr(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     v,err:=tmp.GetPntr(0)
//     test.BasicTest((*int)(nil),v, "Get pntr returned the incorrect value.",t)
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Get on a zero length buffer did not return the correct error.",t,
//     )
//     testCircularBufferGetPntrHelper(tmp,t);
//     tmp,err=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     tmp.startEnd.A=2;
//     tmp.startEnd.B=1;
//     testCircularBufferGetPntrHelper(tmp,t);
// }
// 
// func TestCircularBufferPeekPntrFront(t *testing.T) {
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     _,err=tmp.PeekPntrFront()
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Peek front returned the incorrect error.",t,
//     )
//     tmp.PushBack(1)
//     v,err:=tmp.PeekPntrFront()
//     test.BasicTest(1,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PushBack(2)
//     v,err=tmp.PeekPntrFront()
//     test.BasicTest(1,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PushBack(3)
//     v,err=tmp.PeekPntrFront()
//     test.BasicTest(1,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     v,err=tmp.PeekPntrFront()
//     test.BasicTest(2,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     v,err=tmp.PeekPntrFront()
//     test.BasicTest(3,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek front returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     _,err=tmp.PeekPntrFront()
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Peek front returned the incorrect error.",t,
//     )
// }
// 
// func TestCircularBufferPeekPntrBack(t *testing.T) {
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     _,err=tmp.PeekPntrBack()
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Peek back returned the incorrect error.",t,
//     )
//     tmp.PushBack(1)
//     v,err:=tmp.PeekPntrBack()
//     test.BasicTest(1,*v,"Peek back returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek back returned an error when it shouldn't have.",t,
//     );
//     tmp.PushBack(2)
//     v,err=tmp.PeekPntrBack()
//     test.BasicTest(2,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek back returned an error when it shouldn't have.",t,
//     );
//     tmp.PushBack(3)
//     v,err=tmp.PeekPntrBack()
//     test.BasicTest(3,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek back returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     v,err=tmp.PeekPntrBack()
//     test.BasicTest(3,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek back returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     v,err=tmp.PeekPntrBack()
//     test.BasicTest(3,*v,"Peek front returned the incorrect value.",t);
//     test.BasicTest(nil,err,
//         "Peek back returned an error when it shouldn't have.",t,
//     );
//     tmp.PopFront()
//     _,err=tmp.PeekPntrBack()
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Peek back returned the incorrect error.",t,
//     )
// }
// 
// func TestCircularBufferEmplace(t *testing.T){
//     tmp,_:=NewCircularBuffer[int,widgets.BuiltinInt](5)
//     err:=tmp.Emplace(0,4)
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Calling emplace out of bounds did not return the correct error.",t,
//     )
//     for i:=0; i<4; i++ {
//         tmp.PushBack(i)
//     }
//     err=tmp.Emplace(4,4)
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Calling emplace out of bounds did not return the correct error.",t,
//     )
//     err=tmp.Emplace(-1,4)
//     test.ContainsError(customerr.ValOutsideRange,err,
//         "Calling emplace out of bounds did not return the correct error.",t,
//     )
//     tmp.PushBack(4)
//     for i:=0; i<5; i++ {
//         err=tmp.Emplace(i,i+1)
//         test.BasicTest(nil,err,
//             "Emplace returned an error when it should not have.",t,
//         )
//     }
//     for i:=0; i<5; i++ {
//         v,_:=tmp.Get(i);
//         test.BasicTest(i+1,v,
//             "Emplace did not set the correct value.",t,
//         )
//     }
// }
// 
// func TestCircularBufferPopFront(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     tmp.PushBack(0);
//     for i:=1; i<21; i++ {
//         tmp.PushBack(i);
//         v,err:=tmp.PopFront();
//         test.BasicTest(nil,err,
//             "Pop returned an error when it should not have.",t,
//         );
//         test.BasicTest(i-1,v,"Pop did not return correct value.",t);
//     }
//     v,err:=tmp.PopFront();
//     test.BasicTest(nil,err,
//         "Pop returned an error when it should not have.",t,
//     );
//     test.BasicTest(20,v,"Pop did not return correct value.",t);
//     _,err=tmp.PopFront();
//     test.ContainsError(containerTypes.Empty,err,"Pop did not return correct error.",t)
// }
// 
// func TestCircularBufferPopBack(t *testing.T){
//     tmp,err:=NewCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     tmp.PushBack(0);
//     for i:=1; i<21; i++ {
//         tmp.PushFront(i);
//         v,err:=tmp.PopBack();
//         test.BasicTest(nil,err,
//             "Pop returned an error when it should not have.",t,
//         );
//         test.BasicTest(i-1,v,"Pop did not return correct value.",t);
//     }
//     v,err:=tmp.PopBack();
//     test.BasicTest(nil,err,
//         "Pop returned an error when it should not have.",t,
//     );
//     test.BasicTest(20,v,"Pop did not return correct value.",t);
//     _,err=tmp.PopBack();
//     test.ContainsError(containerTypes.Empty,err,"Pop did not return correct error.",t)
// }
// 
// func TestCircularBufferClear(t *testing.T){
//     tmp,_:=NewCircularBuffer[int,widgets.BuiltinInt](5)
//     for i:=0; i<5; i++ {
//         tmp.PushBack(i)
//     }
//     tmp.Clear()
//     test.BasicTest(0,tmp.numElems,
//         "Clear did not reset num elems",t,
//     )
//     test.BasicTest(0,tmp.startEnd.A,
//         "Clear did not reset the start index.",t,
//     )
//     test.BasicTest(len(tmp.vals)-1,tmp.startEnd.B,
//         "Clear did not reset the end index.",t,
//     )
//     test.BasicTest(5,len(tmp.vals),
//         "Clear did not change the underlying slice.",t,
//     )
//     for i:=0; i<5; i++ {
//         test.BasicTest(0,tmp.vals[i],
//             "Clear did not clear the underlying slice.",t,
//         )
//     }
// }

// func testCircularBufferElemsHelper(
//     c SyncedCircularBuffer[int,widgets.BuiltinInt],
//     t *testing.T,
// ){
//     for i:=0; i<len(c.vals); i++ {
//         c.PushBack(i);
//     }
//     cnt:=0
//     c.Elems().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
//         cnt++
//         test.BasicTest(index,val,"Element was skipped while iterating.",t);
//         return iter.Continue,nil;
//     });
//     test.BasicTest(len(c.vals),cnt,
//         "All the elements were not iterated over.",t,
//     )
// }
// func TestCircularBufferElems(t *testing.T){
//     tmp,err:=NewSyncedCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     testCircularBufferElemsHelper(tmp,t);
//     tmp,err=NewSyncedCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     tmp.startEnd.A=2;
//     tmp.startEnd.B=1;
//     testCircularBufferElemsHelper(tmp,t);
// }
// 
// func testCircularBufferPntrElemsHelper(
//     c CircularBuffer[int,widgets.BuiltinInt], 
//     t *testing.T,
// ){
//     for i:=0; i<len(c.vals); i++ {
//         c.PushBack(i);
//     }
//     cnt:=0
//     c.PntrElems().ForEach(func(index int, val *int) (iter.IteratorFeedback, error) {
//         cnt++
//         test.BasicTest(index,*val,"Element was skipped while iterating.",t);
//         *val=100;
//         return iter.Continue,nil;
//     });
//     c.Elems().ForEach(func(index int, val int) (iter.IteratorFeedback, error) {
//         test.BasicTest(100,val,"Element was not updated while iterating.",t);
//         return iter.Continue,nil;
//     });
//     test.BasicTest(len(c.vals),cnt,
//         "All the elements were not iterated over.",t,
//     )
// }
// func TestCircularBufferPntrElems(t *testing.T){
//     tmp,err:=NewSyncedCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     testCircularBufferElemsHelper(tmp,t);
//     tmp,err=NewSyncedCircularBuffer[int,widgets.BuiltinInt](5);
//     test.BasicTest(nil,err,
//         "NewCircularBuffer returned an error when it should not have.",t,
//     );
//     tmp.startEnd.A=2;
//     tmp.startEnd.B=1;
//     testCircularBufferElemsHelper(tmp,t);
// }

// func TestCircularBufferEq(t *testing.T){
//     v,_:=NewCircularBuffer[int,widgets.BuiltinInt](4)
//     v2,_:=NewCircularBuffer[int,widgets.BuiltinInt](4)
//     comp:=func(l *int, r *int) bool { return *l==*r }
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
//     v.Push(0,1)
//     test.BasicTest(false,v.Eq(&v2,comp),
// 	"Eq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Eq(&v,comp),
// 	"Eq returned a false positive.",t,
//     )
//     v2.Push(0,1)
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
//     v.Push(1,2,3,4)
//     test.BasicTest(false,v.Eq(&v2,comp),
// 	"Eq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Eq(&v,comp),
// 	"Eq returned a false positive.",t,
//     )
//     v2.Push(1,2,3,4)
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
// }
// 
// func TestCircularBufferNeq(t *testing.T){
//     v,_:=NewCircularBuffer[int,widgets.BuiltinInt](4)
//     v2,_:=NewCircularBuffer[int,widgets.BuiltinInt](4)
//     comp:=func(l *int, r *int) bool { return *l==*r }
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
//     v.Push(0,1)
//     test.BasicTest(true,v.Neq(&v2,comp),
// 	"Neq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Neq(&v,comp),
// 	"Neq returned a false negative.",t,
//     )
//     v2.Push(0,1)
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
//     v.Push(1,2,3,4)
//     test.BasicTest(true,v.Neq(&v2,comp),
// 	"Neq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Neq(&v,comp),
// 	"Neq returned a false negative.",t,
//     )
//     v2.Push(1,2,3,4)
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
// }
