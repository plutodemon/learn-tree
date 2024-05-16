package leetcode

// question_34
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	start := left
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target+1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	end := left - 1
	return []int{start, end}
}

// question_2529
func maximumCount(nums []int) int {
	small, large := 0, 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if num > 0 {
			large++
			continue
		}
		small++
	}
	if small > large {
		return small
	}
	return large
}
