package src_test

import (
	"testing"
)

func maxSubArray(nums []int) int {
	ret := nums[0]
	subsum := nums[0]

	// maxSub(nums[0:i+1]) = max(maxSub(nums[0:i]), nums[i+1])

	for i := 1; i < len(nums); i++ {
		if subsum > 0 {
			subsum += nums[i]
		} else {
			subsum = nums[i]
		}

		if subsum > ret {
			ret = subsum
		}
	}
	return ret
}

func TestFunction(t *testing.T) {
	t.Log(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
