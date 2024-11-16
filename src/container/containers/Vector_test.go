package containers

import (
	"testing"

	"github.com/barbell-math/util/src/test"
	"github.com/barbell-math/util/src/widgets"
)

//go:generate ../../../bin/containerInterfaceTests -type=Vector -category=dynamic -interface=Vector -genericDecl=[int] -factory=generateVector
//go:generate ../../../bin/containerInterfaceTests -type=SyncedVector -category=dynamic -interface=Vector -genericDecl=[int] -factory=generateSyncedVector
//go:generate ../../../bin/containerInterfaceTests -type=Vector -category=dynamic -interface=Deque -genericDecl=[int] -factory=generateVector
//go:generate ../../../bin/containerInterfaceTests -type=SyncedVector -category=dynamic -interface=Deque -genericDecl=[int] -factory=generateSyncedVector
//go:generate ../../../bin/containerInterfaceTests -type=Vector -category=dynamic -interface=Queue -genericDecl=[int] -factory=generateVector
//go:generate ../../../bin/containerInterfaceTests -type=SyncedVector -category=dynamic -interface=Queue -genericDecl=[int] -factory=generateSyncedVector
//go:generate ../../../bin/containerInterfaceTests -type=Vector -category=dynamic -interface=Stack -genericDecl=[int] -factory=generateVector
//go:generate ../../../bin/containerInterfaceTests -type=SyncedVector -category=dynamic -interface=Stack -genericDecl=[int] -factory=generateSyncedVector
//go:generate ../../../bin/containerInterfaceTests -type=Vector -category=dynamic -interface=Set -genericDecl=[int] -factory=generateVector
//go:generate ../../../bin/containerInterfaceTests -type=SyncedVector -category=dynamic -interface=Set -genericDecl=[int] -factory=generateSyncedVector

func generateVector(capacity int) Vector[int, widgets.BuiltinInt] {
	v, _ := NewVector[int, widgets.BuiltinInt](capacity)
	return v
}

func generateSyncedVector(capacity int) SyncedVector[int, widgets.BuiltinInt] {
	v, _ := NewSyncedVector[int, widgets.BuiltinInt](capacity)
	return v
}

func ExampleVector_typeCasting() {
	// Vectors can be type casted back and forth to regular slices. Note that
	// when you do this you loose any type information provided by the widget
	// interface.
	v, _ := NewVector[string, widgets.BuiltinString](3)
	s := []string(v)
	_ = s

	s2 := make([]string, 4)
	v2 := Vector[string, widgets.BuiltinString](s2)
	_ = v2
}

func TestVectorWidgetInterface(t *testing.T) {
	var widget widgets.BaseInterface[Vector[string, widgets.BuiltinString]]
	v, _ := NewVector[string, widgets.BuiltinString](0)
	widget = &v
	_ = widget
}

func TestVectorOfVectorsEquality(t *testing.T) {
	v1 := Vector[
		Vector[string, widgets.BuiltinString],
		*Vector[string, widgets.BuiltinString],
	]{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"h", "i", "j"},
	}
	v2 := Vector[
		Vector[string, widgets.BuiltinString],
		*Vector[string, widgets.BuiltinString],
	]{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"h", "i", "j"},
	}
	test.True(v1.Eq(&v1, &v2), t)
	test.True(v1.Eq(&v2, &v1), t)
	v1[0][0] = "blah"
	test.False(v1.Eq(&v1, &v2), t)
	test.False(v1.Eq(&v2, &v1), t)
}

func TestVectorOfVectorsHash(t *testing.T) {
	v1 := Vector[
		Vector[string, widgets.BuiltinString],
		*Vector[string, widgets.BuiltinString],
	]{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"h", "i", "j"},
	}
	v2 := Vector[
		Vector[string, widgets.BuiltinString],
		*Vector[string, widgets.BuiltinString],
	]{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"h", "i", "j"},
	}
	test.Eq(v1.Hash(&v1), v2.Hash(&v2), t)
	v1[0][0] = "blah"
	test.False(v1.Hash(&v1) == v2.Hash(&v2), t)
	h := v1.Hash(&v1)
	for i := 0; i < 100; i++ {
		test.Eq(h, v1.Hash(&v1), t)
	}
	v3 := Vector[int, widgets.BuiltinInt]{500, 600, 700}
	v4 := Vector[int, widgets.BuiltinInt]{700, 600, 500}
	test.False(v3.Hash(&v3) == v4.Hash(&v4), t)
}

func TestVectorZero(t *testing.T) {
	v := Vector[int, widgets.BuiltinInt]{1, 2, 3}
	v.Zero(&v)
	test.SlicesMatch[int]([]int{}, v, t)
}
