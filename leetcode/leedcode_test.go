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

func Test_numberOfPoints(t *testing.T) {
	var a = [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4}}
	numberOfPoints(a)
}

func subset(array []int, index int, data []int, i int, target int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(data[:i]))
		copy(temp, data[:i])
		*result = append(*result, temp)
		return
	}
	if target < array[index] {
		return
	}
	if index >= len(array) {
		return
	}
	data[i] = array[index]
	subset(array, index+1, data, i+1, target-array[index], result)
	subset(array, index+1, data, i, target, result)
}

func combinations(array []int, target int) [][]int {
	data := make([]int, len(array))
	result := [][]int{}
	subset(array, 0, data, 0, target, &result)
	return result
}

func TestDfs(t *testing.T) {
	array := []int{1, 2, 4, 5, 6, 7}
	target := 7
	res := combinations(array, target)
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestB(t *testing.T) {
	var a, b int
	a, b = 1, 5
	fmt.Println((a + b) / 2)
	fmt.Println(a + (b-a)/2)
}

func TestA(t *testing.T) {

}

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}

func Test_removeElement2(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement2(nums, 3))
	fmt.Println(nums)
}

func Test_maxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 100, 1, 5, 3, 6, 4}))
}

func Test_minSubArrayLen(t *testing.T) {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
