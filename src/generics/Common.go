package generics

type (
	// A generic filter
	Filter[V any] func(thing V) bool
)

func ZeroVal[T any]() T {
	var rv T
	return rv
}

func NoFilter[T any](thing T) bool  { return true }
func AllFilter[T any](thing T) bool { return false }

func GenFilter[T comparable](inverse bool, things ...T) Filter[T] {
	return func(thing T) bool {
		rv := inverse
		for i := 0; ((inverse && rv) || (!inverse && !rv)) && i < len(things); i++ {
			if inverse {
				rv = (rv && thing != things[i])
			} else {
				rv = (thing == things[i])
			}
		}
		return rv
	}
}
