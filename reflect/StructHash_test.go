package reflect

import (
	"testing"
	"unsafe"

	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/test"
)

func TestStructHashNonStruct(t *testing.T) {
	i := 0
	_, err := StructHash[int, *int](&i, NewStructHashOpts())
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructHashEmptyStruct(t *testing.T) {
	h, err := StructHash[struct{}, *struct{}](&struct{}{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(hash.Hash(0), h, t)
}

func TestStructHashNoExportedFields(t *testing.T) {
	type s struct{ a int }

	h, err := StructHash[s, *s](&s{a: 10}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(hash.Hash(0), h, t)
}

func TestStructHashPointers(t *testing.T) {
	type P struct{ P *int }
	i := 0

	h, err := StructHash[P, *P](&P{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[P, *P](&P{&i}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[P, *P](
		&P{},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().FollowPntrs(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[P, *P](
		&P{&i},
		NewStructHashOpts().SetOptionsFlag(NewOptionsFlag().FollowPntrs(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(uintptr(unsafe.Pointer(&i))), t)
}

func TestStructHashArray(t *testing.T) {
	type A struct{ A [3]int }
	i := [3]int{1, 2, 3}

	h, err := StructHash[A, *A](&A{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(3), t)

	h, err = StructHash[A, *A](&A{i}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(3).CombineIgnoreZero(1, 2, 3), t)

	h, err = StructHash[A, *A](
		&A{},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().IncludeArrayVals(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(3), t)

	h, err = StructHash[A, *A](
		&A{i},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().IncludeArrayVals(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(3), t)
}

func TestStructHashSlice(t *testing.T) {
	type S struct{ S []int }
	i := []int{1, 2, 3}

	h, err := StructHash[S, *S](&S{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[S, *S](&S{i}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(3).CombineIgnoreZero(1, 2, 3), t)

	h, err = StructHash[S, *S](
		&S{},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().IncludeSliceVals(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[S, *S](
		&S{i},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().IncludeSliceVals(false)),
	)
	test.Nil(err, t)
	test.Eq(
		h,
		hash.Hash(3).CombineIgnoreZero(hash.Hash(uintptr(unsafe.Pointer(&i[0])))),
		t,
	)
}

func TestStructHashMap(t *testing.T) {
	type M struct{ M map[int]int }
	i := map[int]int{1: 2, 3: 4}

	h, err := StructHash[M, *M](&M{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[M, *M](&M{i}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(2).CombineUnorderedIgnoreZero(1, 2, 3, 4), t)

	h, err = StructHash[M, *M](
		&M{},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().IncludeMapVals(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[M, *M](
		&M{i},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().IncludeMapVals(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(2), t)
}

func TestStructHashInterface(t *testing.T) {
	type I struct{ I any }
	i := 1

	h, err := StructHash[I, *I](&I{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[I, *I](&I{i}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(1), t)

	h, err = StructHash[I, *I](
		&I{},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().FollowInterfaces(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[I, *I](
		&I{i},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().FollowInterfaces(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)
}

func TestStructHashStructs(t *testing.T) {
	type S1 struct{ A int }
	type S2 struct{ S S1 }

	h, err := StructHash[S2, *S2](&S2{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[S2, *S2](&S2{S: S1{1}}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(1), t)

	h, err = StructHash[S2, *S2](
		&S2{},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().RecurseStructs(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[S2, *S2](
		&S2{S: S1{1}},
		NewStructHashOpts().
			SetOptionsFlag(NewOptionsFlag().RecurseStructs(false)),
	)
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)
}

func TestStructHashMultipleFields(t *testing.T) {
	type S struct {
		A int
		B uint
	}

	h, err := StructHash[S, *S](&S{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(0), t)

	h, err = StructHash[S, *S](&S{A: 1}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(1), t)

	h, err = StructHash[S, *S](&S{B: 1}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(1), t)

	h, err = StructHash[S, *S](&S{1, 2}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(h, hash.Hash(1).CombineIgnoreZero(2), t)
}
