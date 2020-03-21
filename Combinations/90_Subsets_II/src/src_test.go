package src_test

import (
	"sort"
	"testing"
)

func dfs(nums []int, subset []int, s int, ans *[][]int) {
	tmpVec := make([]int, len(subset))
	copy(tmpVec, subset)
	*ans = append(*ans, tmpVec)

	for i := s; i < len(nums); i++ {
		if i > s && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		dfs(nums, subset, i+1, ans)
		subset = subset[:len(subset)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, nil, 0, &ans)

	return ans
}

func TestSubsets(t *testing.T) {
	nums := []int{4, 4, 4, 1, 4}
	ans := subsetsWithDup(nums)
	t.Log(ans)
	t.Log(len(ans))
}
