package src_test

import "testing"

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := map[string]([]string){}

	for _, str := range strs {
		alphaCnt := [26]int{}
		// calculate how many times each character occurs
		for _, ch := range str {
			alphaCnt[(int)(ch-'a')]++
		}

		var temp string
		// store the pattern, e.g. "moon" => "m1n1o2"
		for i, cnt := range alphaCnt {
			if cnt != 0 {
				temp = temp + string(i+'a') + string(cnt)
			}
		}
		m[temp] = append(m[temp], str)
	}

	for _, strArr := range m {
		res = append(res, strArr)
	}
	return res
}

func TestFunction(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(groupAnagrams(input))
}
