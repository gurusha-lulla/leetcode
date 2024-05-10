// 169. Majority Element
// https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	// Case 1 - Output = 3
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))

	// Case 2 - Output = 2
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

// My Solution: Create a hash map to keep track of all element counts and whenever an element's count becomes more
// than at least 50% of the array length, it is definietly the majority element in the array.
// The -1 return is never reached, as the question explicitly mentions that there will ALWAYS be a majority element.
func majorityElement(nums []int) int {
	n := len(nums)
	mapOfNums := make(map[int]int)

	for _, v := range nums {
		mapOfNums[v] = mapOfNums[v] + 1
		if mapOfNums[v] > n/2 {
			return v
		}
	}

	return -1
}

// NOTE: Did not find any alternate solution, yes Moore has one with a linear pass, but it does not work in all cases. If you
// need to refer that, please refer to this link: https://www.cs.utexas.edu/~moore/best-ideas/mjrty/index.html

// The solution is as follows:

// func majorityElement(nums []int) int {
//     majority_element, majority_element_frequency := nums[0], 1
//     for i := 1; i < len(nums); i++ {
//         if majority_element_frequency == 0 {
//             majority_element, majority_element_frequency = nums[i], 1
//         } else {
//             if nums[i] == majority_element {
//                 majority_element_frequency += 1
//             } else {
//                 majority_element_frequency -= 1
//             }
//         }
//     }
//     return majority_element
// }

// It fails for the input {1, 2, 1, 3, 1, 4, 1, 5} => Output comes 5, which is incorrect
