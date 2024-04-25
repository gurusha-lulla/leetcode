// 27. Remove Element
// https://leetcode.com/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Note: In the output array, we ignore all elements after output length
	// Case 1 - Output = 2, [2,2]
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	nums = []int{3, 2, 2, 3}
	val = 3
	fmt.Println(removeElementAlternative(nums, val))

	// Case 1 - Output = 5, [0,0,1,3,4]
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println(removeElement(nums, val))
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println(removeElementAlternative(nums, val))
}

// My solution: Get the length of the array in a variable called len.
// Loop through the array once, and set the value of any element equal to val, to 100,
// because, as per input conditions, no element of the array can be more than 50.
// Also decrement the len variable by 1. Now sort the array in ascending order using
// sort.Slice(array, func(i,j int) bool { return a[i] < a[j] })
// Good solution - beats 100% runtime
func removeElement(nums []int, val int) int {
	len := len(nums)
	for i, v := range nums {
		if v == val {
			nums[i] = 100
			len--
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums[0:len])
	return len
}

// Alternative Solution: This is not the best solution, however just an alternative approach.
// Loop through all the elements of the array using an index i & through the valid elements in array using index j.
// Each time, check if the element at position i, is not an invalid element that needs to be replaced.
// If the element is to be replaced, skip it, it will get replaced later by a valid element when pointer j comes to this index.
// If the element is VALID, swap position i to position j and increment j by 1.
// Thus, all valid elements will be in the first j positions and j will give length of valid array.
// Sorting is NOT needed here.
func removeElementAlternative(nums []int, val int) int {
	j := 0 // Valid elements

	for i, v := range nums {
		if v != val { // If valid element, shift it to the valid part of the array, tracked by index j
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	fmt.Println(nums[0:j])
	return j
}
