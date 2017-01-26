package treesAndGraphs

import (
	"testing"
)

func TestIsConnected_WhenNodesConnected_ReturnsTrue(test *testing.T) {

	// Arrange
	childrenCount := 3
	node := buildSampleNode(childrenCount)
	childNode := buildSampleNode(childrenCount)
	node.Children = append(node.Children, childNode)

	// Act
	result := node.IsConnectedTo(childNode.Children[1])

	// Assert
	if !result {
		test.Error("Should return true but returned false.")
	}
}

func TestIsConnected_WhenNodesNotConnected_ReturnsFalse(test *testing.T) {

	// Arrange
	childrenCount := 3
	node := buildSampleNode(childrenCount)
	otherNode := &GraphNode{Value: 5}

	// Act
	result := node.IsConnectedTo(otherNode)

	// Assert
	if result {
		test.Error("Should return false but returned true.")
	}
}

func TestIsConnected_WhenNodeConnectedToItself_ReturnsTrue(test *testing.T) {

	// Arrange
	childrenCount := 3
	node := buildSampleNode(childrenCount)
	node.Children = append(node.Children, node)

	// Act
	result := node.IsConnectedTo(node)

	// Assert
	if !result {
		test.Error("Should return true but returned false.")
	}
}

func TestIsConnected_WhenChildrenEmpty_ReturnsFalse(test *testing.T) {

	// Arrange
	node := &GraphNode{Value: 1}
	otherNode := &GraphNode{Value: 2}

	// Act
	result := node.IsConnectedTo(otherNode)

	// Assert
	if result {
		test.Error("Should return false but returned true.")
	}
}
