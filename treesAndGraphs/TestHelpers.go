package treesAndGraphs

func buildSampleNode(childrenCount int) (sampleNode *GraphNode) {
	sampleNode = &GraphNode{Value: 0}
	for i := 1; i <= childrenCount; i++ {
		sampleNode.Children = append(sampleNode.Children, &GraphNode{Value: i})
	}
	return
}
