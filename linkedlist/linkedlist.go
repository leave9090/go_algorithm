package linkedlist

import "fmt"

type (
	ListNode struct {
		value int
		next  *ListNode
	}
)

func (head *ListNode) GetNodeValue() int {
	return head.value
}

func (head *ListNode) GetNodeNext() *ListNode {
	return head.next
}

func (head *ListNode) InsertNode(value int) {
	head.value = value
}
func (head *ListNode) PrintNode() {
	for head != nil {
		fmt.Println(head.value)
		head = head.next
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
func (head *ListNode) add(value int) {
	head.value = value
	head.next = nil
}

func (head *ListNode) rangeLink() {
	point := head.next
	for nil != point {
		fmt.Println(point.value)
		point = point.next
	}
	fmt.Println("--------done----------")
}
