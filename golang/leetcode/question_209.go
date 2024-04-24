package leetcode

// question_209
func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	start, end := 0, 0
	minLen := n + 1
	sum := 0
	for end <= n {
		if sum >= target {
			if (end - start) <= minLen {
				minLen = end - start
			}
			sum -= nums[start]
			start++
			continue
		}
		if end < n {
			sum += nums[end]
		}
		end++
	}
	if minLen > n {
		minLen = 0
	}
	return minLen
}

func MinSubArrayLen1(target int, nums []int) int {
	n, sum, left := len(nums), 0, 0
	minLen := n + 1
	for right := range nums {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if minLen > n {
		minLen = 0
	}
	return minLen
}

// question_713
func NumSubarrayProductLessThanK(nums []int, k int) int {
	sum, count, left := 1, 0, 0
	for right := range nums {
		sum *= nums[right]
		for ; left <= right && sum >= k; left++ {
			sum /= nums[left]
		}
		count += right - left + 1
	}
	return count
}
