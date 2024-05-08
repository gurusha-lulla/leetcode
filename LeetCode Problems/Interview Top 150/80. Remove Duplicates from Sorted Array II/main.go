// 80. Remove Duplicates from Sorted Array II
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	// Note: In the output array, we ignore all elements after output length
	// Case 1 - Output = 5, [1,1,2,2,3]
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	nums = []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicatesAlternative(nums))
	nums = []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicatesOptimal(nums))

	// Case 2 - Output = 7, [0,0,1,1,2,3,3]
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicatesAlternative(nums))
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicatesOptimal(nums))
}

// My Solution: Create a map which will act as a hash map, and store how many times an element has occurred until now.
// Create an element j, which will count how many valid elements have been kept in place in the array up until now.
// Start processing the array, each time you find an element which has occurred less than 2 times, let it stay, store it's count
// in hash map and increase the value of j, as it is a valid element. As soon as you see an element whose value in the hash map is
// exactly 2, that is, it has already occurred twice, we will eliminate all subsequent occurrences of it.
// Splice the array to remove it and reduce the index by 1, as the next element to be processed got spliced out, so we need to process
// the same index again.
func removeDuplicates(nums []int) int {
	mapOfNumbers := make(map[int]int, 0)
	j := 0

	for i := 0; i < len(nums); i++ {
		if mapOfNumbers[nums[i]] == 2 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			mapOfNumbers[nums[i]] = mapOfNumbers[nums[i]] + 1
			j++
		}
	}
	fmt.Println(nums[0:j])
	return j
}

// Alternative Solution: This one is way easier, just maintain a counter j to track the valid part of the array, and a map
// being used as a hash map to store how many times an element has occurred. Range through the array, if found an element
// that has occurred less that 2 times, add it to position j and increment j. Also increment the map counter of the element.
func removeDuplicatesAlternative(nums []int) int {
	mapOfNumbers := make(map[int]int, 0)
	j := 0

	for _, v := range nums {
		if mapOfNumbers[v] < 2 {
			mapOfNumbers[v] = mapOfNumbers[v] + 1
			nums[j] = v
			j++
		}
	}

	fmt.Println(nums[0:j])
	return j
}

// Optimal Solution: Here we don't use any map, instead this solution, while being optimal, can only be used when max allowed duplicates
// are exactly 2. Right away, if the number of elements in the array is less than 2, we don't really have anything to process,
// so return length of nums and nums itself, as they are.
// Next, as the max number of duplicates allowed is 2, we start from the third element.
// Look at the if condition, variable j is used to track valid part of array. For every element,
// The first if checks if current and previous element are not same, if they are not, it is obvious it is a new & thereby valid element.
// Add it at place j and increment j. The second if condition sees if the past 2 elements are not the same, that means the past 2 elements
// are NOT duplicates of each other. If they are not, either the element we add newly is a brand new one, or just 1 duplicate of the
// previous one, so we can add it. If either of these 2 if's is true, the element is valid, else, it is invalid.
func removeDuplicatesOptimal(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[j-1] != nums[i] || nums[j-1] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums[0:j])
	return j
}
