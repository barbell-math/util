package hash

type Hashable[H ~uint32 | ~uint64] interface {
    Hash() H
}
