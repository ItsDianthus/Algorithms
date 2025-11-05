package main

import "fmt"

func removeElement(nums []int, val int) int {
	f, l := 0, len(nums)-1

	for f < l {
		if nums[f] == val {
			for l >= 0 && nums[l] == val {
				l--
				if l < 0 {
					return 0 // все валы
				}
			}
			if f > l {
				return l + 1
			}
			nums[f] = nums[l]
			l--
		}
		f++
	}
	if f == l && nums[f] == val {
		l--
	}

	if len(nums) == 1 && nums[0] == val {
		return 0 // краевой случай если массив состоит из [val]
	}
	return l + 1
}

func main() {
	nums := []int{2, 3, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}
