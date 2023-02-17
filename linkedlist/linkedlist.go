package linkedlist

import "fmt"

type (
	ListNode struct {
		value int
		next  *ListNode
	}
)

func (n *ListNode) GetNodeValue() int {
	return n.value
}

func (n *ListNode) GetNodeNext() *ListNode {
	return n.next
}

func (n *ListNode) InsertNode(value int) {
	n.value = value
}
func (n *ListNode) PrintNode() {
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}

func newList() *ListNode {
	return &ListNode{
		value: -1,
		next:  nil,
	}
}
func (head *ListNode) Insert(i int, e int) bool {
	p := head
	j := 1
	for nil != p && j < i {
		p = p.next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	s := &ListNode{value: e}
	s.next = p.next
	p.next = s
	return true
}
func (n *ListNode) add(value int) {
	n.value = value
	n.next = nil
}

func (n *ListNode) rangeLink() {
	point := n.next
	for nil != point {
		fmt.Println(point.value)
		point = point.next
	}
	fmt.Println("--------done----------")
}
