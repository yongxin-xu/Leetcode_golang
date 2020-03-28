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
		if nums[mid] > nums[low] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if nums[low] < nums[high] {
		return nums[low]
	} else {
		return nums[high]
	}
}

func TestFindMin(t *testing.T) {
	arr := []int{3, 4, 5, 1, 2}
	t.Log(findMin(arr))
}
