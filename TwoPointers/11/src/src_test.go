package src_test

import "testing"

func minInt(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func calculateArea(height []int, left int, right int) int {
	return minInt(height[left], height[right]) * (right - left)
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxAreaVal := 0
	for left < right {
		area := calculateArea(height, left, right)
		if area > maxAreaVal {
			maxAreaVal = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxAreaVal
}

func TestMaxArea(t *testing.T) {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(arr))
}
