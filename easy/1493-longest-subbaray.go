package main

import "fmt"

// Моя краказябля
func longestSubarray(nums []int) int {
	// Проверка на нил массив
	if nums == nil {
		return 0
	}

	var fin []int
	zerocnt := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 { // if its 1
			cnt++
			zerocnt = 0
		} else if zerocnt == 0 { // if it's first zero
			fin = append(fin, cnt)
			cnt = 0
			zerocnt++
		} else if zerocnt < 2 { // if it's second and zero
			fin = append(fin, 0)
			zerocnt++
		}
	}

	// last 1-ses
	if cnt != 0 {
		fin = append(fin, cnt)
	}

	// если получилось что только единицы/разбитые нулями
	if len(fin) == 1 {
		if fin[0] == len(nums) {
			return fin[0] - 1
		}
		return fin[0]
	}
	// если получилось что были только нули
	if len(fin) == 0 {
		return 0
	}

	// пробегаемся по всем
	maxs := 0
	for i := 0; i < len(fin)-1; i++ {
		maxs = max(maxs, fin[i]+fin[i+1])
	}
	//fmt.Println(fin)
	return maxs
}

// решение скользяшим окном
func longestSubarray2(nums []int) int {
	left, zeroCnt, maxL := 0, 0, 0

	if len(nums) == 0 {
		return 0
	}

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCnt++
		}

		for zeroCnt > 1 {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}

		maxL = max(maxL, right-left)
		if maxL == right-left {
			fmt.Println(left, right)
		}

	}
	return maxL
}

func main() {
	fmt.Println(longestSubarray2([]int{0, 0, 1, 0, 1, 0, 1, 1, 1, 0}))
	//longestSubarray([]int{1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0})
}
