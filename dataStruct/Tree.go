package dataStruct

type Tree[N interface{}] struct {
    Data N
    Children []Tree[N]
}
