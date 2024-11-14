package csv

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func TestNewStructToHeaderIndexMappingNonStruct(t *testing.T) {
	var tmp int
	m, err := newStructToHeaderIndexMapping[int](
		headerMapping[int]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions(),
	)
	test.ContainsError(customerr.IncorrectType, err, t)
	test.Eq(0, len(m), t)
}

func TestNewStructToHeaderIndexMappingNominalCase(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newStructToHeaderIndexMapping[Row](
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions(),
	)
	test.Nil(err, t)
	test.Eq(3, len(m), t)
	test.Eq(headerIndex(0), m[structIndex(0)], t)
	test.Eq(headerIndex(1), m[structIndex(1)], t)
	test.Eq(headerIndex(2), m[structIndex(2)], t)
}

func TestNewStructToHeaderIndexMappingDuplicateHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newStructToHeaderIndexMapping[Row](
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().SetHeaders([]string{"One", "Two", "One"}),
	)
	test.ContainsError(DuplicateColName, err, t)
	test.ContainsError(containerTypes.Duplicate, err, t)
	test.Eq(0, len(m), t)
}

func TestNewStructToHeaderIndexMappingSameHeadersSameOrder(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newStructToHeaderIndexMapping[Row](
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().SetHeaders([]string{"One", "Two", "Three"}),
	)
	test.Nil(err, t)
	test.Eq(3, len(m), t)
	test.Eq(headerIndex(0), m[structIndex(0)], t)
	test.Eq(headerIndex(1), m[structIndex(1)], t)
	test.Eq(headerIndex(2), m[structIndex(2)], t)
}

func TestNewStructToHeaderIndexMappingSameHeadersDifferentOrder(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newStructToHeaderIndexMapping[Row](
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().SetHeaders([]string{"Three", "One", "Two"}),
	)
	test.Nil(err, t)
	test.Eq(3, len(m), t)
	test.Eq(headerIndex(1), m[structIndex(0)], t)
	test.Eq(headerIndex(2), m[structIndex(1)], t)
	test.Eq(headerIndex(0), m[structIndex(2)], t)
}

func TestNewStructToHeaderIndexMappingIncorrectHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newStructToHeaderIndexMapping[Row](
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().SetHeaders([]string{"Very", "Bad", "Headers"}),
	)
	test.ContainsError(InvalidHeaders, err, t)
	test.ContainsError(InvalidHeader, err, t)
	test.Eq(0, len(m), t)
}

func TestNewStructToHeaderIndexMappingMissingHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newStructToHeaderIndexMapping[Row](
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().SetHeaders([]string{"One", "Two"}),
	)
	test.Nil(err, t)
	test.Eq(2, len(m), t)
	test.Eq(headerIndex(0), m[structIndex(0)], t)
	test.Eq(headerIndex(1), m[structIndex(1)], t)
}

func TestNewStructToHeaderIndexMappingExtraHeaders(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newStructToHeaderIndexMapping[Row](
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().SetHeaders([]string{"One", "Two", "Three", "Four"}),
	)
	test.ContainsError(InvalidHeaders, err, t)
	test.ContainsError(InvalidHeader, err, t)
	test.Eq(3, len(m), t)
	test.Eq(headerIndex(0), m[structIndex(0)], t)
	test.Eq(headerIndex(1), m[structIndex(1)], t)
	test.Eq(headerIndex(2), m[structIndex(2)], t)
}

func TestNewStructToHeaderIndexMappingExtraHeadersWithDuplicate(t *testing.T) {
	type Row struct {
		One   int
		Two   int
		Three int
	}
	var tmp Row
	m, err := newStructToHeaderIndexMapping[Row](
		headerMapping[Row]{"One": 0, "Two": 1, "Three": 2},
		&tmp,
		NewOptions().SetHeaders([]string{"One", "Two", "Three", "One"}),
	)
	test.ContainsError(DuplicateColName, err, t)
	test.ContainsError(containerTypes.Duplicate, err, t)
	test.Eq(0, len(m), t)
}
