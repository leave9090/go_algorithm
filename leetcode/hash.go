package leetcode

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

//3. 无重复字符的最长子串
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
