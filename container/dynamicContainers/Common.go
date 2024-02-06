// This package serves to define the set of dynamic containers and expose them
// as interfaces.
package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

type ComparisonsOtherConstraint[V any] interface {
	containerTypes.ReadOps[V]
	containerTypes.RWSyncable
	containerTypes.Length
}

type KeyedComparisonsOtherConstraint[K any, V any] interface {
	containerTypes.ReadKeyedOps[K,V]
	containerTypes.ReadOps[V] // TODO - needed??
	containerTypes.RWSyncable
	containerTypes.Length
}
