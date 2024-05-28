package leetcode

import (
	"sort"
)

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

// question_2300
func successfulPairs(spells []int, potions []int, success int64) []int {
	ret := make([]int, 0)
	pLen := len(potions)
	sort.Ints(potions)

	for _, spell := range spells {
		if spell == 0 {
			ret = append(ret, 0)
			continue
		}
		left := 0
		right := pLen - 1
		for left <= right {
			mid := (left + right) >> 1
			if int64(potions[mid])*int64(spell) >= success {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		ret = append(ret, pLen-left)
	}
	return ret
}

// question_2563
func CountFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	nLen := len(nums)

	if nLen == 1 ||
		nums[0]+nums[1] > upper ||
		nums[nLen-1]+nums[nLen-2] < lower {
		return 0
	}

	left := 0
	right := nLen - 1
	ret := int64(0)

	for left <= right {
		if nums[left]+nums[right] > upper {
			right--
			continue
		}
		if nums[left]+nums[right] < lower {
			left++
			continue
		}
		ret++

	}
	return ret
}
