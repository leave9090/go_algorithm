package linkedlist

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

// 尾插法
func (l *LinkedList) TailAdd(value int) {
	newNode := &ListNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}

// 头插法
func (l *LinkedList) HeadAdd(value int) {
	newNode := &ListNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}

func (l *LinkedList) Print() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}

/*
leetcode 203
*/
func removeElements(head *ListNode, val int) *ListNode {

	return nil
}

/*
leetcode 237
*/

/*
leetcode 19
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	first, second := dummy, dummy

	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		second = second.Next
		first = first.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func (l *LinkedList) Length() int {
	var length int
	currentNode := l.Head
	for currentNode != nil {
		length++
		currentNode = currentNode.Next
	}
	return length
}

func (l *LinkedList) removeElements(val int) {
	ll := &ListNode{}
	ll.Next = l.Head
	pre, cur := ll, l.Head
	for cur != nil {
		if cur.Value == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	l.Head = ll.Next
}
