package dataStruct

import (
    "github.com/barbell-math/util/dataStruct/types/static"
)

func staticQueueInterfaceTypeCheck[T any](d static.Queue[T]){}
func staticStackInterfaceTypeCheck[T any](d static.Stack[T]){}
func staticVectorInterfaceTypeCheck[T any](d static.Vector[T]){}
