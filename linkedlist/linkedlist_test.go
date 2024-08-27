package linkedlist

import (
	"fmt"
	"testing"
)

func Test_Linked(t *testing.T) {
	linkedList := &LinkedList{}
	linkedList.TailAdd(1)
	linkedList.TailAdd(1)
	linkedList.TailAdd(1)

	linkedList.TailAdd(2)
	linkedList.TailAdd(3)
	linkedList.TailAdd(1)

	linkedList.TailAdd(4)
	linkedList.TailAdd(5)

	linkedList.Print()
	fmt.Println("=====>", linkedList.Length())

	//linkedList.delete(1)
	linkedList.Print()

}

func Test_removeElements(t *testing.T) {
	linkedList := &LinkedList{}
	linkedList.TailAdd(1)
	linkedList.TailAdd(1)
	linkedList.TailAdd(1)

	linkedList.TailAdd(2)
	linkedList.TailAdd(3)
	linkedList.TailAdd(1)

	linkedList.TailAdd(4)
	linkedList.TailAdd(5)

	linkedList.Print()
	fmt.Println("=====>", linkedList.Length())
	linkedList.removeElements(1)
	//linkedList.Head = removeElements(linkedList.Head, 1)
	linkedList.Print()
	fmt.Println("=====>", linkedList.Length())

}

func Test_removeNthFromEnd(t *testing.T) {
	//linkedList := &LinkedList{}
	//linkedList.TailAdd(1)
	//linkedList.TailAdd(2)
	//linkedList.TailAdd(3)
	//
	//linkedList.TailAdd(4)
	//linkedList.TailAdd(5)
	//linkedList.TailAdd(6)
	//
	//linkedList.TailAdd(7)
	//linkedList.TailAdd(8)
	//linkedList.Print()
	//head := removeNthFromEnd(linkedList.Head, 2)
	newLink()
	head := removeNthFromEnd(newLink(), 2)

	currentNode := head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}

}

// 创建一个链表 1->2->3->4->5
func newLink() *ListNode {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	return head
}

// 打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Value, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}
