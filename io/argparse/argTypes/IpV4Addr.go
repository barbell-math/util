package argTypes

type IpV4Addr [4]int8

var LocalHost IpV4Addr=[4]int8{127,0,0,1}

func (i IpV4Addr)Translate(args []string) (IpV4Addr,error) {
    return i,nil
}

func (i IpV4Addr)ToVal() IpV4Addr {
    return i
}

func (i IpV4Addr)NumArgs() int {
    return 1
}
