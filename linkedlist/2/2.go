package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 头插法
func (l *ListNode) InsertNode(val int) *ListNode {
	node := &ListNode{
		Val:  val,
		Next: l,
	}
	return node
}

// 尾插法
func (l *ListNode) AppendNode(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	for l.Next != nil {
		l = l.Next
	}
	l.Next = node

}

// 删除节点
func (l *ListNode) DeleteNode(val int) {

}
func (l *ListNode) PrintNode() {
	for l.Next != nil {
		println(l.Val)
		l = l.Next
	}
	println(l.Val)
}
