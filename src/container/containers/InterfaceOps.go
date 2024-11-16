package containers

import (
	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/container/containerTypes"
	"github.com/barbell-math/util/src/container/dynamicContainers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/iter"
)

// MapKeyedUnion takes all of the keys value pairs in r and places them in l.
// It keys are present in both l and r, the value from r will be the value that
// is present in the final map.
func MapKeyedUnion[K any, V any](
	l dynamicContainers.Map[K, V],
	r dynamicContainers.Map[K, V],
) {
	r.Keys().ForEach(
		func(index int, key K) (iter.IteratorFeedback, error) {
			val, _ := r.Get(key)
			l.Emplace(basic.Pair[K, V]{key, val})
			return iter.Continue, nil
		},
	)
}

// MapKeyedUnion takes all of the keys value pairs in r and places them in l.
// The key set of l must be fully disjoint from the key set of r, otherwise a
// [containerTypes.Duplicate] error will be returned. The state of l will not be
// known if duplicate keys are found.
func MapDisjointKeyedUnion[K any, V any](
	l dynamicContainers.Map[K, V],
	r dynamicContainers.Map[K, V],
) error {
	return r.Keys().ForEach(
		func(index int, key K) (iter.IteratorFeedback, error) {
			if _, err := l.Get(key); err == nil {
				return iter.Break, customerr.Wrap(
					containerTypes.Duplicate, "%+v", key,
				)
			}
			val, _ := r.Get(key)
			l.Emplace(basic.Pair[K, V]{key, val})
			return iter.Continue, nil
		},
	)
}
