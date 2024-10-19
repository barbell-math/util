package csv

import (
	"testing"

	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/test"
)

func TestToStructsNonStruct(t *testing.T) {
	cnt, err := ToStructs[int](
		iter.SliceElems[[]string](VALID_PARSE),
		NewOptions(),
	).Count()
	test.ContainsError(customerr.IncorrectType, err, t)
	test.Eq(0, cnt, t)
}

func TestToStructsInvalidStruct(t *testing.T) {
	type Row struct {
		One int `csv:"Two"`
		Two int
	}
	cnt, err := ToStructs[Row](
		iter.SliceElems[[]string](VALID_PARSE),
		NewOptions().DateTimeFormat("01/02/2006"),
	).Count()
	test.ContainsError(MalformedCSVStruct, err, t)
	test.ContainsError(DuplicateColName, err, t)
	test.Eq(0, cnt, t)
}

func TestToStructsValidStruct(t *testing.T) {
	cntr := 0
	err := ToStructs[csvTest](
		iter.SliceElems[[]string](VALID_PARSE),
		NewOptions().DateTimeFormat("01/02/2006"),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		VALID_STRUCT[index].Eq(&val, t)
		cntr++
		return iter.Continue, nil
	})
	test.Nil(err, t)
	test.Eq(2, cntr, t)
}

func TestToStructsMissingColumns(t *testing.T) {
	cntr := 0
	err := ToStructs[csvTest](
		iter.SliceElems[[]string](MISSING_COLUMNS_PARSE),
		NewOptions().DateTimeFormat("01/02/2006"),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		MISSING_COLUMNS_STRUCT[index].Eq(&val, t)
		cntr++
		return iter.Continue, nil
	})
	test.Nil(err, t)
	test.Eq(2, cntr, t)
}

func TestToStructsMissingValues(t *testing.T) {
	cntr := 0
	err := ToStructs[csvTest](
		iter.SliceElems[[]string](MISSING_VALUES_PARSE),
		NewOptions().DateTimeFormat("01/02/2006"),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		MISSING_VALUES_STRUCT[index].Eq(&val, t)
		cntr++
		return iter.Continue, nil
	})
	test.Nil(err, t)
	test.Eq(2, cntr, t)
}

func TestToStructsMissingHeadersWhenNotSpecifedAsMissing(t *testing.T) {
	cntr := 0
	err := ToStructs[csvTest](
		iter.SliceElems[[]string](MISSING_HEADERS_PARSE),
		NewOptions(),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		cntr++
		return iter.Continue, nil
	})
	test.ContainsError(MalformedCSVFile, err, t)
	test.Eq(0, cntr, t)
}

func TestToStructsMalformedDifferingRowLengths(t *testing.T) {
	cntr := 0
	err := ToStructs[csvTest](
		iter.SliceElems[[]string](MALFORMED_DIFFERING_ROW_LEN_PARSE),
		NewOptions().DateTimeFormat("01/02/2006"),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		MALFORMED_DIFFERING_ROW_LEN_STRUCT[index].Eq(&val, t)
		cntr++
		return iter.Continue, nil
	})
	test.ContainsError(MalformedCSVFile, err, t)
	test.Eq(1, cntr, t)
}

func TestToStructsMalformedDifferingRowLengthNoHeaders(t *testing.T) {
	cntr := 0
	err := ToStructs[csvTest](
		iter.SliceElems[[]string](MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS_PARSE),
		NewOptions().DateTimeFormat("01/02/2006").OptionsFlag(NewOptionsFlag().HasHeaders(false)),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS_STRUCT[index].Eq(&val, t)
		cntr++
		return iter.Continue, nil
	})
	test.ContainsError(MalformedCSVFile, err, t)
	test.Eq(1, cntr, t)
}

func TestToStructsMalformedTypes(t *testing.T) {
	cntr := 0
	err := ToStructs[csvTest](
		iter.SliceElems[[]string](MALFORMED_INT),
		NewOptions(),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		cntr++
		return iter.Continue, nil
	})
	test.ContainsError(MalformedCSVFile, err, t)
	test.Eq(0, cntr, t)

	err = ToStructs[csvTest](
		iter.SliceElems[[]string](MALFORMED_UINT),
		NewOptions(),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		cntr++
		return iter.Continue, nil
	})
	test.ContainsError(MalformedCSVFile, err, t)
	test.Eq(0, cntr, t)

	err = ToStructs[csvTest](
		iter.SliceElems[[]string](MALFORMED_FLOAT),
		NewOptions(),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		cntr++
		return iter.Continue, nil
	})
	test.ContainsError(MalformedCSVFile, err, t)
	test.Eq(0, cntr, t)

	err = ToStructs[csvTest](
		iter.SliceElems[[]string](MALFORMED_TIME),
		NewOptions().DateTimeFormat("01/02/2006"),
	).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
		cntr++
		return iter.Continue, nil
	})
	test.ContainsError(MalformedCSVFile, err, t)
	test.Eq(0, cntr, t)
}
