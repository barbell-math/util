package containers

// Code generated by ../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/container/tests"
	"testing"
)

func HashGraphToDirectedGraphInterfaceFactory(capacity int) dynamicContainers.DirectedGraph[int, int] {
	v := generateHashGraph(capacity)
	var rv dynamicContainers.DirectedGraph[int, int] = &v
	return rv
}

func TestHashGraph_DynDirectedGraphInterfaceSyncableInterface(t *testing.T) {
	tests.DynDirectedGraphInterfaceSyncableInterface(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceAddressableInterface(t *testing.T) {
	tests.DynDirectedGraphInterfaceAddressableInterface(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceClearInterface(t *testing.T) {
	tests.DynDirectedGraphInterfaceClearInterface(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_ReadDynDirectedGraphInterface(t *testing.T) {
	tests.ReadDynDirectedGraphInterface(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_WriteDynDirectedGraphInterface(t *testing.T) {
	tests.WriteDynDirectedGraphInterface(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceInterface(t *testing.T) {
	tests.DynDirectedGraphInterfaceInterface(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceStaticCapacityInterface(t *testing.T) {
	tests.DynDirectedGraphInterfaceStaticCapacityInterface(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceContainsEdge(t *testing.T) {
	tests.DynDirectedGraphInterfaceContainsEdge(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceContainsEdgePntr(t *testing.T) {
	tests.DynDirectedGraphInterfaceContainsEdgePntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceContainsVertex(t *testing.T) {
	tests.DynDirectedGraphInterfaceContainsVertex(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceContainsVertexPntr(t *testing.T) {
	tests.DynDirectedGraphInterfaceContainsVertexPntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceEdges(t *testing.T) {
	tests.DynDirectedGraphInterfaceEdges(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceEdgePntrs(t *testing.T) {
	tests.DynDirectedGraphInterfaceEdgePntrs(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceVertices(t *testing.T) {
	tests.DynDirectedGraphInterfaceVertices(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphInterfaceVerticePntrs(t *testing.T) {
	tests.DynDirectedGraphInterfaceVerticePntrs(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphLink(t *testing.T) {
	tests.DynDirectedGraphLink(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphLinkPntr(t *testing.T) {
	tests.DynDirectedGraphLinkPntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphNumOutLinks(t *testing.T) {
	tests.DynDirectedGraphNumOutLinks(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphNumOutLinksPntr(t *testing.T) {
	tests.DynDirectedGraphNumOutLinksPntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphOutEdges(t *testing.T) {
	tests.DynDirectedGraphOutEdges(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphOutEdgePntrs(t *testing.T) {
	tests.DynDirectedGraphOutEdgePntrs(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphOutVertices(t *testing.T) {
	tests.DynDirectedGraphOutVertices(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphOutVerticePntrs(t *testing.T) {
	tests.DynDirectedGraphOutVerticePntrs(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphOutEdgesAndVertices(t *testing.T) {
	tests.DynDirectedGraphOutEdgesAndVertices(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphOutEdgesAndVerticePntrs(t *testing.T) {
	tests.DynDirectedGraphOutEdgesAndVerticePntrs(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphEdgesBetween(t *testing.T) {
	tests.DynDirectedGraphEdgesBetween(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphEdgesBetweenPntr(t *testing.T) {
	tests.DynDirectedGraphEdgesBetweenPntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDeleteLink(t *testing.T) {
	tests.DynDirectedGraphDeleteLink(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDeleteLinkPntr(t *testing.T) {
	tests.DynDirectedGraphDeleteLinkPntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDeleteLinks(t *testing.T) {
	tests.DynDirectedGraphDeleteLinks(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDeleteLinksPntr(t *testing.T) {
	tests.DynDirectedGraphDeleteLinksPntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDeleteVertex(t *testing.T) {
	tests.DynDirectedGraphDeleteVertex(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDeleteVertexPntr(t *testing.T) {
	tests.DynDirectedGraphDeleteVertexPntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDeleteEdge(t *testing.T) {
	tests.DynDirectedGraphDeleteEdge(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDeleteEdgePntr(t *testing.T) {
	tests.DynDirectedGraphDeleteEdgePntr(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphClear(t *testing.T) {
	tests.DynDirectedGraphClear(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphKeyedEq(t *testing.T) {
	tests.DynDirectedGraphKeyedEq(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphIntersection(t *testing.T) {
	tests.DynDirectedGraphIntersection(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphUnion(t *testing.T) {
	tests.DynDirectedGraphUnion(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphDifference(t *testing.T) {
	tests.DynDirectedGraphDifference(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphIsSuperset(t *testing.T) {
	tests.DynDirectedGraphIsSuperset(HashGraphToDirectedGraphInterfaceFactory, t)
}

func TestHashGraph_DynDirectedGraphIsSubset(t *testing.T) {
	tests.DynDirectedGraphIsSubset(HashGraphToDirectedGraphInterfaceFactory, t)
}
