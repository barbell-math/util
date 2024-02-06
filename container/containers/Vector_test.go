package containers

import (
	"testing"

	// "github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	// "github.com/barbell-math/util/test"
)

//go:generate go run interfaceTest.go -type=Vector -category=dynamic -interface=Vector -genericDecl=[int] -factory=generateVector
//go:generate go run interfaceTest.go -type=SyncedVector -category=dynamic -interface=Vector -genericDecl=[int] -factory=generateSyncedVector
//go:generate go run interfaceTest.go -type=Vector -category=dynamic -interface=Deque -genericDecl=[int] -factory=generateVector
//go:generate go run interfaceTest.go -type=SyncedVector -category=dynamic -interface=Deque -genericDecl=[int] -factory=generateSyncedVector
//go:generate go run interfaceTest.go -type=Vector -category=dynamic -interface=Queue -genericDecl=[int] -factory=generateVector
//go:generate go run interfaceTest.go -type=SyncedVector -category=dynamic -interface=Queue -genericDecl=[int] -factory=generateSyncedVector
//go:generate go run interfaceTest.go -type=Vector -category=dynamic -interface=Stack -genericDecl=[int] -factory=generateVector
//go:generate go run interfaceTest.go -type=SyncedVector -category=dynamic -interface=Stack -genericDecl=[int] -factory=generateSyncedVector
//go:generate go run interfaceTest.go -type=Vector -category=dynamic -interface=Set -genericDecl=[int] -factory=generateVector
//go:generate go run interfaceTest.go -type=SyncedVector -category=dynamic -interface=Set -genericDecl=[int] -factory=generateSyncedVector

func generateVector() Vector[int,widgets.BuiltinInt] {
    v,_:=NewVector[int,widgets.BuiltinInt](0)
    return v
}

func generateSyncedVector() SyncedVector[int,widgets.BuiltinInt] {
    v,_:=NewSyncedVector[int,widgets.BuiltinInt](0)
    return v
}

func TestVectorTypeCasting(t *testing.T){
    v,_:=NewVector[string,widgets.BuiltinString](3)
    s:=[]string(v)
    _=s

    s2:=make([]string,4)
    v2:=Vector[string,widgets.BuiltinString](s2)
    _=v2
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
