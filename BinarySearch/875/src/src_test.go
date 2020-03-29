package src_test

import "testing"

func maxElem(arr []int) int {
	/* arr is a non-negative array */
	maximum := 0
	for _, val := range arr {
		if val > maximum {
			maximum = val
		}
	}
	return maximum
}

func finishOnTime(piles []int, H int, speed int) bool {
	hoursTaken := 0
	for _, num := range piles {
		hoursTaken += ((num + speed - 1) / speed)
	}
	return (hoursTaken <= H)
}

func minEatingSpeed(piles []int, H int) int {
	low := 1
	high := maxElem(piles)
	minSpeed := high

	for low <= high {
		mid := ((low - high) >> 1) + high
		if finishOnTime(piles, H, mid) {
			minSpeed = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return minSpeed
}

func TestFunction(t *testing.T) {
	arr := []int{30, 11, 23, 4, 20}
	n := 5
	t.Log(minEatingSpeed(arr, n))
}
