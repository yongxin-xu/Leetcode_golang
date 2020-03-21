package src_test

import "testing"

func dfs(k int, n int, s int, temp []int, ans *[][]int) {
	if n == 0 && k == 0 {
		tempAns := make([]int, len(temp))
		copy(tempAns, temp)
		*ans = append(*ans, tempAns)
		return
	}
	if n < s || s > 9 || k == 0 {
		return
	}
	temp = append(temp, s)
	dfs(k-1, n-s, s+1, temp, ans)
	temp = temp[:len(temp)-1]
	dfs(k, n, s+1, temp, ans)
}

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	dfs(k, n, 1, nil, &ans)
	return ans
}

func TestCombinationSum3(t *testing.T) {
	k := 3
	n := 24
	ans := combinationSum3(k, n)
	t.Log(ans)
}
