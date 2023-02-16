package leetcode

/*
70
*/
var s = map[int]int{}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	if i, ok := s[n]; !ok {
		var num = climbStairs(n-1) + climbStairs(n-2)
		s[n] = num
		return num
	} else {
		return i
	}

}
