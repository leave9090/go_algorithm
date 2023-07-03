package array

import (
	"fmt"
	"sort"
)

/*
1. two num
*/
func twoNum(nums []int, num int) (int, int) {
	numMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := numMap[num-v]; ok {
			return i, j
		}
		numMap[v] = i
	}
	return 0, 0
}

// 485. 最大连续 1 的个数
func findMaxConsecutiveOnes(nums []int) int {
	var numMax, numNow int

	for _, v := range nums {
		if v == 0 {
			numNow = 0
		} else {
			numNow++
		}
		numMax = maxNUm(numMax, numNow)
	}
	return numMax
}

/*
 495. 提莫攻击

输入：timeSeries = [1,2], duration = 2
输出：3

	1 2 3 5
	3
	1 2 3    1 2 3  3
	2 3 4        4  1
	3 4 5        5  1
	5 6 7      6 7  2
*/
func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	expired := 0
	for _, t := range timeSeries {

		if t >= expired {
			ans += duration
		} else {
			ans += t + duration - expired
		}

		expired = t + duration
	}

	return ans
}

/*
414. 第三大的数
*/
func thirdMax(nums []int) int {
	num := 0
	numMap := make(map[int]bool)
	for _, v := range nums {
		numMap[v] = true
	}
	numKey := make([]int, 0)
	for key, _ := range numMap {
		numKey = append(numKey, key)
	}
	sort.Ints(numKey)
	if len(numKey) >= 3 {
		num = numKey[len(numKey)-3]
	} else {
		num = numKey[len(numKey)-1]
	}
	return num
}
func thirdMax2(nums []int) int {
	num := 0
	//sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(sort.IntSlice(nums))

	//fmt.Println(sort.Reverse(sort.IntSlice(nums)))

	return num
}

func maxNUm(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
