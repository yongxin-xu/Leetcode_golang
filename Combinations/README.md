## Templates for Combinations
The basic idea is using depth-first search.

```go
/**
[in]    nums    candidate numbers
[in]    d       depth of recursion
[in]    n       how many number to use
[in]    s       starting index
[in]    temp    a potential solution of answer
*/
func Combination(nums []int, d int, n int, s int, temp []int, ans *[][]int) {
    if d == n:
        tempAns := make([]int, len(temp))
        copy(tempAns, temp)
        *ans = append(*ans, tempAns)
    
    for i := s; i < len(nums), i++ {
        temp = append(temp, nums[i])
        Combination(nums, d+1, n, i+1, temp, ans)
        temp = temp[:len(temp)-1]
    }
}
```