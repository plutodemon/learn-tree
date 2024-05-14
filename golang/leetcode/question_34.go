package leetcode

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	mid := len(nums) / 2
	for left < right {
		if nums[mid] == target {

		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if left > right {
		return []int{-1, -1}
	}
	return []int{left, right}
}
