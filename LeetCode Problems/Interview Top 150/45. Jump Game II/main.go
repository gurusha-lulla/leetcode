// 45. Jump Game II
// https://leetcode.com/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{1, 3, 2}))

	fmt.Println(jumpCorrect([]int{2, 3, 1, 1, 4}))
	fmt.Println(jumpCorrect([]int{2, 3, 0, 1, 4}))
	fmt.Println(jumpCorrect([]int{1, 3, 2}))
}

// MY SOLUTION: Make a jump each time you reset your numberOfAvailableJumps. However, this does not work with an input like 1,3,2;
// because, the only jump (where essentially currentPosition is more than your number of jumps available), is at 3. However jumps made
// are actually two, from 1 -> 3 & 3 -> 2. Hence this fails.
func jump(nums []int) int {
	availableJumps := 0
	numberOfJumpsMade := 0
	for _, jumps := range nums[1:] {
		if jumps > availableJumps {
			availableJumps = jumps
			numberOfJumpsMade++
		}
		availableJumps--
	}
	return numberOfJumpsMade
}

// ACTUAL SOLUTION: If the range of my CURRENT JUMP availablility is [curBegin, curEnd], and the farthest I can make it in the current
// values, is stored in curFarthest.
// Refer: https://www.youtube.com/watch?v=vBdo7wtwlXs
// Refer: https://leetcode.com/problems/jump-game-ii/solutions/18014/concise-o-n-one-loop-java-solution-based-on-greedy/?source=submission-noac
func jumpCorrect(nums []int) int {
	jumps, curEnd, curFarthest := 0, 0, 0
	for i, v := range nums[:len(nums)-1] {
		curFarthest = int(math.Max(float64(curFarthest), float64(i+v)))
		if curEnd == i {
			jumps++
			curEnd = curFarthest
		}
	}
	return jumps
}
