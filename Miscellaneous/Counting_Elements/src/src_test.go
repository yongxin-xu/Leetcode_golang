package src_test

import "testing"

func countElements(arr []int) int {
	cnt := 0
	m := map[int]int{}

	for _, val := range arr {
		m[val] = 1
	}

	for _, val := range arr {
		if v, ok := m[val+1]; ok {
			cnt += v
		}
	}

	return cnt
}

func TestExamples(t *testing.T) {
	t.Log(countElements([]int{1, 2, 3}))
	t.Log(countElements([]int{1, 1, 3, 3, 5, 5, 7, 7}))
	t.Log(countElements([]int{1, 3, 2, 3, 5, 0}))
	t.Log(countElements([]int{1, 1, 2, 2}))
}
