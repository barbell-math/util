package iter

import (
	"testing"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/test"
)

// All of these tests could technically go in the producer test file but there
// has been difficulty getting the channels to work so these tests are kept
// separate.

func TestMapElemsStopEarly(t *testing.T) {
	f := func() basic.Pair[string, int] { return basic.Pair[string, int]{} }
	caseHit := false
	for i := 0; i < 10; i++ {
		cntr := 0
		_, _, found := MapElems(map[string]int{"test": 1, "test2": 2, "test3": 3}, f).Find(
			func(val basic.Pair[string, int]) (bool, error) {
				cntr += 1
				return val.GetA() == "test2", nil
			})
		test.True(found, t)
		caseHit = (caseHit || cntr < 3)
	}
	test.True(caseHit, t)
}

func mapElemsHelper[K comparable, V any](m map[K]V, t *testing.T) {
	f := func() basic.Pair[K, V] { return basic.Pair[K, V]{} }
	mIter := MapElems[K, V](m, f)
	for i := 0; i < len(m); i++ {
		mV, mErr, mBool := mIter(Continue)
		test.Eq(mV.GetB(), m[mV.GetA()], t)
		test.Nil(mErr, t)
		test.True(mBool, t)
	}
	_, mErr, mBool := mIter(Continue)
	test.Nil(mErr, t)
	test.False(mBool, t)
}
func TestMapElems(t *testing.T) {
	mapElemsHelper(map[string]int{}, t)
	mapElemsHelper(map[string]int{"test": 1}, t)
	mapElemsHelper(map[string]int{"test": 1, "test2": 2, "test3": 3}, t)
	mapElemsHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0}, t)
}

func TestMapElemsConsume(t *testing.T) {
	f := func() basic.Pair[string, int] { return basic.Pair[string, int]{} }
	test.Nil(MapElems[string, int](map[string]int{}, f).Consume(), t)
	test.Nil(MapElems[string, int](map[string]int{"test": 1}, f).Consume(), t)
	test.Nil(
		MapElems[string, int](map[string]int{"test": 1, "test2": 2}, f).Consume(),
		t,
	)
	test.Nil(
		MapElems[string, int](
			map[string]int{"test": 1, "test2": 2, "test3": 3}, f,
		).Consume(),
		t,
	)
}

func TestMapKeysStopEarly(t *testing.T) {
	caseHit := false
	for i := 0; i < 10; i++ {
		cntr := 0
		_, _, found := MapKeys(map[string]int{"test": 1, "test2": 2, "test3": 3}).Find(
			func(val string) (bool, error) {
				cntr += 1
				return val == "test2", nil
			})
		test.True(found, t)
		caseHit = (caseHit || cntr < 3)
	}
	test.True(caseHit, t)
}

func mapKeysHelper[K comparable, V any](m map[K]V, t *testing.T) {
	mIter := MapKeys[K, V](m)
	for i := 0; i < len(m); i++ {
		mV, mErr, mBool := mIter(Continue)
		_, ok := m[mV]
		test.True(ok, t)
		test.Nil(mErr, t)
		test.True(mBool, t)
	}
	_, mErr, mBool := mIter(Continue)
	test.Nil(mErr, t)
	test.False(mBool, t)
}
func TestMapKeys(t *testing.T) {
	mapKeysHelper(map[string]int{}, t)
	mapKeysHelper(map[string]int{"test": 1}, t)
	mapKeysHelper(map[string]int{"test": 1, "test2": 2, "test3": 3}, t)
	mapKeysHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0}, t)
}

func TestMapKeysConsume(t *testing.T) {
	test.Nil(MapKeys[string, int](map[string]int{}).Consume(), t)
	test.Nil(MapKeys[string, int](map[string]int{"test": 1}).Consume(), t)
	test.Nil(
		MapKeys[string, int](map[string]int{"test": 1, "test2": 2}).Consume(),
		t,
	)
	test.Nil(
		MapKeys[string, int](
			map[string]int{"test": 1, "test2": 2, "test3": 3},
		).Consume(),
		t,
	)
}

func TestMapValsStopEarly(t *testing.T) {
	caseHit := false
	for i := 0; i < 10; i++ {
		cntr := 0
		_, _, found := MapVals(map[string]int{"test": 1, "test2": 2, "test3": 3}).Find(
			func(val int) (bool, error) {
				cntr += 1
				return val == 2, nil
			})
		test.True(found, t)
		caseHit = (caseHit || cntr < 3)
	}
	test.True(caseHit, t)
}

func mapValsHelper[K comparable, V comparable](m map[K]V, t *testing.T) {
	mIter := MapVals[K, V](m)
	for i := 0; i < len(m); i++ {
		mV, mErr, mBool := mIter(Continue)
		found := false
		for _, v := range m {
			found = (found || v == mV)
		}
		test.True(found, t)
		test.Nil(mErr, t)
		test.True(mBool, t)
	}
	_, mErr, mBool := mIter(Continue)
	test.Nil(mErr, t)
	test.False(mBool, t)
}
func TestMapVals(t *testing.T) {
	mapValsHelper(map[string]int{}, t)
	mapValsHelper(map[string]int{"test": 1}, t)
	mapValsHelper(map[string]int{"test": 1, "test2": 2, "test3": 3}, t)
	mapValsHelper(map[int]float32{1: 1.0, 2: 2.0, 3: 3.0}, t)
}

func TestMapValsConsume(t *testing.T) {
	test.Nil(MapVals[string, int](map[string]int{}).Consume(), t)
	test.Nil(MapVals[string, int](map[string]int{"test": 1}).Consume(), t)
	test.Nil(
		MapVals[string, int](map[string]int{"test": 1, "test2": 2}).Consume(),
		t,
	)
	test.Nil(
		MapVals[string, int](
			map[string]int{"test": 1, "test2": 2, "test3": 3},
		).Consume(),
		t,
	)
}

func TestMapElemsChanClosing(t *testing.T) {
	test.NoPanic(
		func() {
			m := map[string]int{"test": 1, "test2": 2, "test3": 3}
			for i := 0; i < 10000; i++ {
				MapVals[string, int](m).Consume()
				m["four"] = i
			}
		},
		t,
	)
}
