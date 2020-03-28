package src

import "testing"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	low := 0
	high := len(matrix)*len(matrix[0]) - 1

	numCol := len(matrix[0])

	for low <= high {
		mid := ((low - high) >> 1) + high
		rowNo := mid / numCol
		colNo := mid % numCol
		if matrix[rowNo][colNo] == target {
			return true
		} else if matrix[rowNo][colNo] < target {
			low = mid + 1
		} else /* matrix[rowNo][colNo] > target */ {
			high = mid - 1
		}
	}

	return false
}

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	t.Log(searchMatrix(matrix, 13))
	t.Log(searchMatrix(matrix, 9))
	t.Log(searchMatrix(matrix, 3))
	t.Log(searchMatrix(matrix, 1))
	t.Log(searchMatrix(matrix, 5))
	t.Log(searchMatrix(matrix, 7))
	t.Log(searchMatrix(matrix, 10))
	t.Log(searchMatrix(matrix, 11))
	t.Log(searchMatrix(matrix, 16))
	t.Log(searchMatrix(matrix, 20))
	t.Log(searchMatrix(matrix, 23))
	t.Log(searchMatrix(matrix, 30))
	t.Log(searchMatrix(matrix, 34))
	t.Log(searchMatrix(matrix, 50))
	// matrix := [][]int{{1}}
	// t.Log(searchMatrix(matrix, 2))
}
