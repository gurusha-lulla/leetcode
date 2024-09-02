// 122. Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	// fmt.Print(maxProfit([]int{7, 1, 5, 3, 6, 4})) - Some index out of range issue in my solution
	fmt.Print(maxProfitGreedy([]int{7, 1, 5, 3, 6, 4}))
}

// MY SOLUTION - I will keep the buyingPrice in a variable, which will initially be the very first element of array.
// I will keep totalProfit in another variable completely. I will keep the currentProfit, which is the max profit
// I can get by selling after the CURRENT BUYING PRICE. As soon as I get the MAX profit I can get after the current
// buying price, I will shift to the next element as buyingPrice, right after the one I used to subtract.
// ISSUE - This doesn't work as it tries to get maximum profit off of 1 buy and sell. We need to use the greedy algorithm instead.
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	totalProfit := 0
	buyingPrice := prices[0]
	buyingPriceIndex := 0

	currProfit, nextIndex := getMaxProfitObtainableAtCurrentBuyingPrice(buyingPrice, prices, buyingPriceIndex)

	for nextIndex <= len(prices) {
		totalProfit += currProfit
		buyingPrice = prices[nextIndex]
		buyingPriceIndex = nextIndex
		currProfit, nextIndex = getMaxProfitObtainableAtCurrentBuyingPrice(buyingPrice, prices, buyingPriceIndex)
	}

	return totalProfit
}

func getMaxProfitObtainableAtCurrentBuyingPrice(buyingPrice int, prices []int, buyingPriceIndex int) (profit, indexOfSellingDay int) {
	profit = 0
	if buyingPriceIndex < len(prices)-1 {
		for i := buyingPriceIndex + 1; i < len(prices); i++ {
			currProfit := prices[i] - buyingPrice
			if currProfit > profit {
				profit = currProfit
				indexOfSellingDay = i
			}
		}
	}
	if profit == 0 && indexOfSellingDay == 0 {
		indexOfSellingDay = buyingPriceIndex + 1
	}
	return profit, indexOfSellingDay
}

// PROPER SOLUTION - Use the greedy algorithm. As soon as you get a price increase, sell.
// Refer to this for detailed explaination: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/solutions/4836121/simple-beginner-friendly-dry-run-greedy-approach-readable-sol-time-o-n-space-o-1-gits/?envType=study-plan-v2&envId=top-interview-150
func maxProfitGreedy(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buyingPrice := prices[0]
	max := 0
	for _, price := range prices[1:] {
		fmt.Println(price, "\t", buyingPrice)
		if price > buyingPrice {
			max = max + (price - buyingPrice)
		}
		buyingPrice = price
	}
	return max
}
