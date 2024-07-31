package csv

import (
	stdReflect "reflect"
	"strings"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/reflect"
)

type (
	headerMapping[R any] map[string]structIndex
)

func newStructHeaderMapping[R any](r *R, opts *options) (headerMapping[R], error) {
	rv := headerMapping[R]{}
	if !reflect.IsStructVal[R](r) {
		return rv, getBadRowTypeError(r)
	}
	err := iter.Zip[string, stdReflect.StructTag](
		reflect.StructFieldNames[R](r),
		reflect.StructFieldTags[R](r),
		basic.NewPair[string, stdReflect.StructTag],
	).ForEach(func(
		index int,
		val basic.Pair[string, stdReflect.StructTag],
	) (iter.IteratorFeedback, error) {
		// Skip unexported fields
		if strings.ToUpper(val.A[0:1])[0] != val.A[0] {
			return iter.Continue, nil
		}
		desiredVal := val.A
		if opts.getFlag(useStructTags) {
			if v, ok := val.B.Lookup(opts.structTagName); ok {
				desiredVal = v
			}
		}
		if _, ok := rv[desiredVal]; !ok {
			rv[desiredVal] = structIndex(index)
			return iter.Continue, nil
		} else {
			return iter.Break, customerr.AppendError(
				customerr.Wrap(
					DuplicateColName,
					"The struct duplicated a column name: %s.",
					desiredVal,
				),
			)
		}
	})
	return rv, err
}
