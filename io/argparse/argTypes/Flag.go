package argTypes

type Flag bool

func (f Flag)Translate(args []string) (bool,error) {
    if len(args)>0 {
        return false,FlagDoesNotTakeArgs("")
    }
    return true,nil
}

func (f Flag)ToVal() bool {
    return bool(f)
}

func (f Flag)NumArgs() int {
    return NoArgs
}
