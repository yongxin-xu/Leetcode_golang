package src_test

import "testing"

func backspaceCompare(S string, T string) bool {
	i := len(S) - 1
	j := len(T) - 1
	cnt1 := 0
	cnt2 := 0
	for i >= 0 && j >= 0 {
		for i >= 0 && (S[i] == '#' || cnt1 > 0) {
			if S[i] == '#' {
				cnt1++
			} else {
				cnt1--
			}
			i--
		}
		for j >= 0 && (T[j] == '#' || cnt2 > 0) {
			if T[j] == '#' {
				cnt2++
			} else {
				cnt2--
			}
			j--
		}
		if i < 0 || j < 0 {
			break
		}
		if S[i] != T[j] {
			return false
		}
		i--
		j--
	}
	for i >= 0 && (S[i] == '#' || cnt1 > 0) {
		if S[i] == '#' {
			cnt1++
		} else {
			cnt1--
		}
		i--
	}
	for j >= 0 && (T[j] == '#' || cnt2 > 0) {
		if T[j] == '#' {
			cnt2++
		} else {
			cnt2--
		}
		j--
	}
	return i == j
}

func TestExample(t *testing.T) {
	s1 := "ab#c"
	t1 := "ad#c"
	t.Log(backspaceCompare(s1, t1))
	s1 = "ab##"
	t1 = "c#d#"
	t.Log(backspaceCompare(s1, t1))
	s1 = "a##c"
	t1 = "#a#c"
	t.Log(backspaceCompare(s1, t1))
	s1 = "a#c"
	t1 = "b"
	t.Log(backspaceCompare(s1, t1))
	s1 = "bxj##tw"
	t1 = "bxo#j##tw"
	t.Log(backspaceCompare(s1, t1))
	s1 = "nzp#o#g"
	t1 = "b#nzp#o#g"
	t.Log(backspaceCompare(s1, t1))
}
