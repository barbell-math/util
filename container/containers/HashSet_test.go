package containers

import (
	"testing"

	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/test"
)

//go:generate go run interfaceTest.go -type=HashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateHashSet
//go:generate go run interfaceTest.go -type=SyncedHashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateSyncedHashSet

func generateHashSet() HashSet[int,badBuiltinInt] {
    v,_:=NewHashSet[int,badBuiltinInt](0)
    return v
}

func generateSyncedHashSet() SyncedHashSet[int,badBuiltinInt] {
    v,_:=NewSyncedHashSet[int,badBuiltinInt](0)
    return v
}

func TestHashSetEquality(t *testing.T) {
    s1,_:=NewHashSet[int,widgets.BuiltinInt](0)
    s1.AppendUnique(0,1,2,3,4)
    s2,_:=NewHashSet[int,widgets.BuiltinInt](0)
    s2.AppendUnique(0,1,2,3,4)
    test.True(s1.Eq(&s1,&s2),t)
    s2.AppendUnique(5)
    test.False(s1.Eq(&s1,&s2),t)
}

func TestHashSetLt(t *testing.T) {
    test.Panics(
        func() {
            s1,_:=NewHashSet[int,widgets.BuiltinInt](0)
            s1.AppendUnique(0,1,2,3,4)
            s2,_:=NewHashSet[int,widgets.BuiltinInt](0)
            s2.AppendUnique(0,1,2,3,4)
            s1.Lt(&s1,&s2)
        },
        t,
    ) 
}

func TestHashSetHash(t *testing.T) {
    s1,_:=NewHashSet[int,widgets.BuiltinInt](0)
    s1.AppendUnique(0,1,2,3,4)
    s2,_:=NewHashSet[int,widgets.BuiltinInt](0)
    s2.AppendUnique(0,1,2,3,4)
    test.Eq(s1.Hash(&s1),s2.Hash(&s2),t)
    s2.AppendUnique(5)
    test.False(s1.Hash(&s1)==s2.Hash(&s2),t)
    for i:=5; i<100; i++ {
        s2.AppendUnique(i*100)
    }
    h:=s1.Hash(&s1)
    for i:=0; i<100; i++ {
	test.Eq(h,s1.Hash(&s1),t)
    }
}

func TestHashSetZero(t *testing.T) {
    s1,_:=NewHashSet[int,widgets.BuiltinInt](0)
    s1.AppendUnique(0,1,2,3,4)
    s1.Zero(&s1)
    test.Eq(0,s1.Length(),t)
}
