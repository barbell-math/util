package staticContainers


type Variant[T any, U any] interface {
    SetValA(newVal T) Variant[T,U];
    SetValB(newVal U) Variant[T,U];
    HasA() bool;
    HasB() bool;
    ValA() T;
    ValB() U;
    ValAOr(_default T) T;
    ValBOr(_default U) U;
};

type Pair[T any, U any] interface {
    GetA() T
    SetA(v T)
    GetB() U
    SetB(v U)
}
