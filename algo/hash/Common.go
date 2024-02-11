package hash


type (
	Hash uint64
)

func (h Hash)Combine(other uint64) Hash {
	return Hash(uint64(h)^(other+0x517cc1b727220a95+(other<<6)+(other>>2)))
}
