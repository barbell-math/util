package containers

import (
	"testing"

	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/test"
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

func TestWidgetInterface(t *testing.T){
    var widget widgets.WidgetInterface[Vector[string,widgets.BuiltinString]]
    v,_:=NewVector[string,widgets.BuiltinString](0)
    widget=&v
    _=widget
}

func TestVectorOfVectorsEquality(t *testing.T){
    v1:=Vector[
	Vector[string,widgets.BuiltinString],
	*Vector[string,widgets.BuiltinString],
    ]{
	{"a","b","c"},
	{"d","e","f"},
	{"h","i","j"},
    }
    v2:=Vector[
	Vector[string,widgets.BuiltinString],
	*Vector[string,widgets.BuiltinString],
    ]{
	{"a","b","c"},
	{"d","e","f"},
	{"h","i","j"},
    }
    test.True(v1.Eq(&v1,&v2),t)
    test.True(v1.Eq(&v2,&v1),t)
    v1[0][0]="blah"
    test.False(v1.Eq(&v1,&v2),t)
    test.False(v1.Eq(&v2,&v1),t)
}

func TestVectorOfVectorsLt(t *testing.T){
    v1:=Vector[
	Vector[string,widgets.BuiltinString],
	*Vector[string,widgets.BuiltinString],
    ]{
	{"a","b","c"},
	{"d","e","f"},
	{"h","i","j"},
    }
    v2:=Vector[
	Vector[string,widgets.BuiltinString],
	*Vector[string,widgets.BuiltinString],
    ]{
	{"a","b","c"},
	{"d","e","f"},
	{"h","i","j"},
    }
    test.False(v1.Lt(&v1,&v2),t)
    test.False(v1.Lt(&v2,&v1),t)
    v1[0][0]="A"
    test.True(v1.Lt(&v1,&v2),t)
    test.False(v1.Lt(&v2,&v1),t)
    v1[0][0]="a"
    v1[0][1]="B"
    test.True(v1.Lt(&v1,&v2),t)
    test.False(v1.Lt(&v2,&v1),t)
    v1[0][1]="b"
    v1.Delete(2)
    test.True(v1.Lt(&v1,&v2),t)
    test.False(v1.Lt(&v2,&v1),t)
    v2.Delete(1)
    v2.Delete(1)
    test.False(v1.Lt(&v1,&v2),t)
    test.True(v1.Lt(&v2,&v1),t)
}

func TestVectorOfVectorsHash(t *testing.T){
    v1:=Vector[
	Vector[string,widgets.BuiltinString],
	*Vector[string,widgets.BuiltinString],
    ]{
	{"a","b","c"},
	{"d","e","f"},
	{"h","i","j"},
    }
    v2:=Vector[
	Vector[string,widgets.BuiltinString],
	*Vector[string,widgets.BuiltinString],
    ]{
	{"a","b","c"},
	{"d","e","f"},
	{"h","i","j"},
    }
    test.Eq(v1.Hash(&v1),v2.Hash(&v2),t)
    v1[0][0]="blah"
    test.False(v1.Hash(&v1)==v2.Hash(&v2),t)
    h:=v1.Hash(&v1)
    for i:=0; i<100; i++ {
	test.Eq(h,v1.Hash(&v1),t)
    }
    v3:=Vector[int,widgets.BuiltinInt]{500,600,700}
    v4:=Vector[int,widgets.BuiltinInt]{700,600,500}
    test.False(v3.Hash(&v3)==v4.Hash(&v4),t)
}

func TestVectorZero(t *testing.T){
    v:=Vector[int,widgets.BuiltinInt]{1,2,3}
    v.Zero(&v)
    test.SlicesMatch[int]([]int{},v,t)
}
