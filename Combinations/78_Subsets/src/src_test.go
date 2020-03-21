package src_test

import (
	"testing"
)

func dfs(nums []int, subset []int, s int, ans *[][]int) {
	tmpVec := make([]int, len(subset))
	copy(tmpVec, subset)
	*ans = append(*ans, tmpVec)

	for i := s; i < len(nums); i++ {
		subset = append(subset, nums[i])
		dfs(nums, subset, i+1, ans)
		subset = subset[:len(subset)-1]
	}
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	dfs(nums, nil, 0, &ans)

	return ans
}

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	ans := subsets(nums)
	t.Log(ans)
	t.Log(len(ans))
}
