package leetcode

import "strings"

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

// question_3
func LengthOfLongestSubstring(s string) int {
	count := 0
	left := 0
	for right := range s {
		substr := string(s[right])
		for strings.Contains(s[left:right], substr) {
			left++
			continue
		}
		lenRet := len(s[left : right+1])
		if lenRet > count {
			count = lenRet
		}
	}
	return count
}

// question_862
func ShortestSubarray(nums []int, k int) int {
	n := len(nums)
	minLen := n + 1
	left := 0
	sum := 0
	less := 0
	for right, num := range nums {
		if num >= k {
			return 1
		}
		if num < 0 {
			less++
		}
		sum += num
		if sum <= 0 {
			left = right
			sum = num
			if num < 0 {
				less = 1
			}
			continue
		}
		oldLeft := left
		oldLess := less
		oldSum := sum
		for oldLeft <= right && (oldSum >= k || oldLess > 0) {
			if right-oldLeft+1 < minLen && oldSum >= k {
				minLen = right - oldLeft + 1
				less = oldLess
				left = oldLeft
				sum = oldSum
			}
			leftNum := nums[oldLeft]
			if leftNum < 0 {
				oldLess--
			}
			oldSum -= leftNum
			oldLeft++
		}
	}
	if minLen > n {
		minLen = -1
	}
	return minLen
}
