package src_test

import "testing"

func peakIndexInMountainArray(A []int) int {
	low := 0
	high := len(A) - 1

	for low <= high {
		mid := ((low - high) >> 1) + high
		// the input guarantees that 3 <= A.length <= 10000

		if (A[mid-1] < A[mid]) && (A[mid] < A[mid+1]) {
			/* 1st. A[mid-1] < A[mid] < A[mid+1] */
			low = mid + 1
		} else if (A[mid-1] > A[mid]) && (A[mid] > A[mid+1]) {
			/* 2nd. A[mid-1] > A[mid] > A[mid+1] */
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func TestFunction(t *testing.T) {
	arr := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	peakIndex := peakIndexInMountainArray(arr)
	t.Log(peakIndex)
}
