package containers

import (
	"testing"

	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/container/containerTypes"
	"github.com/barbell-math/util/src/test"
	"github.com/barbell-math/util/src/widgets"
)

func TestMapKeyedUnionBothEmpty(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	MapKeyedUnion[int, int](&m1, &m2)

	test.Eq(m1.Length(), 0, t)
	test.Eq(m2.Length(), 0, t)
}

func TestMapKeyedUnionLeftEmpty(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)

	m2.Emplace(basic.Pair[int, int]{1, 2})

	MapKeyedUnion[int, int](&m1, &m2)

	keys, err := m1.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](keys, []int{1}, t)

	vals, err := m1.Vals().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](vals, []int{2}, t)

	test.True(m1.Eq(&m1, &m2), t)
}

func TestMapKeyedUnionRightEmpty(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)

	m1.Emplace(basic.Pair[int, int]{1, 2})

	MapKeyedUnion[int, int](&m1, &m2)

	keys, err := m1.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](keys, []int{1}, t)

	vals, err := m1.Vals().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](vals, []int{2}, t)
}

func TestMapKeyedUnionBothNonEmpty(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)

	m1.Emplace(basic.Pair[int, int]{1, 2})
	m2.Emplace(basic.Pair[int, int]{3, 4})

	MapKeyedUnion[int, int](&m1, &m2)

	keys, err := m1.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](keys, []int{1, 3}, t)

	vals, err := m1.Vals().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](vals, []int{2, 4}, t)
}

func TestMapKeyedUnionBothNonEmptyDupKeys(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)

	m1.Emplace(basic.Pair[int, int]{1, 2})
	m2.Emplace(basic.Pair[int, int]{1, 3})

	MapKeyedUnion[int, int](&m1, &m2)

	keys, err := m1.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](keys, []int{1}, t)

	vals, err := m1.Vals().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](vals, []int{3}, t)
}

func TestMapDisjointKeyedUnionBothEmpty(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	err := MapDisjointKeyedUnion[int, int](&m1, &m2)
	test.Nil(err, t)

	test.Eq(m1.Length(), 0, t)
	test.Eq(m2.Length(), 0, t)
}

func TestMapDisjointKeyedUnionLeftEmpty(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)

	m2.Emplace(basic.Pair[int, int]{1, 2})

	err := MapDisjointKeyedUnion[int, int](&m1, &m2)
	test.Nil(err, t)

	keys, err := m1.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](keys, []int{1}, t)

	vals, err := m1.Vals().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](vals, []int{2}, t)

	test.True(m1.Eq(&m1, &m2), t)
}

func TestMapDisjointKeyedUnionRightEmpty(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)

	m1.Emplace(basic.Pair[int, int]{1, 2})

	err := MapDisjointKeyedUnion[int, int](&m1, &m2)
	test.Nil(err, t)

	keys, err := m1.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](keys, []int{1}, t)

	vals, err := m1.Vals().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](vals, []int{2}, t)
}

func TestMapDisjointKeyedUnionBothNonEmpty(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)

	m1.Emplace(basic.Pair[int, int]{1, 2})
	m2.Emplace(basic.Pair[int, int]{3, 4})

	err := MapDisjointKeyedUnion[int, int](&m1, &m2)
	test.Nil(err, t)

	keys, err := m1.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](keys, []int{1, 3}, t)

	vals, err := m1.Vals().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[int](vals, []int{2, 4}, t)
}

func TestMapDisjointKeyedUnionBothNonEmptyDupKeys(t *testing.T) {
	m1, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)
	m2, _ := NewHashMap[int, int, widgets.BuiltinInt, widgets.BuiltinInt](0)

	m1.Emplace(basic.Pair[int, int]{1, 2})
	m2.Emplace(basic.Pair[int, int]{1, 3})

	err := MapDisjointKeyedUnion[int, int](&m1, &m2)
	test.ContainsError(containerTypes.Duplicate, err, t)
}
