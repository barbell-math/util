package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func vectorReadInterface[T ~int, U any](c dynamicContainers.ReadVector[T,U]){}
func vectorWriteInterface[T ~int, U any](c dynamicContainers.WriteVector[T,U]){}
func vectorInterface[T ~int, U any](c dynamicContainers.Vector[T,U]){}

func VectorInterfaceSyncableInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.RWSyncable=factory()
    _=container
}

func VectorInterfaceLengthInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.Length=factory()
    _=container
}

func VectorInterfaceCapacityInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.Capacity=factory()
    _=container
}

func VectorInterfaceWriteOpsInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.WriteOps[K,V]=factory()
    _=container
}

func VectorInterfaceWriteKeyedOpsInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.WriteKeyedOps[K,V]=factory()
    _=container
}

func VectorInterfaceReadOpsInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.ReadOps[K,V]=factory()
    _=container
}

func VectorInterfaceReadKeyedOpsInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.ReadKeyedOps[K,V]=factory()
    _=container
}

func VectorInterfaceDeleteOpsInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.DeleteOps[K,V]=factory()
    _=container
}

func VectorInterfaceDeleteKeyedOpsInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.DeleteKeyedOps[K,V]=factory()
    _=container
}

func VectorInterfaceFirstElemReadInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.FirstElemRead[V]=factory()
    _=container
}

func VectorInterfaceFirstElemWriteInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.FirstElemWrite[V]=factory()
    _=container
}

func VectorInterfaceFirstElemDeleteInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.FirstElemDelete[V]=factory()
    _=container
}

func VectorInterfaceLastElemReadInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.LastElemRead[V]=factory()
    _=container
}

func VectorInterfaceLastElemWriteInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.LastElemWrite[V]=factory()
    _=container
}

func VectorInterfaceLastElemDeleteInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    var container containerTypes.LastElemDelete[V]=factory()
    _=container
}

func ReadVectorInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    vectorReadInterface[K,V](factory());
}

func WriteVectorInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    vectorWriteInterface[K,V](factory());
}

func VectorInterfaceInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    vectorInterface[K,V](factory());
}

func VectorInterfaceStaticCapacityInterface[K ~int, V any](
    factory func() dynamicContainers.Vector[K,V], 
    t *testing.T,
){
    test.Panics(
        func () {
            var c any
            c=factory()
            c2:=c.(containerTypes.StaticCapacity)
            _=c2
        },
        "Code did not panic when casting a dynamic vector to a static vector.",t,
    )
}
// func VectorInterfaceStaticTypeInterface(t *testing.T){
//     test.Panics(
//         func () {
//             var c any
//             c,_=NewVector[int,widgets.BuiltinInt]((5)
//             c2:=c.(static.Vector[int,widgets.BuiltinInt]()
//             _=c2
//         },
//         "Code did not panic when casting a dynamic vector to a static vector.",t,
//     )
//     test.Panics(
//         func () {
//             var c any
//             c,_=NewVector[int,widgets.BuiltinInt]((5)
//             c2:=c.(static.Queue[int])
//             _=c2
//         },
//         "Code did not panic when casting a dynamic vector to a static queue.",t,
//     )
//     test.Panics(
//         func () {
//             var c any
//             c,_=NewVector[int,widgets.BuiltinInt]((5)
//             c2:=c.(static.Stack[int])
//             _=c2
//         },
//         "Code did not panic when casting a dynamic vector to a static stack.",t,
//     )
//     test.Panics(
//         func () {
//             var c any
//             c,_=NewVector[int,widgets.BuiltinInt]((5)
//             c2:=c.(static.Deque[int])
//             _=c2
//         },
//         "Code did not panic when casting a dynamic vector to a static deque.",t,
//     )
// }

func VectorInterfaceGet(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    _,err:=container.Get(0)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Get did not return the correct error with invalid index.",t,
    )
    for i:=0; i<5; i++ {
        container.Append(i)
    }
    for i:=0; i<5; i++ {
	_v,err:=container.Get(i)
	test.BasicTest(i,_v,
	    "Get did not return the correct value.",t,
	)
	test.BasicTest(nil,err,
	    "Get returned an error when it shouldn't have.",t,
	)
    }
    _,err=container.Get(-1)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Get did not return the correct error with invalid index.",t,
    )
    _,err=container.Get(6)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Get did not return the correct error with invalid index.",t,
    )
}

func VectorInterfaceGetPntr(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    _,err:=container.GetPntr(0)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Get did not return the correct error with invalid index.",t,
    )
    for i:=0; i<5; i++ {
        container.Append(i)
    }
    for i:=0; i<5; i++ {
	_v,err:=container.GetPntr(i)
	test.BasicTest(i,*_v,
	    "Get did not return the correct value.",t,
	)
	test.BasicTest(nil,err,
	    "Get returned an error when it shouldn't have.",t,
	)
    }
    _,err=container.GetPntr(-1)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Get pntr did not return the correct error with invalid index.",t,
    )
    _,err=container.GetPntr(6)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Get pntr did not return the correct error with invalid index.",t,
    )
}

func vectorContainsHelper(
    v dynamicContainers.Vector[int,int], 
    l int, 
    t *testing.T,
){
    for i:=0; i<l; i++ {
		v.Append(i)
    }
    for i:=0; i<l; i++ {
	test.BasicTest(true,v.Contains(i),
	    "Contains returned a false negative.",t,
	)
    }
    test.BasicTest(false,v.Contains(-1),
	"Contains returned a false positive.",t,
    )
    test.BasicTest(false,v.Contains(l),
	"Contains returned a false positive.",t,
    )
}
func VectorInterfaceContains(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    vectorContainsHelper(factory(),0,t)
    vectorContainsHelper(factory(),1,t)
    vectorContainsHelper(factory(),2,t)
    vectorContainsHelper(factory(),5,t)
}

func vectorKeyOfHelper(
    v dynamicContainers.Vector[int,int], 
    l int, 
    t *testing.T,
){
    for i:=0; i<l; i++ {
		v.Append(i)
    }
    for i:=0; i<l; i++ {
	k,found:=v.KeyOf(i)
	test.BasicTest(i,k,
	    "KeyOf did not return the correct index.",t,
	)
	test.BasicTest(true,found,
	    "KeyOf returned a false negative.",t,
	)
    }
    _,found:=v.KeyOf(-1)
    test.BasicTest(false,found,
	"KeyOf returned a false positive.",t,
    )
    _,found=v.KeyOf(-1)
    test.BasicTest(false,v.Contains(l),
	"KeyOf returned a false positive.",t,
    )
}
func VectorInterfaceKeyOf(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    vectorKeyOfHelper(factory(),0,t)
    vectorKeyOfHelper(factory(),1,t)
    vectorKeyOfHelper(factory(),2,t)
    vectorKeyOfHelper(factory(),5,t)
}

func VectorInterfaceEmplace(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    err:=container.Emplace(0,6)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Emplace did not return the correct error with invalid index.",t,
    )
    for i:=0; i<5; i++ {
        container.Append(i)
    }
    for i:=0; i<5; i++ {
	err:=container.Emplace(i,i+1)
	test.BasicTest(nil,err,
	    "Get returned an error when it shouldn't have.",t,
	)
    }
    for i:=0; i<5; i++ {
        iterV,_:=container.Get(i)
	test.BasicTest(i+1,iterV,
	    "Emplace did not set the value correctly.",t,
	)
    }
    err=container.Emplace(-1,6)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Emplace did not return the correct error with invalid index.",t,
    )
    err=container.Emplace(6,6)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Emplace did not return the correct error with invalid index.",t,
    )
}

func VectorInterfaceAppend(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    for i:=0; i<5; i++ { container.Append(i) }
    for i:=0; i<5; i++ {
        iterV,_:=container.Get(i)
	test.BasicTest(i,iterV,
	    "Append did not add the value correctly.",t,
	)
    }
    container.Append(5,6,7)
    for i:=0; i<8; i++ {
        iterV,_:=container.Get(i)
	test.BasicTest(i,iterV,
	    "Append did not add the value correctly.",t,
	)
    }
}

func vectorPushHelper(
    v dynamicContainers.Vector[int,int], 
    idx int, 
    l int, 
    t *testing.T,
){
    for i:=0; i<l-1; i++ { v.Append(i) }
    err:=v.Push(idx,l-1)
    test.BasicTest(nil,err,
        "Push returned an error when it shouldn't have.",t,
    )
    test.BasicTest(l,v.Length(),
        "Push did not increment the number of elements.",t,
    )
    for i:=0; i<v.Length(); i++ {
        var exp int
        v,_:=v.Get(i)
        if i<idx {
            exp=i
        } else if i==idx {
            exp=l-1
        } else {
            exp=i-1
        }
        test.BasicTest(exp,v,
            "Push did not put the value in the correct place.",t,
        )
    }
}
func VectorInterfacePush(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    for i:=2; i>=0; i-- { container.Push(0,i) }
    for i:=3; i<5; i++ { container.Push(container.Length(),i) }
    for i:=0; i<5; i++ {
        iterV,_:=container.Get(i)
	test.BasicTest(i,iterV,
	    "Push did not put the values in the correct place.",t,
	)
    }
    for i:=0; i<5; i++ {
	vectorPushHelper(container,i,5,t)
    }
}

func vectorPopHelper(
    factory func() dynamicContainers.Vector[int,int],
    l int, 
    num int, 
    t *testing.T,
){
    // fmt.Println("Permutation: l: ",l," num: ",num)
    container:=factory()
    for i:=0; i<l; i++ {
	if i%4==0 {
	    container.Append(-1)
	} else {
	    container.Append(i)
	}
    }
    // fmt.Println("Init:   ",v)
    n:=container.Pop(-1,num)
    exp:=factory()
    cntr:=0
    for i:=0; i<l; i++ {
	if i%4==0 {
	    if cntr<num {
	        cntr++
	        continue
	    } else {
	        exp.Append(-1)
	    }
	} else {
	    exp.Append(i)
	}
    }
    test.BasicTest(exp.Length(),container.Length(),
	"Pop did not remove value from the list correctly.",t,
    )
    test.BasicTest(cntr,n,
	"Pop did not pop the correct number of values.",t,
    )
    // fmt.Println("EXP:    ",exp)
    // fmt.Println("Final:  ",v)
    for i:=0; i<container.Length(); i++ {
        iterV,_:=container.Get(i)
	expIterV,_:=exp.Get(i)
	test.BasicTest(expIterV,iterV,
	    "Pop did not shift the values correctly.",t,
	)
    }
}
func VectorInterfacePop(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    for i:=0; i<13; i++ {
		for j:=0; j<13; j++ {
		    vectorPopHelper(factory,i,j,t)
		}
    }
}

func VectorInterfaceDelete(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    for i:=0; i<6; i++ {
        container.Append(i)
    }
    for i:=container.Length()-1; i>=0; i-- {
	container.Delete(i)
	test.BasicTest(i,container.Length(),"Delete removed to many values.",t)
	for j:=0; j<i; j++ {
            iterV,_:=container.Get(j)
	    test.BasicTest(j,iterV,"Delete changed the wrong value.",t)
	}
    }
    err:=container.Delete(0)
    test.ContainsError(customerr.ValOutsideRange,err,
	"Delete returned an incorrect error.",t,
    )
}

func VectorInterfaceClear(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    for i:=0; i<6; i++ {
        container.Append(i)
    }
    container.Clear()
    test.BasicTest(0,container.Length(),"Clear did not reset the underlying vector.",t)
    test.BasicTest(0,container.Capacity(),"Clear did not reset the underlying vector.",t)
}

func VectorInterfacePeekPntrFront(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    _v,err:=container.PeekPntrFront()
    test.BasicTest((*int)(nil),_v,
		"Peek pntr front did not return the correct value.",t,	
    )
    test.ContainsError(customerr.ValOutsideRange,err,
		"Peek pntr front returned an incorrect error.",t,
    )
    container.Append(1)
    _v,err=container.PeekPntrFront()
    test.BasicTest(1,*_v,
		"Peek pntr front did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
		"Peek pntr front returned an error when it shouldn't have.",t,
    )
    container.Append(2)
    _v,err=container.PeekPntrFront()
    test.BasicTest(1,*_v,
		"Peek pntr front did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
		"Peek pntr front returned an error when it shouldn't have.",t,
    )
}

func VectorInterfacePeekFront(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    _,err:=container.PeekFront()
    test.ContainsError(customerr.ValOutsideRange,err,
	"Peek front returned an incorrect error.",t,
    )
    container.Append(1)
    _v,err:=container.PeekFront()
    test.BasicTest(1,_v,
	"Peek front did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek front returned an error when it shouldn't have.",t,
    )
    container.Append(2)
    _v,err=container.PeekFront()
    test.BasicTest(1,_v,
	"Peek front did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek front returned an error when it shouldn't have.",t,
    )
}

func VectorInterfacePeekPntrBack(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    _v,err:=container.PeekPntrBack()
    test.BasicTest((*int)(nil),_v,
	"Peek pntr back did not return the correct value.",t,	
    )
    test.ContainsError(customerr.ValOutsideRange,err,
	"Peek pntr back returned an incorrect error.",t,
    )
    container.Append(1)
    _v,err=container.PeekPntrBack()
    test.BasicTest(1,*_v,
	"Peek pntr back did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek pntr back returned an error when it shouldn't have.",t,
    )
    container.Append(2)
    _v,err=container.PeekPntrBack()
    test.BasicTest(2,*_v,
	"Peek pntr back did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek pntr back returned an error when it shouldn't have.",t,
    )
}

func VectorInterfacePeekBack(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    _,err:=container.PeekBack()
    test.ContainsError(customerr.ValOutsideRange,err,
	"Peek back returned an incorrect error.",t,
    )
    container.Append(1)
    _v,err:=container.PeekBack()
    test.BasicTest(1,_v,
	"Peek back did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek back returned an error when it shouldn't have.",t,
    )
    container.Append(2)
    _v,err=container.PeekBack()
    test.BasicTest(2,_v,
	"Peek back did not return the correct value.",t,	
    )
    test.BasicTest(nil,err,
	"Peek back returned an error when it shouldn't have.",t,
    )
}

func VectorInterfacePopFront(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    for i:=0; i<4; i++ {
        container.Append(i)
    }
    for i:=0; i<4; i++ {
	f,err:=container.PopFront()
	test.BasicTest(i,f,
	    "Pop front returned the incorrect value.",t,
	)
	test.BasicTest(nil,err,
	    "Pop front returned an error when it shoudn't have.",t,
	)
    }
    _,err:=container.PopFront()
    test.ContainsError(containerTypes.Empty,err,
	"Pop front returned an incorrect error.",t,
    )
}

func VectorInterfacePopBack(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    for i:=0; i<4; i++ {
        container.Append(i)
    }
    for i:=3; i>=0; i-- {
	f,err:=container.PopBack()
	test.BasicTest(i,f,
	    "Pop front returned the incorrect value.",t,
	)
	test.BasicTest(nil,err,
	    "Pop front returned an error when it shoudn't have.",t,
	)
    }
    _,err:=container.PopBack()
    test.ContainsError(containerTypes.Empty,err,
	"Pop front returned an incorrect error.",t,
    )
}

func VectorInterfacePushFront(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    for i:=0; i<4; i++ {
		container.PushFront(i)
		test.BasicTest(i+1,container.Length(),
		    "Push front did not add the value correctly.",t,
		)
		for j:=0; j<i+1; j++ {
    	        iterV,_:=container.Get(j)
		    test.BasicTest(i-j,iterV,
			"Push front did not put the value in the correct place.",t,
		    )
		}
    }
}

func VectorInterfacePushBack(
    factory func() dynamicContainers.Vector[int,int], 
    t *testing.T,
){
    container:=factory()
    for i:=0; i<4; i++ {
		container.PushBack(i)
		test.BasicTest(i+1,container.Length(),
		    "Push back did not add the value correctly.",t,
		)
		for j:=0; j<i+1; j++ {
    	    iterV,_:=container.Get(j)
			test.BasicTest(j,iterV,
				"Push back did not put the value in the correct place.",t,
			)
		}
    }
}

//func VectorElemsHelper(
//    v Vector[int,int],
//    l int, 
//    t *testing.T,
//){
//    for i:=0; i<l; i++ {
//        v.PushBack(i);
//    }
//    cnt:=0
//    v.Elems().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
//        cnt++
//        test.BasicTest(index,val,"Element was skipped while iterating.",t);
//        return iter.Continue,nil;
//    });
//    test.BasicTest(l,cnt,
//        "All the elements were not iterated over.",t,
//    )
//}
//func VectorInterfaceElems(t *testing.T){
//    tmp,err:=NewSyncedVector[int,widgets.BuiltinInt](0);
//    test.BasicTest(nil,err,
//        "NewCircularBuffer returned an error when it should not have.",t,
//    );
//    VectorElemsHelper(tmp,0,t);
//    tmp,err=NewSyncedVector[int,widgets.BuiltinInt](0);
//    test.BasicTest(nil,err,
//        "NewCircularBuffer returned an error when it should not have.",t,
//    );
//    VectorElemsHelper(tmp,1,t);
//    tmp,err=NewSyncedVector[int,widgets.BuiltinInt](0);
//    test.BasicTest(nil,err,
//        "NewCircularBuffer returned an error when it should not have.",t,
//    );
//    VectorElemsHelper(tmp,2,t);
//}
//
//func VectorPntrElemsHelper(
//    v SyncedVector[int,widgets.BuiltinInt], 
//    l int, 
//    t *testing.T,
//){
//    for i:=0; i<l; i++ {
//        v.PushBack(i);
//    }
//    cnt:=0
//    v.PntrElems().ForEach(func(index int, val *int) (iter.IteratorFeedback, error) {
//        cnt++
//        test.BasicTest(index,*val,"Element was skipped while iterating.",t);
//        *val=100;
//        return iter.Continue,nil;
//    });
//    v.Elems().ForEach(func(index int, val int) (iter.IteratorFeedback, error) {
//        test.BasicTest(100,val,"Element was not updated while iterating.",t);
//        return iter.Continue,nil;
//    });
//    test.BasicTest(l,cnt,
//        "All the elements were not iterated over.",t,
//    )
//}
//func VectorInterfaceElemPntrs(t *testing.T){
//    tmp,err:=NewSyncedVector[int,widgets.BuiltinInt](0);
//    test.BasicTest(nil,err,
//        "NewCircularBuffer returned an error when it should not have.",t,
//    );
//    VectorElemsHelper(tmp,0,t);
//    tmp,err=NewSyncedVector[int,widgets.BuiltinInt](0);
//    test.BasicTest(nil,err,
//        "NewCircularBuffer returned an error when it should not have.",t,
//    );
//    VectorElemsHelper(tmp,1,t);
//    tmp,err=NewSyncedVector[int,widgets.BuiltinInt](0);
//    test.BasicTest(nil,err,
//        "NewCircularBuffer returned an error when it should not have.",t,
//    );
//    VectorElemsHelper(tmp,2,t);
//}

// func VectorInterfaceEq(t *testing.T){
//     v:=Vector[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     v2:=Vector[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     comp:=func(l *int, r *int) bool { return *l==*r }
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
//     v.Delete(3)
//     test.BasicTest(false,v.Eq(&v2,comp),
// 	"Eq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Eq(&v,comp),
// 	"Eq returned a false positive.",t,
//     )
//     v=Vector[int,widgets.BuiltinInt]([]int{0})
//     v2=Vector[int,widgets.BuiltinInt]([]int{0})
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
//     v.Delete(0)
//     test.BasicTest(false,v.Eq(&v2,comp),
// 	"Eq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Eq(&v,comp),
// 	"Eq returned a false positive.",t,
//     )
//     v=Vector[int,widgets.BuiltinInt]([]int{})
//     v2=Vector[int,widgets.BuiltinInt]([]int{})
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
// }
// 
// func VectorInterfaceNeq(t *testing.T){
//     v:=Vector[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     v2:=Vector[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     comp:=func(l *int, r *int) bool { return *l==*r }
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
//     v.Delete(3)
//     test.BasicTest(true,v.Neq(&v2,comp),
// 	"Neq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Neq(&v,comp),
// 	"Neq returned a false negative.",t,
//     )
//     v=Vector[int,widgets.BuiltinInt]([]int{0})
//     v2=Vector[int,widgets.BuiltinInt]([]int{0})
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
//     v.Delete(0)
//     test.BasicTest(true,v.Neq(&v2,comp),
// 	"Neq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Neq(&v,comp),
// 	"Neq returned a false negative.",t,
//     )
//     v=Vector[int,widgets.BuiltinInt]([]int{})
//     v2=Vector[int,widgets.BuiltinInt]([]int{})
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
// }

