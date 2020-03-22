## Permutations
Basically, there are two ways to implement a permutation.

The first is to enumerate numbers for a position
The second is to enumerate positions for a number

Two templates

1. Enumerate numbers for a postition
```go
func dfs(nums []int, used []bool, path []int, cur int ans *[][]int) {
	if cur == len(nums) {
		oneAns := make([]int, len(path))
		copy(oneAns, path)
		*ans = append(*ans, oneAns)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			path[i] = nums[i]
			dfs(nums, used, path, cur+1, ans)
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	ans := [][]int{}
    used := make([]bool, len(nums))
    path := make([]int, len(nums))
	dfs(nums, used, path, 0, &ans)
	return ans
}
```

2. Enumerate positions for a number
```go
func dfs(nums []int, used []bool, path []int, ans *[][]int) {
	if len(nums) == len(path) {
		oneAns := make([]int, len(path))
		copy(oneAns, path)
		*ans = append(*ans, oneAns)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			path = append(path, nums[i])
			dfs(nums, used, path, ans)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	used := make([]bool, len(nums))
	dfs(nums, used, nil, &ans)
	return ans
}
```