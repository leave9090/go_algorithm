package leetcode

import (
	"fmt"
	"math"
)

// 1.两数之和
func twoSum(data []int, num int) []int {
	numMap := map[int]int{}
	for i, value := range data {
		if v, ok := numMap[num-value]; ok {
			return []int{v, i}
		}
		numMap[value] = i
	}
	return nil
}

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	var num int

	return num
}

func maxNum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

/*
7. 整数反转  137
*/
func reverse(x int) int {
	s := 0

	fmt.Println(math.MinInt32, math.MaxInt32)
	for x != 0 {
		if (s < math.MinInt32/10) || (s > math.MaxInt32/10) {
			return 0
		}
		l := x % 10
		x /= 10
		s = s*10 + l
	}

	return s
}

/*
字符串反转 344
输入：s = ["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
*/
func reverseString(s []byte) {
	n := len(s)
	for i, j := 0, n-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	fmt.Println(string(s))
}

/*
链表反转 206
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	return head
}

/*
2. 两数相加
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var list *ListNode

	return list
}

/*
	2  2

2*2
2 3
2*2*2
*/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 || x == 1.0 {
		return x
	}

	if n < 0 {
		n *= -1
		x = 1.0 / x
	}

	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)

}

/*
2848.与车相交的点
*/
func numberOfPoints(nums [][]int) int {
	for _, i := range nums {
		fmt.Println(i)
	}
	return 0
}

/*
 */
func removeElement(nums []int, val int) int {
	var num int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[num] = nums[i]
			num++
		}
	}
	nums = nums[:num]
	return num
}

/*
双指针
*/
func removeElement2(nums []int, val int) int {

	var left, right int
	left, right = 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	nums = nums[:left]
	return left

}

/*
121  贪心算法
*/
func maxProfit(prices []int) int {
	min := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > res {
			res = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return res
}

/*
209.长度最小的子数组
暴力
*/
func minSubArrayLen(target int, nums []int) int {
	var length int
	var result = math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				length = j - i + 1
				if result > length {
					result = length
					break
				}
			}
		}
	}
	if result == math.MaxInt32 {
		result = 0
	}
	return result

}

/*
209.长度最小的子数组
滑动窗口
*/
func minSubArrayLen2(target int, nums []int) int {
	i := 0 //起始位置
	sum := 0
	var result = math.MaxInt32
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			lens := j - i + 1
			if lens < result {
				result = lens
			}

			sum -= nums[i]
			i++
		}

	}
	if result == math.MaxInt32 {
		return 0
	} else {
		return result
	}
}

func reverseString2(s []byte) {
	var i, j int
	j = len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

/*
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	var left, right int
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// 斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func fib3(n int) int {
	num := make(map[int]int)

	return helper(n, num)
}
func helper(n int, num map[int]int) int {
	if n < 2 {
		return n
	}
	if v, ok := num[n]; ok {
		return v
	}
	num[n] = helper(n-1, num) + helper(n-2, num)
	return num[n]
}

// 不同路径
func uniquePaths(m int, n int) int {
	l := make([][]int, m)

	for i := range l {
		l[i] = make([]int, n)
		l[i][0] = 1
	}
	for i := 0; i < n; i++ {
		l[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			l[i][j] = l[i-1][j] + l[i][j-1]
		}
	}
	return l[m-1][n-1]
}
