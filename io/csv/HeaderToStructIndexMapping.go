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
	headerToStructIndexMapping[R any] map[headerIndex]structIndex
)

func newHeaderToStructIndexMapping[R any](
	src iter.Iter[[]string],
	structHeaderMapping headerMapping[R],
	r *R,
	opts *options,
) (headerToStructIndexMapping[R], error) {
	rv := headerToStructIndexMapping[R]{}
	if !reflect.IsStructVal[R](r) {
		return rv, getBadRowTypeError(r)
	}

	if !opts.getFlag(hasHeaders) {
		// No way to get the correct number of headers without pulling a row of
		// data which would need to be parsed later so just get all the indexes.
		return rv, rv.getHeadersUntil(math.MaxInt, r)
	}

	headers, err, ok := src.PullOne()
	if !ok || err != nil {
		return headerToStructIndexMapping[R]{}, err
	}
	if opts.getFlag(ignoreHeaders) {
		// Get all struct indexes up to the total number of headers. The headers
		// are ignored and the structs field ordering determines the indexes.
		return rv, rv.getHeadersUntil(len(headers), r)
	}
	if err := rv.detectHeaderDuplicates(headers); err != nil {
		return rv, err
	}

	return rv, iter.SliceElemPntrs[string](headers).ForEach(
		func(index int, val *string) (iter.IteratorFeedback, error) {
			if structIdx, ok := structHeaderMapping[*val]; ok {
				rv[headerIndex(index)] = structIdx
			} else {
				return iter.Break, customerr.Wrap(
					InvalidHeader,
					"Invalid column: %s",
					*val,
				)
			}
			return iter.Continue, nil
		},
	)
}

func (i *headerToStructIndexMapping[R]) getHeadersUntil(num int, r *R) error {
	hIndex := 0
	return reflect.StructFieldNames[R](r).ForEach(
		func(index int, val string) (iter.IteratorFeedback, error) {
			if strings.ToUpper(val[0:1])[0] != val[0] {
				return iter.Continue, nil
			}
			(*i)[headerIndex(hIndex)] = structIndex(index)
			hIndex++
			if hIndex < num {
				return iter.Continue, nil
			} else {
				return iter.Break, nil
			}
		},
	)
}

func (i *headerToStructIndexMapping[R]) detectHeaderDuplicates(headers []string) error {
	set, err := containers.NewHashSet[
		*string,
		widgets.Pntr[string, widgets.BuiltinString],
	](0)
	if err != nil {
		return err
	}
	if err := containers.Unique[*string](
		iter.SliceElemPntrs[string](headers),
		&set,
		true,
	); err != nil {
		return customerr.AppendError(
			customerr.Wrap(
				DuplicateColName,
				"The file headers duplicated a col name.",
			),
			err,
		)
	}
	return nil
}
