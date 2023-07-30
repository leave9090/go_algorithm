package binary_tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Val: val, Left: nil, Right: nil}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	}
	if val > root.Val {
		root.Right = insert(root.Right, val)
	}
	return root
}
