package dataStruct

import (
	"github.com/barbell-math/util/dataStruct/types"
	"github.com/barbell-math/util/dataStruct/types/dynamic"
	"github.com/barbell-math/util/dataStruct/types/static"
)

func staticQueueInterfaceTypeCheck[T any](d static.Queue[T]){}
func staticStackInterfaceTypeCheck[T any](d static.Stack[T]){}
func staticDequeInterfaceTypeCheck[T any](d static.Deque[T]){}
func staticVectorInterfaceTypeCheck[T any](d static.Vector[T]){}
func dynQueueInterfaceTypeCheck[T any](d dynamic.Queue[T]){}
func dynStackInterfaceTypeCheck[T any](d dynamic.Stack[T]){}
func dynDequeInterfaceTypeCheck[T any](d dynamic.Deque[T]){}
func dynVectorInterfaceTypeCheck[T any](d dynamic.Vector[T]){}

func writeInterfaceTypeCeck[K any, V any](d types.Write[K,V]){}
func readInterfaceTypeCeck[K any, V any](d types.Write[K,V]){}

func equalsInterfaceTypeCheck[T any, V any](d types.Equals[T,V]){}
