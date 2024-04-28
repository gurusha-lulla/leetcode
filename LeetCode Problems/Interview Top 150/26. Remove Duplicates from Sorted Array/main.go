// 26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	// Case 1 - Output = 2, [1,2]
	nums := []int{1, 2, 2}
	fmt.Println(removeDuplicates(nums))

	// Case 2 - Output = 5, [0,0,1,1,1,2,2,3,3,4]
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

// My Solution: Maintain 2 pointers in the array, i to loop through all elements in the array and j to track the valid part of the array.
// Start with the first element, we don't want to replace the current element, hence in the loop, j++ first and then add element.
// But yes, as soon as any element is found which is different from the one at position j, increment j and add the current element.
// This approach works because the array is in sorted order.
// So for example, if at j the current element is 0, we encounter 1, then after 0 we will add 1.
// Again if we encounter 1, now the element at j is also 1, so we don't add anything.
// Once we encounter 2, we add that and increment j, and so on.
// In the end, our value of j is the index of the last valid element, but index is always 1 less than length,
// hence we return j after incrementing it by 1.
func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	j++
	fmt.Println(nums[0:j])
	return j
}
