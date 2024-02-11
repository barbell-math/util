package hash


func Combine(l uint64, r uint64) uint64 {
	return l^(r+0x517cc1b727220a95+(r<<6)+(r>>2))
}
