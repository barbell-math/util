package containers

import (
	"testing"

	// "github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/test"
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
    test.BasicTest(true,v1.Eq(&v1,&v2),
	"The widget equality check did not return true when it should have.",t,
    )
    test.BasicTest(true,v1.Eq(&v2,&v1),
	"The widget equality check did not return true when it should have.",t,
    )
    v1[0][0]="blah"
    test.BasicTest(false,v1.Eq(&v1,&v2),
	"The widget equality check did not return true when it should have.",t,
    )
    test.BasicTest(false,v1.Eq(&v2,&v1),
	"The widget equality check did not return true when it should have.",t,
    )
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
    test.BasicTest(false,v1.Lt(&v1,&v2),
	"The widget Lt method did not return false when comparing equal vectors.",t,
    )
    test.BasicTest(false,v1.Lt(&v2,&v1),
	"The widget Lt method did not return false when comparing equal vectors.",t,
    )
    v1[0][0]="A"
    test.BasicTest(true,v1.Lt(&v1,&v2),
	"The widget Lt method did not returned a false negative.",t,
    )
    test.BasicTest(false,v1.Lt(&v2,&v1),
	"The widget Lt method did not returned a false positive.",t,
    )
    v1[0][0]="a"
    v1[0][1]="B"
    test.BasicTest(true,v1.Lt(&v1,&v2),
	"The widget Lt method did not returned a false negative.",t,
    )
    test.BasicTest(false,v1.Lt(&v2,&v1),
	"The widget Lt method did not returned a false positive.",t,
    )
    v1[0][1]="b"
    v1.Delete(2)
    test.BasicTest(true,v1.Lt(&v1,&v2),
	"The widget Lt method did not returned a false negative.",t,
    )
    test.BasicTest(false,v1.Lt(&v2,&v1),
	"The widget Lt method did not returned a false positive.",t,
    )
    v2.Delete(1)
    v2.Delete(1)
    test.BasicTest(false,v1.Lt(&v1,&v2),
	"The widget Lt method did not returned a false positive.",t,
    )
    test.BasicTest(true,v1.Lt(&v2,&v1),
	"The widget Lt method did not returned a false negative.",t,
    )
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
    test.BasicTest(v1.Hash(&v1),v2.Hash(&v2),
	"The widget hash method did not return the same value for identical vectors.",t,
    )
    v1[0][0]="blah"
    test.BasicTest(false,v1.Hash(&v1)==v2.Hash(&v2),
	"The widget hash method did not return different values for different vectors.",t,
    )
}

func TestVectorZero(t *testing.T){
    v:=Vector[int,widgets.BuiltinInt]{1,2,3}
    v.Zero(&v)
    test.SlicesMatch[int]([]int{},v,t)
}
