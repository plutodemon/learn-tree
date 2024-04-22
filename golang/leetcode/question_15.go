package leetcode

import (
	"math"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ret := make([][]int, 0)
	for i := range nums {
		if i > n-2 {
			break
		}
		// 优化：nums[i]+nums[i+1]+nums[i+2]大于0，后面的数都大于0
		if i < n-2 && nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if (i > 0 && nums[i] == nums[i-1]) ||
			// 优化：nums[i]+nums[len(nums)-1]+nums[len(nums)-2]小于0，后面的数都小于0
			nums[i]+nums[n-1]+nums[n-2] < 0 {
			continue
		}
		target := nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := target + nums[left] + nums[right]
			if sum == 0 {
				ret = append(ret, []int{target, nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
				continue
			}
			if sum < 0 {
				left++
				continue
			}
			right--
		}
	}
	return ret
}

// question_167
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
			continue
		}
		right--
	}
	return []int{}
}

// question_2824
func countPairs(nums []int, target int) int {
	left, right, count := 0, len(nums)-1, 0
	sort.Ints(nums)
	for left < right {
		if nums[left]+nums[right] < target {
			count += right - left
			left++
			continue
		}
		right--
	}
	return count
}

// question_16
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	minDiff := math.MaxInt
	for i := range nums {
		if i > n-3 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 优化一
		s := nums[i] + nums[i+1] + nums[i+2]
		if s > target { // 后面无论怎么选，选出的三个数的和不会比 s 还小
			if s-target < minDiff {
				sum = s // 由于下面直接 break，这里无需更新 minDiff
			}
			break
		}
		// 优化二
		s = nums[i] + nums[n-2] + nums[n-1]
		if s < target { // x 加上后面任意两个数都不超过 s，所以下面的双指针就不需要跑了
			if target-s < minDiff {
				minDiff = target - s
				sum = s
			}
			continue
		}
		left, right := i+1, n-1
		for left < right {
			s = nums[i] + nums[left] + nums[right]
			if s == target {
				return target
			}
			if s > target {
				if s-target < minDiff { // s 与 target 更近
					minDiff = s - target
					sum = s
				}
				right--
			} else { // s < target
				if target-s < minDiff { // s 与 target 更近
					minDiff = target - s
					sum = s
				}
				left++
			}
		}
	}
	return sum
}
