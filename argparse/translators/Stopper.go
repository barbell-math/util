package translators

type (
	Stopper[T any] struct { Err error }
)

func (s Stopper[T]) Translate(arg string) (T, error) {
	var tmp T
	return tmp, s.Err
}
