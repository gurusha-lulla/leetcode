// 274. H-Index
// https://leetcode.com/problems/h-index/description/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{4, 4, 0, 0}))

	fmt.Println(hIndexCorrect([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndexCorrect([]int{4, 4, 0, 0}))
}

// MY SOLUTION: Sort the array in descending order, loop through the array. If you find the number of papers is equal to the number
// of citations, return that number.
// This solution is not able to handle the many corner cases, it DOES NOT WORK.
func hIndex(citations []int) int {

	if len(citations) < 1 {
		return 0
	}

	sort.Slice(citations, func(i int, j int) bool {
		return citations[i] > citations[j]
	})

	currentHIndex := citations[0]
	citationsWithCurrentHIndex := 0

	for _, currentCitation := range citations {
		if citationsWithCurrentHIndex >= currentHIndex {
			return currentHIndex
		}
		if currentCitation >= currentHIndex {
			citationsWithCurrentHIndex++
		} else {
			currentHIndex = currentCitation
			citationsWithCurrentHIndex++
		}
	}

	return citationsWithCurrentHIndex

}

// ACTUAL SOLUTION: This problem needs to be solved by counting sort.
// REFER: https://www.youtube.com/watch?v=mgG5KFTvfPw
func hIndexCorrect(citations []int) int {
	n := len(citations)
	paper_counts := make([]int, n+1)

	// Make an array paper_counts, to store which paper has how many counts. Like if 6 citations paper is encountered, increment count
	// for index 6, and so on.
	for _, v := range citations {
		indexToBeIncremented := int(math.Min(float64(n), float64(v)))
		paper_counts[indexToBeIncremented] = paper_counts[indexToBeIncremented] + 1
	}

	// Start with the last paper, and see if number of papers with that many citations, is less than the h index.
	// As soon as number of papers becomes greater than the h index, that means you found your max h index.
	h := n
	papers := paper_counts[n]

	for papers < h {
		h = h - 1
		papers = papers + paper_counts[h]
	}

	return h
}
