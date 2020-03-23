package src_test

import (
	"fmt"
	"testing"
)

func merge_strings(str1 []byte, str2 []byte) string {
	minLen := myMin(len(str1), len(str2))
	for i := 0; i < minLen; i++ {
		tempStr1 := string(str1[len(str1)-minLen+i:])
		tempStr2 := string(str2[:minLen-i])
		if tempStr1 == tempStr2 {
			return string(str1) + string(str2[minLen-i:])
		}
	}
	return string(str1) + string(str2)
}

func dfs(input []string, path []string, used []bool, ans *[][]string) {
	if len(path) == len(input) {
		tempAns := make([]string, len(path))
		copy(tempAns, path)
		*ans = append(*ans, tempAns)
	}

	for i := 0; i < len(input); i++ {
		if !used[i] {
			used[i] = true
			path = append(path, input[i])
			dfs(input, path, used, ans)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

func permute(input []string, ans *[][]string) string {
	used := make([]bool, len(input))
	dfs(input, nil, used, ans)

	var shortestSuper string = ""

	for _, strArray := range *ans {
		var temp string = ""
		for _, substr := range strArray {
			temp = merge_strings([]byte(temp), []byte(substr))
		}
		if len(shortestSuper) == 0 {
			shortestSuper = temp
		} else if len(temp) < len(shortestSuper) {
			shortestSuper = temp
		}
	}

	return shortestSuper
}

func shortestSuperstring(A []string) string {
	permutations := [][]string{}
	ans := permute(A, &permutations)
	return ans
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
	fmt.Println(answer == "gctaagttcatgcatc")
}

func TestMergeStrings(t *testing.T) {
	fmt.Println(merge_strings([]byte("abcde"), []byte("cdet")))
	fmt.Println(merge_strings([]byte("bcde"), []byte("cdeat")))
}
