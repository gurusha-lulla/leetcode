// 88. Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	//Case 1 - Output = [1,2,2,3,5,6]
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	nums1 = []int{1, 2, 3, 0, 0, 0}
	m = 3
	nums2 = []int{2, 5, 6}
	n = 3
	mergeOptimal(nums1, m, nums2, n)

	//Case 2 - Output = [1]
	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge(nums1, m, nums2, n)
	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	mergeOptimal(nums1, m, nums2, n)

	//Case 2 - Output = [1]
	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	merge(nums1, m, nums2, n)
	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	mergeOptimal(nums1, m, nums2, n)
}

// My solution: Create a new array, while there are still elements in both nums1 AND nums2,
// check which element is smaller and append it in new array,
// as soon as any one of the input arrays is exhausted,
// append all elements from the other array into the new one
func merge(nums1 []int, m int, nums2 []int, n int) {
	newArr := make([]int, 0)
	nums1Index := 0
	nums2Index := 0

	for nums1Index < m && nums2Index < n {
		if nums1[nums1Index] < nums2[nums2Index] {
			newArr = append(newArr, nums1[nums1Index])
			nums1Index++
		} else {
			newArr = append(newArr, nums2[nums2Index])
			nums2Index++
		}
	}
	if nums1Index < m {
		newArr = append(newArr, nums1[nums1Index:m]...)
	}
	if nums2Index < n {
		newArr = append(newArr, nums2[nums2Index:n]...)
	}
	copy(nums1, newArr)

	fmt.Println(nums1)
}

// Optimal Solution: Here time complexity remains the same, space complexity becomes O(1)
// We start from the ends of the array, checking the biggest element and adding it to the first array, nums1 array
func mergeOptimal(nums1 []int, m int, nums2 []int, n int) {

	// Starting index will be the last position in the answer array, which is m+n-1
	// This raises the question - what if only nums1 had elements? Our base check is only n != 0
	// Answer is, in this case, our output will be picked from nums1 only, and nums2 is anyways empty, arrays are already merged
	// In this case, we don't need to do anything
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	fmt.Println(nums1)
}
