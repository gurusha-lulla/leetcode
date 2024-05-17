// 189. Rotate Array
// https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	// Case 1 - Output = [5,6,7,1,2,3,4]
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 3
	rotateAlternate(nums, k)
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 3
	rotateCyclic(nums, k)

	// Case 2 - Output = [3,99,-1,-100]
	nums = []int{-1, -100, 3, 99}
	k = 2
	rotate(nums, k)
	nums = []int{-1, -100, 3, 99}
	k = 2
	rotateAlternate(nums, k)
	nums = []int{-1, -100, 3, 99}
	k = 2
	rotateCyclic(nums, k)
}

// My Solution: The base conditions are simple, if the array has 1 element or less, no rotation is possible, return.
// If the value of k, that is number of rotations to be done is 0, no rotations needed, exit.
// Now, while rotations are to be done, that is k is not 0, I take the last number and copy the array in such a way,
// That the current last number becomes the first number, and all the further numbers, from the first to the second last number,
// that is, except the number to be copied, are pushed after the currentLastNumber. We then decrement k, indicating that
// one rotation got completed like this.
// THIS SOLUTION TIMES OUT FOR LARGE INPUT.
func rotate(nums []int, k int) {
	// Base Conditions - If array has only 1 element or less, no rotation can happen.
	// If k is 0, which is a non negative number, there's no rotation to be made
	if len(nums) < 2 || k == 0 {
		return
	}

	for k != 0 {
		currentLastNumber := nums[len(nums)-1]
		copy(nums, append([]int{currentLastNumber}, nums[0:len(nums)-1]...))
		k--
	}

	fmt.Println(nums)
}

// My Solution 2: The base conditions are simple, if the array has 1 element or less, no rotation is possible, return.
// If the value of k, that is number of rotations to be done is 0, no rotations needed, exit.
// If k is more than array length, we waste time doing full rotations, hence just mod it by length of array.
// Just copy the last k elements in the first k positions & all elements except for last k elements, after that.
// This works for LARGE INPUT ALSO.
func rotateAlternate(nums []int, k int) {
	k = k % len(nums)
	// Base Conditions - If array has only 1 element or less, no rotation can happen.
	// If k is 0, which is a non negative number, there's no rotation to be made
	if len(nums) < 2 || k == 0 {
		return
	}
	copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k+1]...))
	fmt.Println(nums)
}

// MOST EFFICIENT SOLUTION: Start with the first element, replace the (start+k)th element with start element
// while keeping the (start+k)th element value in a temp variable. Now move to ((start+k) + k)th element
// and replace the value with last temp and update the temp with curent value. On every replacement,
// increment a global counter to keep track of number of replacements,
// this will help prevent repplacement of already replaced value.
func rotateCyclic(nums []int, k int) {
	k = k % len(nums)
	// Base Conditions - If array has only 1 element or less, no rotation can happen.
	// If k is 0, which is a non negative number, there's no rotation to be made
	if len(nums) < 2 || k == 0 {
		return
	}

	maxNumberOfMoves := len(nums)
	numOfMovesMade := 0
	start := 0

	for numOfMovesMade < maxNumberOfMoves {
		temp := nums[start]
		nextIndex := start

		for {
			nextIndex = (nextIndex + k) % len(nums)
			nums[nextIndex], temp = temp, nums[nextIndex]
			numOfMovesMade++
			if start == nextIndex {
				break
			}
		}

		start++
	}
	fmt.Println(nums)
}
