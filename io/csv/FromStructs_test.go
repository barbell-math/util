package csv

import (
	"testing"

	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/test"
)

func TestFromStructsNonStruct(t *testing.T) {
	cnt, err := FromStructs[int](
		iter.SliceElems[int]([]int{0, 1, 2}),
		NewOptions(),
	).Count()
	test.ContainsError(customerr.IncorrectType, err, t)
	test.Eq(0, cnt, t)
}

func TestFromStructsInvalidStruct(t *testing.T) {
	type Row struct {
		One int `csv:"Two"`
		Two int
	}
	cnt, err := FromStructs[Row](
		iter.SliceElems[Row]([]Row{
			{One: 1, Two: 2},
		}),
		NewOptions().DateTimeFormat("01/02/2006"),
	).Count()
	test.ContainsError(MalformedCSVStruct, err, t)
	test.ContainsError(DuplicateColName, err, t)
	test.Eq(0, cnt, t)
}

func fromStructsEqualityHelper(exp [][]string, got [][]string, t *testing.T) {
	test.Eq(len(exp), len(got), t)
	for i := 0; i < len(got); i++ {
		test.Eq(len(exp[i]), len(got[i]), t)
		for j := 0; j < len(got[i]); j++ {
			test.Eq(exp[i][j], got[i][j], t)
		}
	}
}

func TestFromStructsValidStruct(t *testing.T) {
	res, err := FromStructs[csvTest](
		iter.SliceElems[csvTest](VALID_STRUCT),
		NewOptions().DateTimeFormat("01/02/2006"),
	).Collect()
	test.Nil(err, t)
	fromStructsEqualityHelper(VALID_FROM_STRUCT, res, t)
}

func TestFromStructsValidStructDontWriteHeaders(t *testing.T) {
	res, err := FromStructs[csvTest](
		iter.SliceElems[csvTest](VALID_STRUCT),
		NewOptions().DateTimeFormat("01/02/2006").WriteHeaders(false),
	).Collect()
	test.Nil(err, t)
	test.Eq(len(VALID_FROM_STRUCT)-1, len(res), t)
	for i := 0; i < len(res); i++ {
		test.Eq(len(VALID_FROM_STRUCT[i+1]), len(res[i]), t)
		for j := 0; j < len(res); j++ {
			test.Eq(VALID_FROM_STRUCT[i+1][j], res[i][j], t)
		}
	}
}

func TestFromStructsMissingColumnsWithHeadersSpecified(t *testing.T) {
	res, err := FromStructs[csvTest](
		iter.SliceElems[csvTest](MISSING_COLUMNS_STRUCT),
		NewOptions().DateTimeFormat("01/02/2006").Headers([]string{
			"I8", "Ui", "Ui8", "F32", "S", "B", "T",
		}),
	).Collect()
	test.Nil(err, t)
	fromStructsEqualityHelper(MISSING_COLUMNS_FROM_STRUCT, res, t)
}

func TestFromStructMissingValues(t *testing.T) {
	res, err := FromStructs[csvTest](
		iter.SliceElems[csvTest](MISSING_VALUES_STRUCT),
		NewOptions().DateTimeFormat("01/02/2006"),
	).Collect()
	test.Nil(err, t)
	fromStructsEqualityHelper(MISSING_VALUES_FROM_STRUCT, res, t)
}

func TestFromStructsMissingHeaders(t *testing.T) {
	res, err := FromStructs[csvTest](
		iter.SliceElems[csvTest](MISSING_HEADERS_STRUCT),
		NewOptions().DateTimeFormat("01/02/2006").WriteHeaders(false),
	).Collect()
	test.Nil(err, t)
	fromStructsEqualityHelper(MISSING_HEADERS_FROM_STRUCT, res, t)
}

func TestFromStructsStrings(t *testing.T) {
	res, err := FromStructs[csvTest](
		iter.SliceElems[csvTest](STRINGS_STRUCT),
		NewOptions().DateTimeFormat("01/02/2006").Headers([]string{"S", "S1"}),
	).Collect()
	test.Nil(err, t)
	fromStructsEqualityHelper(STRINGS_FROM_STRUCT, res, t)
}

func TestFromStructsNewlineString(t *testing.T) {
	type Row struct {
		Str string
	}
	res, err := FromStructs[Row](
		iter.SliceElems[Row]([]Row{{"hello\nworld"}}),
		NewOptions(),
	).Collect()
	test.Nil(err, t)
	test.Eq(2, len(res), t)
	test.SlicesMatch[string]([]string{"Str"}, res[0], t)
	test.SlicesMatch[string]([]string{"\"hello\nworld\""}, res[1], t)
}

func TestFromStructsQuotedString(t *testing.T) {
	type Row struct {
		Str string
	}
	res, err := FromStructs[Row](
		iter.SliceElems[Row]([]Row{{"\"hello\n\"world\"\""}}),
		NewOptions(),
	).Collect()
	test.Nil(err, t)
	test.Eq(2, len(res), t)
	test.SlicesMatch[string]([]string{"Str"}, res[0], t)
	test.SlicesMatch[string]([]string{"\"\"\"hello\n\"\"world\"\"\"\"\""}, res[1], t)
}

func TestFromStructsWithTags(t *testing.T) {
	type Row struct {
		I int    `csv:"int"`
		S string `csv:"string"`
	}
	res, err := FromStructs[Row](
		iter.SliceElems[Row]([]Row{{1, "One"}}),
		NewOptions(),
	).Collect()
	test.Nil(err, t)
	test.Eq(2, len(res), t)
	test.SlicesMatch[string]([]string{"int", "string"}, res[0], t)
	test.SlicesMatch[string]([]string{"1", "One"}, res[1], t)
}
