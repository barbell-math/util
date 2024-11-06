// This package defines a non-cryptographically secure hash type that is aimed
// towards being used for applications like hash maps.
package hash

type (
	Hash uint64
)

// Combines the supplied hashes with the hash that called the method in an
// order dependent fassion, meaning the same values sent in a different order
// are likely to produce a different hash. The hash that called the method will
// always be the zero-th value, meaning it will be the starting hash value
// before any combining operations are performed. After the first value the
// hashes will be combined in the order they are given.
func (h Hash) Combine(hashes ...Hash) Hash {
	newH := h
	for _, iterH := range hashes {
		newH = Hash(
			uint64(newH) ^ (uint64(iterH) + 0x517cc1b727220a95 + (uint64(newH) << 6) + (uint64(newH) >> 2)))
	}
	return newH
}

// Operates the same a [Hash.Combine] except zero hash values will be skipped
// when calculating hash values.
func (h Hash) CombineIgnoreZero(hashes ...Hash) Hash {
	newH := h
	for _, iterH := range hashes {
		if newH ==0 {
			newH=iterH
		} else if iterH!=0 {
			newH = Hash(
				uint64(newH) ^ (uint64(iterH) + 0x517cc1b727220a95 + (uint64(newH) << 6) + (uint64(newH) >> 2)))
		}
	}
	return newH
}

// Combines the supplied hashes with the hash that called the method in an
// order independent fassion, meaning the same values sent in a different
// order will produce the same hash. The hash that called the method will always
// be the zero-th value, meaning it will be the starting hash value before any
// cobining operations are performed.
func (h Hash) CombineUnordered(hashes ...Hash) Hash {
	newH := h
	for _, iterH := range hashes {
		newH = Hash(uint64(newH) ^ uint64(iterH))
	}
	return newH
}

// Operates the same a [Hash.CombineUnorderedIgnoreZero] except zero hash values
// will be skipped when calculating hash values.
func (h Hash) CombineUnorderedIgnoreZero(hashes ...Hash) Hash {
	newH := h
	for _, iterH := range hashes {
		if newH==0 {
			newH=iterH
		} else if iterH!=0 {
			newH = Hash(uint64(newH) ^ uint64(iterH))
		}
	}
	return newH
}
