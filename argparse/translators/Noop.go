package translators

type (
	Noop[T any] struct{}
)

func (_ Noop[T]) Translate(arg string) (T, error) {
	var tmp T
	return tmp, nil
}
