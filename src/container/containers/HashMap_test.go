package containers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/test"
	"github.com/barbell-math/util/src/widgets"
)

//go:generate ../../../bin/containerInterfaceTests -type=HashMap -category=dynamic -interface=Map -genericDecl=[int,int] -factory=generateHashMap
//go:generate ../../../bin/containerInterfaceTests -type=SyncedHashMap -category=dynamic -interface=Map -genericDecl=[int,int] -factory=generateSyncedHashMap

func generateHashMap(capacity int) HashMap[int, int, badBuiltinInt, widgets.BuiltinInt] {
	m, _ := NewHashMap[int, int, badBuiltinInt, widgets.BuiltinInt](capacity)
	return m
}

func generateSyncedHashMap(capacity int) SyncedHashMap[
	int,
	int,
	badBuiltinInt,
	widgets.BuiltinInt,
] {
	m, _ := NewSyncedHashMap[int, int, badBuiltinInt, widgets.BuiltinInt](capacity)
	return m
}

func TestHashMapWidgetInterface(t *testing.T) {
	var widget widgets.BaseInterface[HashMap[string, string, widgets.BuiltinString, widgets.BuiltinString]]
	v, _ := NewHashMap[string, string, widgets.BuiltinString, widgets.BuiltinString](0)
	widget = &v
	_ = widget
}

func TestHashMapEq(t *testing.T) {
	m1, _ := NewHashMap[int, string, widgets.BuiltinInt, widgets.BuiltinString](0)
	m2, _ := NewHashMap[int, string, widgets.BuiltinInt, widgets.BuiltinString](0)
	m1.Emplace(
		basic.Pair[int, string]{0, "zero"},
		basic.Pair[int, string]{1, "one"},
		basic.Pair[int, string]{2, "two"},
		basic.Pair[int, string]{3, "three"},
	)
	m2.Emplace(
		basic.Pair[int, string]{0, "zero"},
		basic.Pair[int, string]{1, "one"},
		basic.Pair[int, string]{2, "two"},
		basic.Pair[int, string]{3, "three"},
	)
	test.True(m1.Eq(&m1, &m2), t)
	test.True(m2.Eq(&m1, &m2), t)
	m2.Delete(0)
	test.False(m1.Eq(&m1, &m2), t)
	test.False(m2.Eq(&m1, &m2), t)
}

func TestHashMapHash(t *testing.T) {
	m1, _ := NewHashMap[int, string, widgets.BuiltinInt, widgets.BuiltinString](0)
	m2, _ := NewHashMap[int, string, widgets.BuiltinInt, widgets.BuiltinString](0)
	m1.Emplace(
		basic.Pair[int, string]{0, "zero"},
		basic.Pair[int, string]{1, "one"},
		basic.Pair[int, string]{2, "two"},
		basic.Pair[int, string]{3, "three"},
	)
	m2.Emplace(
		basic.Pair[int, string]{0, "zero"},
		basic.Pair[int, string]{1, "one"},
		basic.Pair[int, string]{2, "two"},
		basic.Pair[int, string]{3, "three"},
	)
	test.Eq(m1.Hash(&m1), m2.Hash(&m2), t)
	m2.Set(basic.Pair[int, string]{0, "nil"})
	test.Neq(m1.Hash(&m1), m2.Hash(&m2), t)
	h := m1.Hash(&m1)
	for i := 0; i < 100; i++ {
		test.Eq(h, m1.Hash(&m1), t)
	}
}

func TestHashMapZero(t *testing.T) {
	m1, _ := NewHashMap[int, string, widgets.BuiltinInt, widgets.BuiltinString](0)
	m1.Emplace(
		basic.Pair[int, string]{0, "zero"},
		basic.Pair[int, string]{1, "one"},
		basic.Pair[int, string]{2, "two"},
		basic.Pair[int, string]{3, "three"},
	)
	m1.Clear()
	test.Eq(0, m1.Length(), t)
	test.Eq(0, len(m1.internalHashMapImpl), t)
}

func TestHashMapFormat(t *testing.T) {
	m1, _ := NewHashMap[int, string, widgets.BuiltinInt, widgets.BuiltinString](0)
	m1.Emplace(
		basic.Pair[int, string]{0, "zero"},
		basic.Pair[int, string]{1, "one"},
		basic.Pair[int, string]{2, "two"},
		basic.Pair[int, string]{3, "three"},
	)
	res := fmt.Sprintf("%v", m1)
	test.Eq(len("hashMap[0:zero 1:one 2:two 3:three]"), len(res), t)
	test.True(strings.Contains(res, "0:zero"), t)
	test.True(strings.Contains(res, "1:one"), t)
	test.True(strings.Contains(res, "2:two"), t)
	test.True(strings.Contains(res, "3:three"), t)

	res = fmt.Sprintf("%v", &m1)
	test.Eq(len("hashMap[0:zero 1:one 2:two 3:three]"), len(res), t)
	test.True(strings.Contains(res, "0:zero"), t)
	test.True(strings.Contains(res, "1:one"), t)
	test.True(strings.Contains(res, "2:two"), t)
	test.True(strings.Contains(res, "3:three"), t)
}

func TestHashMapString(t *testing.T) {
	m1, _ := NewHashMap[int, string, widgets.BuiltinInt, widgets.BuiltinString](0)
	m1.Emplace(
		basic.Pair[int, string]{0, "zero"},
		basic.Pair[int, string]{1, "one"},
		basic.Pair[int, string]{2, "two"},
		basic.Pair[int, string]{3, "three"},
	)
	res := m1.String()
	test.Eq(len("hashMap[0:zero 1:one 2:two 3:three]"), len(res), t)
	test.True(strings.Contains(res, "0:zero"), t)
	test.True(strings.Contains(res, "1:one"), t)
	test.True(strings.Contains(res, "2:two"), t)
	test.True(strings.Contains(res, "3:three"), t)
}
