package hash

import (
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestCombine(t *testing.T) {
	h1 := Hash(69)
	h2 := Hash(420)
	h3 := Hash(5280)

	test.Neq(h1.Combine(h2), h2.Combine(h1), t)
	test.Neq(h1.Combine(h3), h3.Combine(h1), t)
	test.Neq(h2.Combine(h1), h1.Combine(h2), t)
	test.Neq(h2.Combine(h3), h3.Combine(h2), t)
	test.Neq(h3.Combine(h1), h1.Combine(h3), t)
	test.Neq(h3.Combine(h2), h2.Combine(h3), t)

	// TODO - these specific orderings seem to always produce the same values, why?
	// test.Neq(h1.Combine(h2.Combine(h3)),h2.Combine(h1.Combine(h3)),t)
	test.Neq(h1.Combine(h2.Combine(h3)), h2.Combine(h3.Combine(h1)), t)
	test.Neq(h1.Combine(h2.Combine(h3)), h3.Combine(h2.Combine(h1)), t)

	test.Eq(h1.Combine(h2, h3), h1.Combine(h2).Combine(h3), t)
	test.Eq(h1.Combine(h3, h2), h1.Combine(h3).Combine(h2), t)
	test.Neq(h1.Combine(h2, h3), h1.Combine(h3, h2), t)

	test.Eq(h2.Combine(h1, h3), h2.Combine(h1).Combine(h3), t)
	test.Eq(h2.Combine(h3, h1), h2.Combine(h3).Combine(h1), t)
	test.Neq(h2.Combine(h1, h3), h2.Combine(h3, h1), t)

	test.Eq(h3.Combine(h1, h2), h3.Combine(h1).Combine(h2), t)
	test.Eq(h3.Combine(h2, h1), h3.Combine(h2).Combine(h1), t)
	test.Neq(h3.Combine(h1, h2), h3.Combine(h3, h2), t)
}

func TestCombineUnordered(t *testing.T) {
	h1 := Hash(69)
	h2 := Hash(420)
	h3 := Hash(5280)
	h4 := Hash(123456)
	h5 := Hash(7890)

	test.Eq(h1.CombineUnordered(h2), h2.CombineUnordered(h1), t)
	test.Eq(h1.CombineUnordered(h3), h3.CombineUnordered(h1), t)
	test.Eq(h2.CombineUnordered(h1), h1.CombineUnordered(h2), t)
	test.Eq(h2.CombineUnordered(h3), h3.CombineUnordered(h2), t)
	test.Eq(h3.CombineUnordered(h1), h1.CombineUnordered(h3), t)
	test.Eq(h3.CombineUnordered(h2), h2.CombineUnordered(h3), t)

	test.Eq(
		h1.CombineUnordered(h2.CombineUnordered(h3)),
		h2.CombineUnordered(h1.CombineUnordered(h3)),
		t,
	)
	test.Eq(
		h1.CombineUnordered(h2.CombineUnordered(h3)),
		h2.CombineUnordered(h3.CombineUnordered(h1)),
		t,
	)
	test.Eq(
		h1.CombineUnordered(h2.CombineUnordered(h3)),
		h3.CombineUnordered(h2.CombineUnordered(h1)),
		t,
	)

	test.Eq(
		h1.CombineUnordered(h2, h3, h4, h5),
		h1.CombineUnordered(h2, h3, h5, h4),
		t,
	)
	test.Eq(
		h1.CombineUnordered(h2, h3, h4, h5),
		h2.CombineUnordered(h3, h4, h1, h5),
		t,
	)
}
