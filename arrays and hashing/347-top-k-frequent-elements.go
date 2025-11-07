package main

import (
	"fmt"
)

// medium

func topKFrequent1(nums []int, k int) []int {
	myMap := make(map[int]int)
	buckets := make([][]int, len(nums))
	var res []int

	for i := 0; i < len(nums); i++ {
		if _, exists := myMap[nums[i]]; exists {
			myMap[nums[i]] += 1
		} else {
			myMap[nums[i]] = 1
		}
	}

	for key, val := range myMap {
		buckets[val-1] = append(buckets[val-1], key)
	}

	buckNotZero := len(buckets) - 1
	for k > 0 {
		for buckets[buckNotZero] == nil { // Здесь находим ненулевой бакет
			buckNotZero--
		}
		res = append(res, buckets[buckNotZero]...) // надо заджоинить
		k -= len(buckets[buckNotZero])
		buckNotZero--
	}

	return res
}

// Замена на бакеты - мапу сделало beats 6,69 -> beats 92,35 по memory
// ничего лишнего, убрана проверка на мапу
func topKFrequent(nums []int, k int) []int {
	myMap := make(map[int]int)
	buckets := make(map[int][]int)

	var res []int

	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] += 1
	}

	for key, val := range myMap {
		buckets[val] = append(buckets[val], key)
	}

	buckNotZero := len(nums)
	for k > 0 {
		for buckets[buckNotZero] == nil { // Здесь находим ненулевой бакет
			buckNotZero--
		}
		res = append(res, buckets[buckNotZero]...) // надо заджоинить
		k -= len(buckets[buckNotZero])
		buckNotZero--
	}

	return res
}

func main() {
	myMap := make(map[int]int)
	myMap[0] += 1
	fmt.Println(myMap)

	//fmt.Println(topKFrequent([]int{1}, 1))
}
