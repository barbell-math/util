package hash


type (
	Hash uint64
)

// Combines the supplied hashes with the hash that called the method in an 
// orderer dependent fassion, meaning the same values sent in a different order 
// are likely to produce a different hash. The hash that called the method will 
// always be the zero-th value, meaning it will be the starting hash value 
// before any combining operations are performed. After the first value the 
// hashes will be combined in the order they are given.
func (h Hash)Combine(hashes ...Hash) Hash {
	newH:=h
	for _,iterH:=range(hashes) {
		newH=Hash(
			uint64(newH)^(
				uint64(iterH)+0x517cc1b727220a95+(
					uint64(newH)<<6)+(uint64(newH)>>2)))
	}
	return newH
}

// Combines the supplied hashes with the hash that called the method in an 
// orderer independent fassion, meaning the same values sent in a different
// order will produce the same hash. The hash that called the method will always
// be the zero-th value, meaning it will be the starting hash value before any 
// cobining operations are performed.
func (h Hash)CombineUnordered(hashes ...Hash) Hash {
	newH:=h
	for _,iterH:=range(hashes) {
		newH=Hash(uint64(newH)^uint64(iterH))
	}
	return newH
}
