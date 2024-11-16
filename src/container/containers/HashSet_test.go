package containers

import (
	"testing"

	"github.com/barbell-math/util/src/iter"
	"github.com/barbell-math/util/src/test"
	"github.com/barbell-math/util/src/widgets"
)

//go:generate ../../../bin/containerInterfaceTests -type=HashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateHashSet
//go:generate ../../../bin/containerInterfaceTests -type=SyncedHashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateSyncedHashSet

func generateHashSet(capacity int) HashSet[int, badBuiltinInt] {
	v, _ := NewHashSet[int, badBuiltinInt](capacity)
	return v
}

func generateSyncedHashSet(capacity int) SyncedHashSet[int, badBuiltinInt] {
	v, _ := NewSyncedHashSet[int, badBuiltinInt](capacity)
	return v
}

func TestHashSetWidgetInterface(t *testing.T) {
	var widget widgets.BaseInterface[HashSet[string, widgets.BuiltinString]]
	v, _ := NewHashSet[string, widgets.BuiltinString](0)
	widget = &v
	_ = widget
}

func TestHashSetEquality(t *testing.T) {
	s1, _ := NewHashSet[int, widgets.BuiltinInt](0)
	s1.AppendUnique(0, 1, 2, 3, 4)
	s2, _ := NewHashSet[int, widgets.BuiltinInt](0)
	s2.AppendUnique(0, 1, 2, 3, 4)
	test.True(s1.Eq(&s1, &s2), t)
	s2.AppendUnique(5)
	test.False(s1.Eq(&s1, &s2), t)
}

func TestHashSetHash(t *testing.T) {
	s1, _ := NewHashSet[int, widgets.BuiltinInt](0)
	s1.AppendUnique(0, 1, 2, 3, 4)
	s2, _ := NewHashSet[int, widgets.BuiltinInt](0)
	s2.AppendUnique(0, 1, 2, 3, 4)
	test.Eq(s1.Hash(&s1), s2.Hash(&s2), t)
	s2.AppendUnique(5)
	test.Neq(s1.Hash(&s1), s2.Hash(&s2), t)
	h := s1.Hash(&s1)
	for i := 0; i < 100; i++ {
		test.Eq(h, s1.Hash(&s1), t)
	}
}

func TestHashSetZero(t *testing.T) {
	s1, _ := NewHashSet[int, widgets.BuiltinInt](0)
	s1.AppendUnique(0, 1, 2, 3, 4)
	s1.Zero(&s1)
	test.Eq(0, s1.Length(), t)
}

func popAndGetAffectedHashesHelper(
	setVals iter.Iter[int],
	popVal int,
	popHash HashSetHash,
	initialMap map[HashSetHash]int,
	res map[OldHashSetHash]NewHashSetHash,
	t *testing.T,
) {
	s1, _ := NewHashSet[int, badBuiltinInt2](0)
	setVals.ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		s1.AppendUnique(val)
		return iter.Continue, nil
	})
	test.MapsMatch[HashSetHash, int](initialMap, s1.internalHashSetImpl, t)
	deletedHash, vals, num := s1.popAndGetAffectedHashes(&popVal)
	test.Eq(popHash, deletedHash, t)
	test.Eq(1, num, t)
	s1.Pop(popVal)
	test.MapsMatch[OldHashSetHash, NewHashSetHash](res, vals, t)
}
func TestGetHashesAffectedByPop(t *testing.T) {
	initialMap := map[HashSetHash]int{}
	for i := 0; i < 8; i++ {
		initialMap[HashSetHash(i)] = i
	}
	for i := 0; i < 4; i++ {
		popAndGetAffectedHashesHelper(
			iter.Range[int](0, 8, 1),
			i,
			HashSetHash(i),
			initialMap,
			map[OldHashSetHash]NewHashSetHash{
				4: NewHashSetHash(i),
				5: 4,
				6: 5,
				7: 6,
			},
			t,
		)
	}
	for i := 4; i < 8; i++ {
		expRes := map[OldHashSetHash]NewHashSetHash{}
		for j := i; j < 7; j++ {
			expRes[OldHashSetHash(j+1)] = NewHashSetHash(j)
		}
		popAndGetAffectedHashesHelper(
			iter.Range[int](0, 8, 1),
			i,
			HashSetHash(i),
			initialMap,
			expRes,
			t,
		)
	}
	popAndGetAffectedHashesHelper(
		iter.SliceElems[int]([]int{0, 2, 3, 4, 5, 6, 7}),
		4,
		1,
		map[HashSetHash]int{
			0: 0,
			1: 4,
			2: 2,
			3: 3,
			4: 5,
			5: 6,
			6: 7,
		},
		map[OldHashSetHash]NewHashSetHash{
			4: 1,
			5: 4,
			6: 5,
		},
		t,
	)
	popAndGetAffectedHashesHelper(
		iter.SliceElems[int]([]int{0, 1, 2, 5, 6, 7}),
		0,
		0,
		map[HashSetHash]int{
			0: 0,
			1: 1,
			2: 2,
			3: 5,
			4: 6,
			5: 7,
		},
		map[OldHashSetHash]NewHashSetHash{},
		t,
	)
	popAndGetAffectedHashesHelper(
		iter.SliceElems[int]([]int{0, 4, 8, 7, 11}),
		0,
		0,
		map[HashSetHash]int{
			0: 0,
			1: 4,
			2: 8,
			3: 7,
			4: 11,
		},
		map[OldHashSetHash]NewHashSetHash{
			1: 0,
			2: 1,
		},
		t,
	)
}
