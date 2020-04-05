package src_test

import "testing"

func maxProfit(prices []int) int {
	maxProf := 0
	if len(prices) > 1 {
		buyPrice := prices[0]

		for i := 1; i < len(prices); i++ {
			if prices[i] <= prices[i-1] {
				// We should sell the stock on yesterday
				maxProf = maxProf + prices[i-1] - buyPrice
				// Attempt to buy
				buyPrice = prices[i]
			} else if i == len(prices)-1 {
				maxProf = maxProf + prices[i] - buyPrice
			}
		}

	}

	return maxProf
}

func TestExample(t *testing.T) {
	t.Log(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	t.Log(maxProfit([]int{1, 2, 3, 4, 5}))
	t.Log(maxProfit([]int{7, 6, 4, 3, 1}))
}
