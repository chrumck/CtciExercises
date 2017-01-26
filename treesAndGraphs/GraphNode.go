package treesAndGraphs

import (
	"container/list"
)

// GraphNode represents a graph node.
type GraphNode struct {
	Value    interface{}
	Children []*GraphNode
	visited  bool
}

// IsConnectedTo checks if the node is connected to the other node.
// Uses unidirectional Breadth-First Search
func (node *GraphNode) IsConnectedTo(otherNode *GraphNode) bool {
   nodesToCheck := list.New()
	node.visited = true
	nodesToCheck.PushFront(node)

	for nodesToCheck.Len() != 0 {
      nodeToCheck := nodesToCheck.Remove(nodesToCheck.Back()).(*GraphNode)
		for _, child := range nodeToCheck.Children {
			if child == otherNode {
				return true
			}
			child.visited = true
			nodesToCheck.PushFront(child)
		}
	}

	return false
}
