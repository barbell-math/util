package containers

import (
	"testing"

	"github.com/barbell-math/util/src/iter"
	"github.com/barbell-math/util/src/test"
	"github.com/barbell-math/util/src/widgets"
)

//go:generate ../../../bin/containerInterfaceTests -type=HookedHashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateHookedHashSet
//go:generate ../../../bin/containerInterfaceTests -type=SyncedHookedHashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateSyncedHookedHashSet

type testSetHooks struct {
	addOpCntr int
	delOpCntr int
	clrOpCntr int
	addTestOp func(hashLoc HashSetHash)
	delTestOp func(
		deletedHash HashSetHash,
		updatedHashes map[OldHashSetHash]NewHashSetHash,
	)
}

func (t *testSetHooks) addOp(hashLoc HashSetHash) {
	t.addOpCntr++
	if t.addTestOp != nil {
		t.addTestOp(hashLoc)
	}
}
func (t *testSetHooks) deleteOp(
	deletedHash HashSetHash,
	updatedHashes map[OldHashSetHash]NewHashSetHash,
) {
	t.delOpCntr++
	if t.delTestOp != nil {
		t.delTestOp(deletedHash, updatedHashes)
	}
}
func (t *testSetHooks) clearOp() {
	t.clrOpCntr++
}

func generateHookedHashSet(capacity int) HookedHashSet[int, badBuiltinInt] {
	v, _ := NewHookedHashSet[int, badBuiltinInt](&testSetHooks{}, capacity)
	return v
}

func generateSyncedHookedHashSet(
	capacity int,
) SyncedHookedHashSet[int, badBuiltinInt] {
	v, _ := NewSyncedHookedHashSet[int, badBuiltinInt](&testSetHooks{}, capacity)
	return v
}

func TestHookedHashSetWidgetInterface(t *testing.T) {
	var widget widgets.BaseInterface[HookedHashSet[
		string, widgets.BuiltinString,
	]]

	v, _ := NewHookedHashSet[string, widgets.BuiltinString](&testSetHooks{}, 0)
	widget = &v
	_ = widget
}

func TestHookedHashSetHashOperation(t *testing.T) {
	s, _ := NewHookedHashSet[int, badBuiltinInt2](&testSetHooks{}, 0)
	for i := 0; i < 16; i++ {
		s.AppendUnique(i)
	}
	for i := 0; i < 16; i++ {
		res, ok := s.getHashPosition(&i)
		test.True(ok, t)
		test.Eq(HashSetHash(i), res, t)

		res1, err := s.GetFromHash(HashSetHash(i))
		test.Nil(err, t)
		test.Eq(i, res1, t)
	}
}

func TestHookedHashSetAppendUniqueHook(t *testing.T) {
	var valueToAdd int
	var s HookedHashSet[int, badBuiltinInt2]
	testHooks := testSetHooks{
		addTestOp: func(hashLoc HashSetHash) {
			res, err := s.GetFromHash(hashLoc)
			test.Nil(err, t)
			test.Eq(valueToAdd, res, t)
			test.Eq(valueToAdd, s.internalHashSetImpl[hashLoc], t)
		},
	}
	s, _ = NewHookedHashSet[int, badBuiltinInt2](&testHooks, 0)
	for i := 0; i < 16; i++ {
		valueToAdd = i
		s.AppendUnique(i)
	}
	test.Eq(16, testHooks.addOpCntr, t)
}

func hookedHashSetPopHookHelper(
	setVals iter.Iter[int],
	popVal int,
	popHash HashSetHash,
	initialMap map[HashSetHash]int,
	res map[OldHashSetHash]NewHashSetHash,
	t *testing.T,
) {
	var s HookedHashSet[int, badBuiltinInt2]
	testHooks := testSetHooks{
		delTestOp: func(
			deletedHash HashSetHash,
			updatedHashes map[OldHashSetHash]NewHashSetHash,
		) {
			test.Eq(popHash, deletedHash, t)
			test.MapsMatch[OldHashSetHash, NewHashSetHash](
				res,
				updatedHashes,
				t,
			)
		},
	}
	s, _ = NewHookedHashSet[int, badBuiltinInt2](&testHooks, 0)
	setVals.ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		s.AppendUnique(val)
		return iter.Continue, nil
	})
	test.MapsMatch[HashSetHash, int](initialMap, s.internalHashSetImpl, t)
	s.Pop(popVal)
	test.Eq(1, testHooks.delOpCntr, t)
}
func TestHookedHashSetPopHook(t *testing.T) {
	initialMap := map[HashSetHash]int{}
	for i := 0; i < 8; i++ {
		initialMap[HashSetHash(i)] = i
	}
	for i := 0; i < 4; i++ {
		hookedHashSetPopHookHelper(
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
		hookedHashSetPopHookHelper(
			iter.Range[int](0, 8, 1),
			i,
			HashSetHash(i),
			initialMap,
			expRes,
			t,
		)
	}
}

func hookedHashSetPopPntrHookHelper(
	setVals iter.Iter[int],
	popVal int,
	popHash HashSetHash,
	initialMap map[HashSetHash]int,
	res map[OldHashSetHash]NewHashSetHash,
	t *testing.T,
) {
	var s HookedHashSet[int, badBuiltinInt2]
	testHooks := testSetHooks{
		delTestOp: func(
			deletedHash HashSetHash,
			updatedHashes map[OldHashSetHash]NewHashSetHash,
		) {
			test.Eq(popHash, deletedHash, t)
			test.MapsMatch[OldHashSetHash, NewHashSetHash](
				res,
				updatedHashes,
				t,
			)
		},
	}
	s, _ = NewHookedHashSet[int, badBuiltinInt2](&testHooks, 0)
	setVals.ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		s.AppendUnique(val)
		return iter.Continue, nil
	})
	test.MapsMatch[HashSetHash, int](initialMap, s.internalHashSetImpl, t)
	s.PopPntr(&popVal)
	test.Eq(1, testHooks.delOpCntr, t)
}
func TestHookedHashSetPopPntrHook(t *testing.T) {
	initialMap := map[HashSetHash]int{}
	for i := 0; i < 8; i++ {
		initialMap[HashSetHash(i)] = i
	}
	for i := 0; i < 4; i++ {
		hookedHashSetPopPntrHookHelper(
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
		hookedHashSetPopPntrHookHelper(
			iter.Range[int](0, 8, 1),
			i,
			HashSetHash(i),
			initialMap,
			expRes,
			t,
		)
	}
}

func TestHookedHashSetClearHook(t *testing.T) {
	var s HookedHashSet[int, badBuiltinInt2]
	testHooks := testSetHooks{}
	s, _ = NewHookedHashSet[int, badBuiltinInt2](&testHooks, 0)
	for i := 0; i < 8; i++ {
		s.AppendUnique(i)
	}
	s.Clear()
	test.Eq(1, testHooks.clrOpCntr, t)
}
