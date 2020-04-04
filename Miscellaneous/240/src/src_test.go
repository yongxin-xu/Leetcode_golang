package src_test

import "testing"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	ld_corner_x := 0
	ld_corner_y := len(matrix) - 1

	for ld_corner_x <= len(matrix[0])-1 &&
		ld_corner_y >= 0 {
		if matrix[ld_corner_y][ld_corner_x] == target {
			return true
		} else if matrix[ld_corner_y][ld_corner_x] > target {
			ld_corner_y--
		} else {
			ld_corner_x++
		}
	}

	return false
}

func TestExample1(t *testing.T) {
	m := [][]int{{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}

	t.Log(searchMatrix(m, 5))
	t.Log(searchMatrix(m, 20))
}
