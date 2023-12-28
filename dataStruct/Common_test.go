package dataStruct

import (
    "github.com/barbell-math/util/dataStruct/types/static"
    "github.com/barbell-math/util/dataStruct/types/dynamic"
)

func staticQueueInterfaceTypeCheck[T any](d static.Queue[T]){}
func staticStackInterfaceTypeCheck[T any](d static.Stack[T]){}
func staticDequeInterfaceTypeCheck[T any](d static.Deque[T]){}
func staticVectorInterfaceTypeCheck[T any](d static.Vector[T]){}
func dynVectorInterfaceTypeCheck[T any](d dynamic.Vector[T]){}
