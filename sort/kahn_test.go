package sort

import (
	"fmt"
	"testing"
)

func TestKahnAlgorithm(t *testing.T) {
	// 创建一个示例图
	g := &Graph{
		Nodes: []int{3, 2, 6, 2, 5},
		Edges: [][]int{
			{2, 3},
			{3},
			{3, 4},
			{4},
			{},
		},
	}

	// 执行拓扑排序
	order, err := KahnAlgorithm(g)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Topological Order:", order)
	}
}
