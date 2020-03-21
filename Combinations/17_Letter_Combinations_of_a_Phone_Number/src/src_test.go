package src_test

import (
	"testing"
)

// URL: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func dfs(digits string, dictionary *[8][]byte,
	l int, cur string, ans *[]string) {
	if l == len(digits) {
		*ans = append(*ans, cur)
		return
	}
	for _, ch := range dictionary[digits[l]-'2'] {
		temp := []byte(cur)
		temp = append(temp, ch)
		cur = string(temp)
		dfs(digits, dictionary, l+1, cur, ans)
		temp = temp[:len(temp)-1]
		cur = string(temp)
	}
}

func bfs(digits string, dictionary *[8][]byte) []string {

	ans := []string{}
	if digits == "" {
		return ans
	}
	ans = append(ans, "")
	for _, digit := range digits {
		tmp := []string{}
		for _, ch := range dictionary[digit-'2'] {
			for _, s := range ans {
				strTemp := []byte(s)
				strTemp = append(strTemp, byte(ch))
				s = string(strTemp)
				tmp = append(tmp, s)
			}
		}
		ans, tmp = tmp, ans
	}

	return ans
}

func letterCombinations(digits string) []string {
	dictionary := [8][]byte{
		{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'},
		{'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'},
		{'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}

	ans := []string{}
	if digits == "" {
		return ans
	}
	cur := ""
	dfs(digits, &dictionary, 0, cur, &ans)
	return ans
}

func letterCombinations2(digits string) []string {
	dictionary := [8][]byte{
		{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'},
		{'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'},
		{'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}

	ans := bfs(digits, &dictionary)
	return ans
}

func TestLetterCombinationsDFS(t *testing.T) {
	input := "23"
	ret := letterCombinations(input)
	t.Log(ret)
}

func TestLetterCombinationsBFS(t *testing.T) {
	input := "23"
	ret := letterCombinations2(input)
	t.Log(ret)
}
