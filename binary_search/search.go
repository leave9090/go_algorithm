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
