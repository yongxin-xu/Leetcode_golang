package src_test

import (
	"testing"
	"unicode"
)

func stringRemvoeLast(str string) string {
	temp := []byte(str)
	temp = temp[:len(temp)-1]
	return string(temp)
}

func stringAppendLast(str string, c byte) string {
	temp := []byte(str)
	temp = append(temp, c)
	return string(temp)
}

func dfs(str string, substr string, cur int, ans *[]string) {
	if cur == len(str) {
		*ans = append(*ans, substr)
		return
	}

	if unicode.IsLetter(rune(str[cur])) {
		lowwer_char := unicode.ToLower(rune(str[cur]))
		substr = stringAppendLast(substr, byte(lowwer_char))
		dfs(str, substr, cur+1, ans)

		upper_char := unicode.ToUpper(rune(str[cur]))
		substr = stringRemvoeLast(substr)
		substr = stringAppendLast(substr, byte(upper_char))
		dfs(str, substr, cur+1, ans)
	} else {
		substr = stringAppendLast(substr, str[cur])
		dfs(str, substr, cur+1, ans)
	}
}

func letterCasePermutation(S string) []string {
	ans := []string{}
	dfs(S, "", 0, &ans)
	return ans
}

func TestLetterCasePermutation(t *testing.T) {
	str := "a1b2"
	arr := letterCasePermutation(str)
	t.Log(arr)
	t.Log(len(arr))
}

func TestAppend(t *testing.T) {
	str := "a1b2"
	str = stringAppendLast(str, '3')
	t.Log(str)

}
