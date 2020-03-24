package src_test

import "testing"

func search(nums []int, target int) int {
	ans := -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (high-low)>>1 + low
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			/* nums[mid] == target */
			return mid
		}
	}
	return ans
}

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	ans := search(nums, target)
	t.Log(ans)
}
