package main

import (
	"fmt"
	"sort"
)

// easy

// Ну падумаешь заоверинжинирила, зато работает)))
func twoSum1(nums []int, target int) []int {
	myMap := make(map[int][]int)

	for idx, val := range nums {
		if _, exist := myMap[val]; exist {
			myMap[val] = append(myMap[val], idx)
		} else {
			myMap[val] = []int{idx}
		}
	}
	sort.Ints(nums)

	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] > target-nums[r] {
			r--
		} else if nums[l] < target-nums[r] {
			l++
		} else {
			if nums[l] == nums[r] {
				return myMap[nums[r]]
			}
			return []int{myMap[nums[r]][0], myMap[nums[l]][0]}
		}
	}
	return []int{-1}
}

// Окей адекватное решение
func twoSum(nums []int, target int) []int {
	mySet := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if idx, exist := mySet[nums[i]]; exist {
			return []int{i, idx}
		} else {
			mySet[target-nums[i]] = i
		}
	}
	return []int{-1}
}

func main() {
	fmt.Println(twoSum([]int{-3, 4, 3, 90}, 0))
}
