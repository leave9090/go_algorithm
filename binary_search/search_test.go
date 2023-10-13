package binary_search

import (
	"fmt"
	"sort"
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

func Test_Sort(t *testing.T) {
	ss := []int{1, 4, 2, 7, 32, 5, 32, 3, 4, 23, 1}
	sort.Ints(ss)
	fmt.Println(ss)

}

func Test_mySqrt(t *testing.T) {

	fmt.Println(mySqrt(2))
}

func Test_searchBinary(t *testing.T) {
	fmt.Println(searchBinary([]int{1, 2, 3, 3, 3, 5, 5, 5, 6}, 3))
	fmt.Println(4 >> 2)
}
func searchBinary(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
