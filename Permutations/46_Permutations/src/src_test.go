package src_test

import "testing"

func dfs(nums []int, used []bool, path []int, ans *[][]int) {
	if len(nums) == len(path) {
		oneAns := make([]int, len(path))
		copy(oneAns, path)
		*ans = append(*ans, oneAns)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			path = append(path, nums[i])
			dfs(nums, used, path, ans)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	used := make([]bool, len(nums))
	dfs(nums, used, nil, &ans)
	return ans
}

func TestPermutate(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	ans := permute(arr)
	t.Log(ans)
	t.Log(len(ans))
}
