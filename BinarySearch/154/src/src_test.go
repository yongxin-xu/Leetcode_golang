package src_test

import "testing"

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high-1 {
		if nums[low] < nums[high] {
			break
		}

		mid := ((low - high) >> 1) + high
		if nums[mid] == nums[low] {
			low++
		} else if nums[mid] < nums[low] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if nums[low] <= nums[high] {
		return nums[low]
	} else {
		return nums[high]
	}
}

func TestFindMin(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	t.Log(findMin(arr))
}
