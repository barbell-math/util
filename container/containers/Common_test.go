package containers

import (
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
)

func syncableInterface(v containerTypes.RWSyncable){}
func lengthInterface(v containerTypes.Length){}
func capacityInterface(v containerTypes.Capacity){}
func readOpsInterface[T any, U any](v containerTypes.ReadOps[T,U]){}
func readKeyedOpsInterface[T any, U any](v containerTypes.ReadKeyedOps[T,U]){}
func writeOpsInterface[T any, U any](v containerTypes.WriteOps[T,U]){}
func writeKeyedOpsInterface[T any, U any](v containerTypes.WriteKeyedOps[T,U]){}
func deleteOpsInterface[T any, U any](v containerTypes.DeleteOps[T,U]){}
func deleteKeyedOpsInterface[T any, U any](v containerTypes.DeleteKeyedOps[T,U]){}
func firstElemReadInterface[T any](v containerTypes.FirstElemRead[T]){}
func firstElemWriteInterface[T any](v containerTypes.FirstElemWrite[T]){}
func firstElemDeleteInterface[T any](v containerTypes.FirstElemDelete[T]){}
func lastElemReadInterface[T any](v containerTypes.LastElemRead[T]){}
func lastElemWriteInterface[T any](v containerTypes.LastElemWrite[T]){}
func lastElemDeleteInterface[T any](v containerTypes.LastElemDelete[T]){}

func vectorReadInterface[U any](c dynamicContainers.ReadVector[int,U]){}
func vectorWriteInterface[U any](c dynamicContainers.WriteVector[int,U]){}
func vectorInterface[U any](c dynamicContainers.Vector[int,U]){}
