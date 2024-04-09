package sort

import (
	"fmt"
)

// Graph 表示有向无环图
type Graph struct {
	Nodes []int
	Edges [][]int
}

// KahnAlgorithm 实现Kahn算法进行拓扑排序
func KahnAlgorithm(g *Graph) ([]int, error) {
	inDegrees := make([]int, len(g.Nodes))
	for _, edge := range g.Edges {
		for _, v := range edge {
			inDegrees[v]++
		}
	}

	queue := []int{}
	for i, degree := range inDegrees {
		if degree == 0 {
			queue = append(queue, i)
		}
	}

	topologicalOrder := []int{}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			topologicalOrder = append(topologicalOrder, g.Nodes[node])

			for _, neighbor := range g.Edges[node] {
				inDegrees[neighbor]--
				if inDegrees[neighbor] == 0 {
					queue = append(queue, neighbor)
				}
			}
		}
		queue = queue[size:]
	}

	if len(topologicalOrder) != len(g.Nodes) {
		return nil, fmt.Errorf("Graph has a cycle, cannot perform topological sort")
	}

	return topologicalOrder, nil
}
