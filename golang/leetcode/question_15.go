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

// question_18
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ret := make([][]int, 0)
	for i := range nums {
		if i > n-4 {
			break
		}
		if i < n-4 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if j < n-5 && nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
					continue
				}
				if sum < target {
					left++
					continue
				}
				right--
			}
		}
	}
	return ret
}

// question_611
func TriangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	count := 0
	for i := range nums {
		if i > n-3 {
			break
		}
		if nums[i] == 0 {
			continue
		}
		k := i + 2
		for j := i + 1; j < n; j++ {
			for k < n && nums[i]+nums[j] > nums[k] {
				k++
			}
			count += k - j - 1
		}
	}
	return count
}

func TriangleNumber1(nums []int) int {
	sort.Ints(nums)
	count := 0
	for k := 2; k < len(nums); k++ {
		c := nums[k]
		i := 0
		j := k - 1
		for i < j {
			if nums[i]+nums[j] > c {
				count += j - i
				j--
				continue
			}
			i++
		}
	}
	return count
}

// question_11
func maxArea(height []int) int {
	area := 0
	left, right := 0, len(height)-1
	for left < right {
		sum := minNum(height[left], height[right]) * (right - left)
		if sum > area {
			area = sum
		}
		if height[left] < height[right] {
			left++
			continue
		}
		right--
	}
	return area
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// question_42 前后缀分解
func Trap(height []int) int {
	n := len(height)
	preMax, sufMax := make([]int, n), make([]int, n)
	preMax[0] = height[0]
	sufMax[n-1] = height[n-1]
	sum := 0
	for i := 1; i < n; i++ {
		if height[i] < preMax[i-1] {
			preMax[i] = preMax[i-1]
			continue
		}
		preMax[i] = height[i]
	}
	for i := n - 2; i >= 0; i-- {
		if height[i] < sufMax[i+1] {
			sufMax[i] = sufMax[i+1]
			continue
		}
		sufMax[i] = height[i]
	}
	for i := 0; i < n; i++ {
		sum += minNum(preMax[i], sufMax[i]) - height[i]
	}
	return sum
}

func Trap1(height []int) int {
	n := len(height)
	left, right := 0, n-1
	sum := 0
	preMax, sufMax := height[0], height[n-1]
	for left <= right {
		if preMax < height[left] {
			preMax = height[left]
		}
		if sufMax < height[right] {
			sufMax = height[right]
		}
		if preMax < sufMax {
			sum += preMax - height[left]
			left++
			continue
		}
		sum += sufMax - height[right]
		right--
	}
	return sum
}
