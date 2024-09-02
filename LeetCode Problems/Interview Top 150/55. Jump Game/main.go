// 55. Jump Game
// https://leetcode.com/problems/jump-game/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))

	fmt.Println(canJumpCorrect([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJumpCorrect([]int{3, 2, 1, 0, 4}))
}

// MY SOLUTION - This was supposed to be a greedy approach, taking as many jumps as possible in one shot.
// However, in an input like [2, 5, 0, 0] it is pretty evident that I can take a jump of 1, then of 2
// and make the requirement, however my greedy approach would take 2 directly, and reach 0.
// After that it would not be able to make any jumps. This is incorrect.
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	requiredJumps := len(nums) - 1
	currentPosition := 0
	valueAtCurrentPosition := nums[0]
	for currentPosition < requiredJumps {
		currentPosition = currentPosition + valueAtCurrentPosition
		if currentPosition < requiredJumps {
			valueAtCurrentPosition = nums[currentPosition]
			if valueAtCurrentPosition == 0 && currentPosition < requiredJumps {
				return false
			}
		}
	}
	return true
}

// SOLUTION: This seems to be a greedy solution only, where we remove 1 jump each time we make it. As soon as we see the
// current available jumps in array be greater than what we have, we update our available jumps to that. If still available
// jumps run out, we return false.
// Refer: https://leetcode.com/problems/jump-game/solutions/4534808/super-simple-intuitive-8-line-python-solution-beats-99-92-of-users/?envType=study-plan-v2&envId=top-interview-150
func canJumpCorrect(nums []int) bool {
	availableJumps := 0
	for _, jumps := range nums {
		if availableJumps < 0 {
			return false
		} else if jumps > availableJumps {
			availableJumps = jumps
		}
		availableJumps--
	}
	return true
}
