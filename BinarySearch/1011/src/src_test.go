package src

import (
	"testing"
)

func sumOfArr(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func maxOfArr(arr []int) int {
	max := 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func shipOnTime(weights []int, D int, capacity int) bool {
	days := 1
	tempCap := capacity

	for _, weight := range weights {
		if weight > tempCap {
			days++
			tempCap = capacity - weight
		} else {
			tempCap -= weight
		}
	}

	return (days <= D)
}

func shipWithinDays(weights []int, D int) int {
	low := maxOfArr(weights)
	high := sumOfArr(weights)
	minDays := high

	for low <= high {
		mid := ((low - high) >> 1) + high
		if shipOnTime(weights, D, mid) {
			minDays = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return minDays
}

func TestFunction(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := 5
	t.Log(shipWithinDays(arr, d))
	arr2 := []int{3, 2, 2, 4, 1, 4}
	d2 := 3
	t.Log(shipWithinDays(arr2, d2))
	arr3 := []int{1, 2, 3, 1, 1}
	d3 := 4
	t.Log(shipWithinDays(arr3, d3))
}
