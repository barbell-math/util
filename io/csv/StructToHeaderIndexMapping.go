package csv

import (
	"math"
	"strings"

	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/widgets"
	"github.com/barbell-math/util/container/containers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/reflect"
)

type (
	structToHeaderIndexMapping[R any] map[structIndex]headerIndex
)

func newStructToHeaderIndexMapping[R any](
	structHeaderMapping headerMapping[R],
	r *R,
	opts *options,
) (structToHeaderIndexMapping[R], error) {
	rv := structToHeaderIndexMapping[R]{}
	if !reflect.IsStructVal[R](r) {
		return rv, getBadRowTypeError(r)
	}

	if !opts.getFlag(headersSupplied) {
		// Get all public data from the struct
		return rv, rv.getHeadersUntil(math.MaxInt, r)
	}

	if err := rv.detectHeaderDuplicates(opts); err != nil {
		return rv, err
	}

	return rv, iter.SliceElemPntrs[string](opts.headers).ForEach(
		func(index int, val *string) (iter.IteratorFeedback, error) {
			if structIdx, ok := structHeaderMapping[*val]; ok {
				rv[structIdx] = headerIndex(index)
			} else {
				return iter.Break, customerr.AppendError(
					InvalidHeaders,
					customerr.Wrap(
						InvalidHeader,
						"Invalid column: %s",
						*val,
					),
				)
			}
			return iter.Continue, nil
		},
	)
}

func (i *structToHeaderIndexMapping[R]) getHeadersUntil(num int, r *R) error {
	hIndex := 0
	return reflect.StructFieldNames[R](r).ForEach(
		func(index int, val string) (iter.IteratorFeedback, error) {
			if strings.ToUpper(val[0:1])[0] != val[0] {
				return iter.Continue, nil
			}
			(*i)[structIndex(hIndex)] = headerIndex(index)
			hIndex++
			if hIndex < num {
				return iter.Continue, nil
			} else {
				return iter.Break, nil
			}
		},
	)
}

func (i *structToHeaderIndexMapping[R]) detectHeaderDuplicates(
	opts *options,
) error {
	set, err := containers.NewHashSet[
		*string,
		widgets.Pntr[string, widgets.BuiltinString],
	](0)
	if err != nil {
		return err
	}
	if err := containers.Unique[*string](
		iter.SliceElemPntrs[string](opts.headers),
		&set,
		true,
	); err != nil {
		return customerr.AppendError(
			InvalidHeaders,
			customerr.Wrap(
				DuplicateColName,
				"The file headers duplicated a col name.",
			),
			err,
		)
	}
	return nil
}
