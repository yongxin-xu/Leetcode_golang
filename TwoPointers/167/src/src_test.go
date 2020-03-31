package src

import "testing"

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{-1, -1}
	}

	idx1 := 0
	idx2 := len(numbers) - 1
	for idx1 < idx2 {
		if numbers[idx1]+numbers[idx2] == target {
			return []int{idx1 + 1, idx2 + 1}
		} else if numbers[idx1]+numbers[idx2] > target {
			idx2--
		} else {
			idx1++
		}
	}

	return []int{-1, -1}
}

func TestTwoSum(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	t.Log(twoSum(arr, 9))
	t.Log(twoSum(arr, 13))
	t.Log(twoSum(arr, 17))
	t.Log(twoSum(arr, 18))
	t.Log(twoSum(arr, 22))
	t.Log(twoSum(arr, 26))
}
