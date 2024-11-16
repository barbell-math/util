package generics

func ZeroVal[T any]() T {
	var rv T
	return rv
}
