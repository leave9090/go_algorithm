package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Linked(t *testing.T) {
	linkedList := &LinkedList{}
	linkedList.TailAdd(1)
	linkedList.TailAdd(2)
	linkedList.TailAdd(3)

	linkedList.Print()
	fmt.Println("=====>", linkedList.Length())

	linkedList.Del(2)
	linkedList.Print()

}

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
