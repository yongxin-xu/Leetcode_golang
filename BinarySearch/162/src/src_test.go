package src

import "testing"

func findPeakElement(nums []int) int {
	arrLen := len(nums)
	if arrLen <= 1 || nums[0] > nums[1] {
		return 0
	}
	if nums[arrLen-1] > nums[arrLen-2] {
		return arrLen - 1
	}
	low := 0
	high := arrLen - 1
	for low <= high {
		mid := ((low - high) >> 1) + high
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] < nums[mid+1] {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}

func TestFindPeakElement(t *testing.T) {
	arr := []int{1, 2, 1, 3, 5, 6, 4}
	t.Log(findPeakElement(arr))
}
