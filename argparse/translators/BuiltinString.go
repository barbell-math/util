package translators

type (
	BuiltinString struct{}
)

func (_ BuiltinString) Translate(arg string) (string, error) {
	return arg, nil
}
