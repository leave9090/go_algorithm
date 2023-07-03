package array

import (
	"fmt"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 0, 1, 1, 1, 0, 1}))
}

func Test_thirdMax2(t *testing.T) {
	thirdMax2([]int{1, 2, 7, 2, 4, 9, 3, 4, 5})
}

func Test_twoNum(t *testing.T) {
	fmt.Println(twoNum([]int{2, 4, 5, 7, 8}, 6))
}
