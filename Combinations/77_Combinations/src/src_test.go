package src_test

import "testing"

func dfs(n int, numbers_left int, s int,
	visited []bool, temp []int, ans *[][]int) {
	if numbers_left == 0 {
		temp_ans := make([]int, len(temp), cap(temp))
		copy(temp_ans, temp)
		*ans = append(*ans, temp_ans)
		return
	}
	for i := s; i <= n; i++ {
		if visited[i] {
			continue
		}
		temp = append(temp, i)
		visited[i] = true
		dfs(n, numbers_left-1, i+1, visited, temp, ans)
		visited[i] = false
		temp = temp[:len(temp)-1]
	}
}

func combine(n int, k int) [][]int {
	ans := [][]int{}
	if n < k || n <= 0 || k <= 0 {
		return ans
	}
	visited := make([]bool, n+1)
	dfs(n, k, 1, visited, nil, &ans)
	return ans
}

func TestCombine(t *testing.T) {
	ans := combine(5, 3)
	t.Log(ans)
}
