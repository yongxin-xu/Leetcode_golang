package src_test

import "testing"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low-high)/2 + high /* mid = (low+high)/2 */
		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
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
	return -1
}

func TestSearch(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	t.Log(search(nums, target))

	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	t.Log(search(nums2, target2))

	nums2 = []int{1, 2, 3, 5, 6, 7, 8}
	target2 = 4
	t.Log(search(nums2, target2))

	nums2 = []int{1, 2, 3, 5, 6, 7, 8}
	target2 = 7
	t.Log(search(nums2, target2))

	nums2 = []int{8, 1, 2, 3, 5, 6, 7}
	for i := 0; i <= 10; i++ {
		target2 = i
		t.Logf("%d: %d", target2, search(nums2, target2))
	}
}
