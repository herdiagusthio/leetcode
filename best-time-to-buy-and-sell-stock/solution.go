package main

func maxProfit(prices []int) int {
	// Edge case: if there are less than 2 prices, no profit can be made.
	if len(prices) < 2 {
		return 0
	}

	var maxProfit, price int
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		price = prices[i]
		if price < minPrice {
			minPrice = price
			continue
		}

		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	examples := [][]int{
		{7, 1, 5, 3, 6, 4}, // expected 5 (buy at 1, sell at 6)
		{7, 6, 4, 3, 1},    // expected 0 (no profit)
		{2, 4, 1},          // expected 2 (buy at 2, sell at 4)
		{},                 // expected 0 (empty)
	}

	for _, ex := range examples {
		println(maxProfit(ex))
	}
}
