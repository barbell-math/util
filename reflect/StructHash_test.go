package reflect

import (
	"testing"

	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/test"
)

type (
	structHashTest struct {
		B bool
		I int
		I8 int8
		I16 int16
		I32 int32
		I64 int64
		U uint
		U8 uint8
		U16 uint16
		U32 uint32
		U64 uint64
		F32 float32
		F64 float64
		C64 complex64
		C128 complex128
		S string
		Uptr uintptr
		P *int
		C chan int
		A [3]int
		Sl []int
		M map[int]int
		F func()
		Intf interface{}
		St structTest
		Iter *structHashTest
	}
)

func TestStructHashNonStruct(t *testing.T) {
	i:=0
	_, err:=StructHash[int, *int](&i, NewStructHashOpts())
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructHashEmptyStruct(t *testing.T) {
	h, err:=StructHash[struct{}, *struct{}](&struct{}{}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(hash.Hash(0), h, t)
}

func TestStructHashSimpleStruct(t *testing.T) {
	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)
}

func TestStructHashNoExportedFields(t *testing.T) {
	type s struct { a int }

	h, err:=StructHash[s, *s](&s{a: 10}, NewStructHashOpts())
	test.Nil(err, t)
	test.Eq(hash.Hash(0), h, t)
}

func TestStructHashPopulatedPointers(t *testing.T) {
	i:=0	// A zero value will be hashed differently from nil

	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{P: &i},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	test.Neq(h, h2, t)
}

func TestStructHashPopulatedChan(t *testing.T) {
	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{C: make(chan int)},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	test.Neq(h, h2, t)
}

func TestStructHashPopulatedArray(t *testing.T) {
	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts().OptionsFlag(NewOptionsFlag().IncludeArrayVals(false)),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{A: [3]int{}},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	test.Neq(h, h2, t)
}

func TestStructHashPopulatedSlice(t *testing.T) {
	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts().OptionsFlag(NewOptionsFlag().IncludeSliceVals(false)),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{Sl: []int{}},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	test.Neq(h, h2, t)
}

func TestStructHashPopulatedMap(t *testing.T) {
	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts().OptionsFlag(NewOptionsFlag().IncludeMapVals(false)),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{M: map[int]int{}},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	test.Neq(h, h2, t)
}

func TestStructHashPopulatedFunc(t *testing.T) {
	f:=func() {}

	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{F: f},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	test.Neq(h, h2, t)
}

func TestStructHashPopulatedInterface(t *testing.T) {
	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{Intf: 5},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	test.Neq(h, h2, t)
}

func TestStructHashPopulatedStruct(t *testing.T) {
	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{St: structTest{One: 1}},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	test.Neq(h, h2, t)
}

func TestStructHashStability(t *testing.T) {
	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)
	test.Eq(h, h2, t)
}

func TestStructHashDontFollowPntrs(t *testing.T) {
	i:=0

	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{P: &i},
		NewStructHashOpts().OptionsFlag(NewOptionsFlag().FollowPntrs(false)),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{P: &i},
		NewStructHashOpts().OptionsFlag(NewOptionsFlag().FollowPntrs(false)),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)
	test.Eq(h, h2, t)

	i=1
	h3, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{P: &i},
		NewStructHashOpts().OptionsFlag(NewOptionsFlag().FollowPntrs(false)),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)
	test.Eq(h, h3, t)
}

func TestStructHashFollowPntrs(t *testing.T) {
	i:=0

	h, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{P: &i},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)

	h2, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{P: &i},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)
	test.Eq(h, h2, t)

	i=1
	h3, err:=StructHash[structHashTest, *structHashTest](
		&structHashTest{P: &i},
		NewStructHashOpts(),
	)
	test.Nil(err, t)
	test.Neq(hash.Hash(0), h, t)
	test.Neq(h, h3, t)
}

// TODO - test follow pntrs
// TODO - test include array vals
// TODO - test include slice vals
// TODO - test include map vals
// TODO - test follow interface
