// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	// Case 1 - Output = 5
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfitOptimal(prices))

	// Case 2 - Output = 0
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfitOptimal(prices))
}

// MY SOLUTION: THIS EXCEEDS THE TIME LIMIT FOR HUGE INPUTS AS IT RUNS IN O(n^2) TIME.
// I maintain 2 pointers, j as the first pointer, and i as the second. Initially set the profit to 0.
// I take an element, and subtract it from all elements after from that, if my subtracton, is greater than current
// max profit, I have found new profit. Else continue. Once all subtractions have been completed, I move to the next element.
// Now this new element will be subtracted from all subsequent elements, and new profit will be calculated.
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	currentMaxProfit := 0
	j := 0

	for j < len(prices)-1 {
		for i := j + 1; i < len(prices); i++ {
			if (prices[i] - prices[j]) > currentMaxProfit {
				currentMaxProfit = prices[i] - prices[j]
			}
		}
		j++
	}

	return currentMaxProfit
}

// OPTIMAL SOLUTION: Here, we maintain the buying price, which should ideally be the lowest in the array we have processed until now,
// so if we come across a price lower than buying price, set the buying price to the lower price. If not, we check if
// current price - buying price is greater than max profit. If yes, set it to max profit and move to next price in array and
// repeat.
func maxProfitOptimal(prices []int) int {
	// Check if the prices slice is empty
	if len(prices) <= 1 {
		return 0 // If it's empty, return 0 as there can be no profit
	}

	// Initialize variables for tracking maximum profit and buying price
	maxProfit := 0           // Initialize maximum profit to 0
	buyingPrice := prices[0] // Initialize buying price to the first price in the slice

	// Iterate through the prices array
	for _, price := range prices {
		// Check if the current price is less than the buying price
		if price < buyingPrice {
			buyingPrice = price // If so, update the buying price to the current price
		} else { // Otherwise, if the current price is greater than or equal to the buying price
			// Calculate the profit if selling at the current price
			profit := price - buyingPrice
			// Update the maximum profit if the profit from selling at the current price is greater
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	// Return the maximum profit obtained
	return maxProfit
}
