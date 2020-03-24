## Binary Search

Templates:

**Simple binary search**
```go
func binarySearch(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	ans := -1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			/* nums[mid] == target */
			return mid
		}
	}
	return ans
}
```

**Find the first/last target**
```go
type DirectionType int

const (
	first DirectionType = 1
	last  DirectionType = 2
)

func binarySearch(nums []int, target int, directionType DirectionType) int {
	high := len(nums) - 1
	low := 0
	ans := -1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			/* nums[mid] == target */
			if directionType == first {
				if (mid == 0) || nums[mid-1] != target {
					return mid
				} else {
					high = mid - 1
				}
			} else {
				if (mid == len(nums)-1) || nums[mid+1] != target {
					return mid
				} else {
					low = mid + 1
				}
			}
		}
	}
	return ans
}
```