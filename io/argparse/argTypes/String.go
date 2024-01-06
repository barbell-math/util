package argTypes

type String string

func (s String)Translate(args []string) (string,error) {
    return string(s),nil
}

func (s String)ToVal() string {
    return string(s)
}

func (f String)NumArgs() int {
    return 1
}
