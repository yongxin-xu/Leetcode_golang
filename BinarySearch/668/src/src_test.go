package src_test

import "testing"

func rankInMatrix(matrix [][]int, val int) int {
	cnt := 0

	i := 0
	for j := len(matrix) - 1; j >= 0 && i <= len(matrix)-1; {
		if matrix[i][j] <= val {
			cnt += (j + 1)
			i++
		} else {
			j--
		}
	}

	return cnt
}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}
	n := len(matrix)
	low := matrix[0][0]
	high := matrix[n-1][n-1]

	// k is always valid
	for low < high {
		mid := ((low - high) >> 1) + high
		if rankInMatrix(matrix, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func TestFunction(t *testing.T) {
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	t.Log(kthSmallest(matrix, 8))
}
