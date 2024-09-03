// 380. Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "math/rand"

func main() {

}

// IMPLEMENTATION IN O(1)
// MY SOLUTION: USE MAP

type RandomizedSet struct {
	dataMap map[int]int
}

func Constructor() RandomizedSet {
	dataMap := make(map[int]int, 0)
	return RandomizedSet{
		dataMap: dataMap,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dataMap[val]; !ok {
		this.dataMap[val] = 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		delete(this.dataMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	// randomValue := time.Now().Nanosecond() % len(this.dataMap)
	randomValue := rand.Intn(len(this.dataMap))
	currentValue := 0
	for k := range this.dataMap {
		if currentValue == randomValue {
			return k
		}
		currentValue++
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
