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
