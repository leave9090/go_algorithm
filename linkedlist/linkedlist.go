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
