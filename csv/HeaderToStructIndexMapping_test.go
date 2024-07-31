package csv

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/test"
)

func TestNewHeaderToStructIndexMappingNonStruct(t *testing.T) {
	var tmp int
	m, err := newHeaderToStructIndexMapping[int](
		iter.SliceElems[[]string]([][]string{{"One", "Two", "Three"}}),
		headerMapping[int]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions(),
	)
	test.ContainsError(customerr.IncorrectType, err, t)
	test.Eq(0, len(m), t)
}

func TestNewHeaderToStructIndexMappingNominalCase(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newHeaderToStructIndexMapping[Row](
		iter.SliceElems[[]string]([][]string{{"One", "Two", "Three"}}),
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions(),
	)
	test.Nil(err, t)
	test.Eq(3, len(m), t)
	test.Eq(structIndex(0), m[headerIndex(0)], t)
	test.Eq(structIndex(1), m[headerIndex(1)], t)
	test.Eq(structIndex(2), m[headerIndex(2)], t)
}

func TestNewHeaderToStructIndexMappingMissingHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newHeaderToStructIndexMapping[Row](
		iter.SliceElems[[]string]([][]string{{"One", "Three"}}),
		headerMapping[Row]{"One": 0, "Three": 2},
		&tmp,
		NewOptions(),
	)
	test.Nil(err, t)
	test.Eq(2, len(m), t)
	test.Eq(structIndex(0), m[headerIndex(0)], t)
	test.Eq(structIndex(2), m[headerIndex(1)], t)
}

func TestNewHeaderToStructIndexMappingMissingStructFields(t *testing.T) {
	type Row struct {
		One   int
		Three int
	}
	var tmp Row
	m, err := newHeaderToStructIndexMapping[Row](
		iter.SliceElems[[]string]([][]string{{"One", "Two", "Three"}}),
		headerMapping[Row]{"One": 0, "Three": 1},
		&tmp,
		NewOptions(),
	)
	test.ContainsError(InvalidHeader, err, t)
	test.Eq(1, len(m), t)
	test.Eq(structIndex(0), m[headerIndex(0)], t)
}

func TestNewHeaderToStructIndexMappingDuplicateHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newHeaderToStructIndexMapping[Row](
		iter.SliceElems[[]string]([][]string{{"One", "Two", "One"}}),
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions(),
	)
	test.ContainsError(DuplicateColName, err, t)
	test.ContainsError(containerTypes.Duplicate, err, t)
	test.Eq(0, len(m), t)
}

func TestNewHeaderToStructIndexMappingIgnoreHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newHeaderToStructIndexMapping[Row](
		iter.SliceElems[[]string]([][]string{{"Some", "Bad", "Headers"}}),
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().IgnoreHeaders(true),
	)
	test.Nil(err, t)
	test.Eq(3, len(m), t)
	test.Eq(structIndex(0), m[headerIndex(0)], t)
	test.Eq(structIndex(1), m[headerIndex(1)], t)
	test.Eq(structIndex(2), m[headerIndex(2)], t)
}

func TestNewHeaderToStructIndexMappingIgnoreHeadersMissingHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newHeaderToStructIndexMapping[Row](
		iter.SliceElems[[]string]([][]string{{"Some", "Headers"}}),
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().IgnoreHeaders(true),
	)
	test.Nil(err, t)
	test.Eq(2, len(m), t)
	test.Eq(structIndex(0), m[headerIndex(0)], t)
	test.Eq(structIndex(1), m[headerIndex(1)], t)
}

func TestNewHeaderToStructIndexMappingIgnoreHeadersMoreHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newHeaderToStructIndexMapping[Row](
		iter.SliceElems[[]string]([][]string{{"Some", "Very", "Bad", "Headers"}}),
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().IgnoreHeaders(true),
	)
	test.Nil(err, t)
	test.Eq(3, len(m), t)
	test.Eq(structIndex(0), m[headerIndex(0)], t)
	test.Eq(structIndex(1), m[headerIndex(1)], t)
	test.Eq(structIndex(2), m[headerIndex(2)], t)
}

func TestNewHeaderToStructIndexMappingNoHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newHeaderToStructIndexMapping[Row](
		iter.SliceElems[[]string]([][]string{{"1", "2", "3"}}),
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().HasHeaders(false),
	)
	test.Nil(err, t)
	test.Eq(3, len(m), t)
	test.Eq(structIndex(0), m[headerIndex(0)], t)
	test.Eq(structIndex(1), m[headerIndex(1)], t)
	test.Eq(structIndex(2), m[headerIndex(2)], t)
}
