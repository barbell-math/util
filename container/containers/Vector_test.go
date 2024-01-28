package containers

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/container/widgets"
	"github.com/barbell-math/util/test"
)

//go:generate go run interfaceTest.go -type=Vector -category=dynamic -interface=Vector -factory=generateVector -info
//go:generate go run interfaceTest.go -type=SyncedVector -category=dynamic -interface=Vector -factory=generateSyncedVector -info
// //go:generate interfaceTest -type=Vector -category=dynamic -interface=Queue -info
// //go:generate interfaceTest -type=Vector -category=dynamic -interface=Stack -info
// //go:generate interfaceTest -type=Vector -category=dynamic -interface=Deque -info

func generateVector() dynamicContainers.Vector[int,int] {
    rv,_:=NewVector[int,widgets.BuiltinInt](0)
    return &rv
}

func generateSyncedVector() dynamicContainers.Vector[int,int] {
    rv,_:=NewSyncedVector[int,widgets.BuiltinInt](0)
    return &rv
}

func TestVectorTypeCasting(t *testing.T){
    v,_:=NewVector[string,widgets.BuiltinString](3)
    s:=[]string(v)
    _=s

    s2:=make([]string,4)
    v2:=Vector[string,widgets.BuiltinString](s2)
    _=v2
}

func testVectorElemsHelper(
    v SyncedVector[int,widgets.BuiltinInt], 
    l int, 
    t *testing.T,
){
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
    tmp,err:=NewSyncedVector[int,widgets.BuiltinInt](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,0,t);
    tmp,err=NewSyncedVector[int,widgets.BuiltinInt](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,1,t);
    tmp,err=NewSyncedVector[int,widgets.BuiltinInt](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,2,t);
}

func testVectorPntrElemsHelper(
    v SyncedVector[int,widgets.BuiltinInt], 
    l int, 
    t *testing.T,
){
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
    tmp,err:=NewSyncedVector[int,widgets.BuiltinInt](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,0,t);
    tmp,err=NewSyncedVector[int,widgets.BuiltinInt](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,1,t);
    tmp,err=NewSyncedVector[int,widgets.BuiltinInt](0);
    test.BasicTest(nil,err,
        "NewCircularBuffer returned an error when it should not have.",t,
    );
    testVectorElemsHelper(tmp,2,t);
}

// func TestVectorEq(t *testing.T){
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
// func TestVectorNeq(t *testing.T){
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
