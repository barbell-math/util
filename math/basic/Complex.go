package basic

import "unsafe"

func RealPart[N Complex, RN Float](v N) RN {
	switch any(v).(type) {
	case complex64:
		return RN(*(*float32)(unsafe.Pointer(&v)))
	case complex128:
		return RN(*(*float64)(unsafe.Pointer(&v)))
	default:
		return RN(0)
	}
}

func ImaginaryPart[N Complex, RN Float](v N) RN {
	switch any(v).(type) {
	case complex64:
		return RN(*(*float32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&v)) + unsafe.Sizeof(float32(0)),
		)))
	case complex128:
		return RN(*(*float64)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&v)) + unsafe.Sizeof(float64(0)),
		)))
	default:
		return RN(0)
	}
}
