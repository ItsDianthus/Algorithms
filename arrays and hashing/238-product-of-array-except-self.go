package main

import (
	"fmt"
)

// medium

func productExceptSelf1(nums []int) []int {
	res := make([]int, 0, len(nums))

	pref := make([]int, len(nums))
	suff := make([]int, len(nums))

	// tbd 2 elem

	for i, j := 0, len(nums)-1; i < len(nums) && j >= 0; i, j = i+1, j-1 {
		if i == 0 && j == len(nums)-1 {
			pref[i] = 1
			suff[j] = 1
		} else {
			pref[i] = pref[i-1] * nums[i-1]
			suff[j] = suff[j+1] * nums[j+1]
		}
	}
	for i := 0; i < len(nums); i++ {
		res = append(res, pref[i]*suff[i])
	}

	return res
}

// Можно суффикс считать через одну переменную, не додумалась. Гениально.
// Сначала в рез складываем все префиксы, потом домножаем на suff
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	suff := 1
	for j := len(nums) - 2; j >= 0; j-- {
		res[j] *= nums[j+1] * suff
		suff = suff * nums[j+1]
	}

	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2}))
}
