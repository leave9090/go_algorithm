package binary_search

import (
	"fmt"
	"testing"
)

func Test_searchRange(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 5, 5, 5, 6}
	target := 5
	fmt.Println(searchRange(nums, target))
}

func Test_searchCeil(t *testing.T) {

	fmt.Println(searchCeil([]int{1, 2, 3, 3, 3, 5, 5, 5, 6}, 3))
	fmt.Println(searchFloor([]int{1, 2, 3, 3, 3, 5, 5, 5, 6}, 3))

	left := 1
	right := 3
	fmt.Println(left + (right-left)>>1)
}

func Test_searchRange2(t *testing.T) {

	//fmt.Println(searchRange2([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9}, 3))
	//fmt.Println(searchRange3([]int{2, 5}, 5))
	fmt.Println(searchRange5([]int{1, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9}, 3))
}
