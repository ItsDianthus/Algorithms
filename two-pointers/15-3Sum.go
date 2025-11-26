package main

import (
	"fmt"
	"sort"
)

// medium
// фикс левый указатель а дальше ту сам

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		fmt.Println(nums[i])
		target := -nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[l]+nums[r] == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				// двигаем правую и левую границу пока значения не изменятся
				// или хотя бы одну из них ы
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

// Пример использования
func main() {
	hist := []int{-4, -1, -1, 0, 1, 2}
	fmt.Println(threeSum(hist))
}
