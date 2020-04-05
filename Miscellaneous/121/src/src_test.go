package src_test

import "testing"

func maxProfit(prices []int) int {
	maxProf := 0
	if len(prices) > 1 {
		currentMin := prices[0]
		for _, price := range prices {
			if price-currentMin > maxProf {
				maxProf = price - currentMin
			}

			if price < currentMin {
				currentMin = price
			}
		}
	}

	return maxProf
}

func TestExample(t *testing.T) {
	t.Log(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	t.Log(maxProfit([]int{7, 6, 4, 3, 1}))
	t.Log(maxProfit([]int{7, 6, 4, 1, 2}))
}
