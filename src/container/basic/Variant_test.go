package basic

import (
	"testing"

	"github.com/barbell-math/util/src/test"
)

type testStruct struct {
	a int
	b float32
}

func TestVariantA(t *testing.T) {
	var tmp int = 5
	v := Variant[int, testStruct]{}
	newV := v.SetValA(tmp)
	test.True(newV.HasA(), t)
	test.False(newV.HasB(), t)
	test.Eq(5, newV.ValA(), t)
	test.Eq(5, newV.ValAOr(1), t)
	test.Eq(testStruct{}, newV.ValBOr(testStruct{}), t)
}

func TestVariantB(t *testing.T) {
	var tmp testStruct = testStruct{a: 1, b: 2}
	v := Variant[int, testStruct]{}
	newV := v.SetValB(tmp)
	test.False(newV.HasA(), t)
	test.True(newV.HasB(), t)
	test.Eq(tmp, newV.ValB(), t)
	test.Eq(1, newV.ValAOr(1), t)
	test.Eq(tmp, newV.ValBOr(testStruct{}), t)
}
