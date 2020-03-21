package src_test

import (
	"sort"
	"testing"
)

func dfs(candidates []int, target int, current_pos int,
	ans *[][]int, temp []int) {
	if target == 0 {
		temp_ans := make([]int, len(temp), cap(temp))
		copy(temp_ans, temp)
		*ans = append(*ans, temp_ans)
		return
	}

	for i := current_pos; i < len(candidates); i++ {
		candidate := candidates[i]
		if target < candidate {
			return
		}
		if i > current_pos && candidates[i] == candidates[i-1] {
			continue
		}
		temp = append(temp, candidate)
		dfs(candidates, target-candidate, i+1, ans, temp)
		temp = temp[:len(temp)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	if len(candidates) == 0 {
		return ans
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	temp := []int{}
	dfs(candidates, target, 0, &ans, temp)
	return ans
}

func TestCombinationSumDFS(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	ans := combinationSum2(candidates, target)
	t.Log(ans)
}
