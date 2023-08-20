package binary_search

func searchRange(nums []int, target int) [2]int {
	result := [2]int{-1, -1}
	leftIdx := binarySearch(nums, target, true)

	if leftIdx == len(nums) || nums[leftIdx] != target {
		return result
	}

	result[0] = leftIdx
	result[1] = binarySearch(nums, target, false) - 1
	return result
}

func binarySearch(nums []int, target int, findFirst bool) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target || (findFirst && nums[mid] == target) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

/**********************************************************************************/
func searchFloor(arr []int, tar int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] >= tar {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l < len(arr) && arr[l] == tar {
		return l
	}
	return -1
}

func searchCeil(arr []int, tar int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] <= tar {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if r >= 0 && arr[r] == tar {
		return r
	}
	return -1
}

/**********************************************************************************/
/*
704 左闭右闭
*/
func searchRange2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)>>1
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

/*
左闭右开
*/
func searchRange3(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)>>1
		if target < nums[middle] {
			right = middle
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

/**********************************************************************************/

/*
34
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/
/*
方法1
*/
func searchRange4(nums []int, target int) []int {
	left := leftRange(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := leftRange(nums, target+1) - 1
	return []int{left, right}
}

func leftRange(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)>>1
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		}
	}
	return left
}

/*
方法2
*/
func leftRangeFunc(nums []int, target int) int {
	left, right := 0, len(nums)-1
	start := -1
	for left <= right {
		middle := left + (right-left)>>1
		if target <= nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
		if nums[middle] == target {
			start = middle
		}
	}
	return start
}
func rightRangeFunc(nums []int, target int) int {
	left, right := 0, len(nums)-1
	end := -1
	for left <= right {
		middle := left + (right-left)>>1
		if target >= nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
		if nums[middle] == target {
			end = middle
		}
	}
	return end
}
func searchRange5(nums []int, target int) []int {
	return []int{leftRangeFunc(nums, target), rightRangeFunc(nums, target)}
}

/**********************************************************************************/

/*
35
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。
*/
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)>>1
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			return middle
		}
	}
	return left
}

/**********************************************************************************/
/*
540. 有序数组中的单一元素
中等
给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
请你找出并返回只出现一次的那个数。
你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

 1,1,2,3,3
 1,2,2,3,3
 1,1,2,2,3
*/
func singleNonDuplicate(nums []int) int {

	return 0
}

/**********************************************************************************/
/*
给你一个非负整数 x ，计算并返回x的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/
func mySqrt(x int) int {
	left, right := 0, x
	end := -1
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid == x {
			return mid
		} else {
			end = mid
			left = mid + 1
		}
	}
	return end
}
