package src_test

import (
	"fmt"
	"math"
	"testing"
)

func dfs(input []string, cur int, used []bool, curLen int, bestLen int, path []int, bestPath *[]int,
	costMatrix *[][]int) {
	if len(path) >= bestLen {
		return
	}
	if len(path) == len(input) {
		bestLen = curLen
		*bestPath = path
		return
	}
	for i := 0; i < len(input); i++ {
		if !used[i] {
			used[i] = true
			path = append(path, i)
			if cur == 0 {
				dfs(input, cur+1, used, len(input[i]), bestLen, path, bestPath, costMatrix)
			} else {
				cost := (*costMatrix)[path[i-1]][i]
				dfs(input, cur+1, used, curLen+cost, bestLen, path, bestPath, costMatrix)
			}
			used[i] = false
		}
	}
}

func shortestSuperstring(A []string) []int {
	used := make([]bool, len(A))
	costMatrix := preprocess(A)
	ans := []int{}
	dfs(A, 0, used, 0, math.MaxInt32, nil, &ans, &costMatrix)
	return ans
}

func preprocess(A []string) [][]int {
	costMatrix := [][]int{}
	for i := 0; i < len(A); i++ {
		temp := []int{}
		for j := 0; j < len(A); j++ {
			hasCommon := false
			minLen := myMin(len(A[i]), len(A[j]))
			for idx := 0; idx < minLen; idx++ {
				tempStr1 := string(A[i][len(A[idx])-minLen+idx:])
				tempStr2 := string(A[j][:minLen-idx])
				if tempStr1 == tempStr2 {
					temp = append(temp, len(A[j])-idx)
				}
			}
			if !hasCommon {
				temp = append(temp, len(A[j]))
			}
		}
		costMatrix = append(costMatrix, temp)
	}
	return costMatrix
}

func myMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestShortestSuperstring(t *testing.T) {
	input := []string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}
	answer := shortestSuperstring(input)
	fmt.Println(answer)
}
