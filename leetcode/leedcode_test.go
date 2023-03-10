package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 3))
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	lengthOfLongestSubstring("abcdefasdffsdddsddd")
}

func Test_reverse(t *testing.T) {
	fmt.Println(reverse(137))
}

func Test_reverseString(t *testing.T) {
	s := []byte{'1', '2', '3'}
	reverseString(s)
}

func Test_myPow(t *testing.T) {
	fmt.Println(myPow(9, 99))
}
