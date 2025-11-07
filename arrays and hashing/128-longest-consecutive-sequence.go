package main

import (
	"fmt"
)

// medium

func longestConsecutive1(nums []int) int {
	mySet := make(map[int]bool)
	maxCnt := 0

	for i := 0; i < len(nums); i++ {
		mySet[nums[i]] = true
	}

	for k, _ := range mySet {
		if !mySet[k-1] { // первый элемент
			cnt := 0
			i := k
			for mySet[i] {
				cnt++
				i++
				maxCnt = max(maxCnt, cnt)
			}
		}
	}

	return maxCnt
}

// Улучшение засчёт фора, обращений к мапе в два раза меньше
func longestConsecutive(nums []int) int {
	mySet := make(map[int]bool)
	maxCnt := 0

	for i := 0; i < len(nums); i++ {
		mySet[nums[i]] = true
	}

	for k, _ := range mySet {
		if !mySet[k-1] { // первый элемент
			cnt := 0
			i := k
			for {
				if _, ok := mySet[i]; !ok {
					break
				}
				cnt++
				i++
			}
			if cnt > maxCnt {
				maxCnt = cnt
			}

		}
	}

	return maxCnt
}

func main() {
	fmt.Println(longestConsecutive([]int{1, 3}))

}
