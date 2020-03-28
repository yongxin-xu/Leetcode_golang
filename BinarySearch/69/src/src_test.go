package src_test

import "testing"

func mySqrt(x int) int {
	low := 1
	high := x
	for low <= high {
		mid := ((low - high) >> 1) + high
		if mid > x/mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

func TestMySqrt(t *testing.T) {
	x := 17
	t.Log(mySqrt(x))
}
