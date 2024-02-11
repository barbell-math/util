package hash


type (
	Hash uint64
)

func (h Hash)Combine(other Hash) Hash {
	return Hash(
		uint64(h)^(
			uint64(other)+0x517cc1b727220a95+(
				uint64(other)<<6)+(uint64(other)>>2)))
}
