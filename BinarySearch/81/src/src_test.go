package src_test

import "testing"

func search(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low-high)/2 + high
		if nums[mid] == target {
			return true
		}
		if nums[low] == nums[mid] {
			low++
			continue
		}
		if nums[low] < nums[mid] {
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

func TestSearch(t *testing.T) {
	nums := []int{1, 1, 1, 3, 1}
	target := 3
	t.Log(search(nums, target))
}
