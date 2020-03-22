package src_test

import (
	"sort"
	"testing"
)

func dfs(nums []int, used []bool, path []int, ans *[][]int) {
	if len(path) == len(nums) {
		oneAns := make([]int, len(path))
		copy(oneAns, path)
		*ans = append(*ans, oneAns)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(nums, used, path, ans)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	used := make([]bool, len(nums))
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, used, nil, &ans)
	return ans
}

func TestPermutate(t *testing.T) {
	arr := []int{1, 1, 2}
	ans := permuteUnique(arr)
	t.Log(ans)
	t.Log(len(ans))
}
