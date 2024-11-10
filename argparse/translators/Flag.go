package translators

type (
	Flag struct{}
)

func (_ Flag) Translate(arg string) (bool, error) {
	return true, nil
}
