package src_test

import "testing"

func sortedSquares(A []int) []int {
	ret := []int{}
	// 1. Find the edge of positive numbers and negatives
	edge := 0
	for _, val := range A {
		if val < 0 {
			edge++
			continue
		}
		break
	}

	// 2. Transformation => square
	for i, val := range A {
		A[i] = val * val
	}

	// 3. Merge [0 ... edge] and [edge ... A.size()-1]
	idx1 := edge - 1
	idx2 := edge
	for idx1 >= 0 && idx2 < len(A) {
		if A[idx1] <= A[idx2] {
			ret = append(ret, A[idx1])
			idx1--
		} else {
			ret = append(ret, A[idx2])
			idx2++
		}
	}

	for idx1 >= 0 {
		ret = append(ret, A[idx1])
		idx1--
	}

	for idx2 < len(A) {
		ret = append(ret, A[idx2])
		idx2++
	}

	return ret
}

func TestFunction(t *testing.T) {
	arr1 := []int{-4, -1, 0, 3, 10}
	arr2 := []int{1, 3, 5, 7, 9}
	arr3 := []int{-9, -7, -5, -3, -1}
	arr4 := []int{-7, -3, 2, 3, 11}

	t.Log(sortedSquares(arr1))
	t.Log(sortedSquares(arr2))
	t.Log(sortedSquares(arr3))
	t.Log(sortedSquares(arr4))
}
