package containers

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"
import "github.com/barbell-math/util/container/basic"

// A pass through widget to represent the aliased type graphLink
// This is meant to be used with the containers from the [containers] package.
// Returns true if both graphLink's are equal. Uses the Eq operator provided by the basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash] widget internally.
func (_ *graphLink) Eq(l *graphLink, r *graphLink) bool {
	var tmp basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash]
	return tmp.Eq((*basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash])(l), (*basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash])(r))
}

// Returns true if a is less than r. Uses the Lt operator provided by the basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash] widget internally.
func (_ *graphLink) Lt(l *graphLink, r *graphLink) bool {
	var tmp basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash]
	return tmp.Lt((*basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash])(l), (*basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash])(r))
}

// Provides a hash function for the value that it is wrapping. The value that is returned will be supplied by the basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash] widget internally.
func (_ *graphLink) Hash(other *graphLink) hash.Hash {
	var tmp basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash]
	return tmp.Hash((*basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash])(other))
}

// Zeros the supplied value. The operation that is performed will be determined by the basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash] widget internally.
func (_ *graphLink) Zero(other *graphLink) {
	var tmp basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash]
	tmp.Zero((*basic.WidgetPair[edgeHash, vertexHash, *edgeHash, *vertexHash])(other))
}
