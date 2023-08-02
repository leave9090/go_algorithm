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
	pseudoHead := &ListNode{}
	pseudoHead.Next = head
	prev, cur := pseudoHead, head

	for cur != nil {
		if cur.Value == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return pseudoHead.Next
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

func (l *LinkedList) Del(value int) {
	if l.Head == nil {
		return
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next

		}
	}
}

//func (l *LinkedList) DelLink(value int) *LinkedList {
//	link := new(LinkedList)
//	if l.Head == nil {
//		return link
//	}
//	currentNode := l.Head
//	for currentNode.Next != nil {
//		if currentNode.Value == value {
//
//		}
//	}
//}
