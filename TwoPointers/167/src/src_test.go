package src

import "testing"

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{-1, -1}
	}

	for i := 0; i < len(numbers); i++ {
		second := target - numbers[i]

		low := i + 1
		high := len(numbers) - 1

		for low <= high {
			mid := low + ((high - low) >> 1)
			if second < numbers[mid] {
				high = mid - 1
			} else if second > numbers[mid] {
				low = mid + 1
			} else {
				return []int{i + 1, mid + 1}
			}
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
