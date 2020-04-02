package src_test

import "testing"

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}

	m := make(map[int]bool)
	m[n] = true

	for {
		val := 0
		temp := n
		for temp > 0 {
			v := temp % 10
			temp /= 10
			val = val + v*v
		}

		if val == 1 {
			return true
		}

		_, ok := m[val]

		if ok {
			break
		} else {
			m[val] = true
		}
		n = val
	}

	return false
}

func TestIsHappy(t *testing.T) {
	t.Log(isHappy(19))
}
