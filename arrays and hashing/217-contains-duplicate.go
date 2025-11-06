package main

import (
	"fmt"
	"sort"
)

// easy

func containsDuplicate(nums []int) bool {
	mySet := make(map[int]bool)
	
	for i := 0; i < len(nums); i++ {
		if mySet[nums[i]] {
			return true
		}
		mySet[nums[i]] = true
	}
	return false
}

// Сортировка + сравнение рядом стоящих элементов
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 5, 6}))
}
