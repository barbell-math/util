package dataStruct

type Tree[N any] struct {
    Val N
    Children []Tree[N]
}
