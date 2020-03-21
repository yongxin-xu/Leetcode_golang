package src_test

import (
	"sort"
	"testing"
)

func dfs(candidates []int, target int, current_pos int, ans *[][]int, temp []int) {
	if target == 0 {
		tempAns := make([]int, len(temp), cap(temp))
		copy(tempAns, temp)
		*ans = append(*ans, tempAns)
		return
	}

	for i := current_pos; i < len(candidates); i++ {
		candidate := candidates[i]
		if candidate > target {
			break
		}
		temp = append(temp, candidate)
		dfs(candidates, target-candidate, i, ans, temp)
		temp = temp[:len(temp)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	if target == 0 || len(candidates) == 0 {
		return ans
	}
	dfs(candidates, target, 0, &ans, nil)
	return ans
}

func TestCombinationSumDFS(t *testing.T) {
	candidates := []int{8, 7, 4, 3}
	target := 11
	ans := combinationSum(candidates, target)
	t.Log(ans)
}
