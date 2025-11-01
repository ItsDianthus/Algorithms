package main

import (
	"fmt"
	"iter"
	"net/netip"
)

func summaryRanges(nums []int) []string {
	var res []string

	if len(nums) == 0 {
		return nil
	}
	st := nums[0]
	fn := st
	nums = nums[1:]
	for _, num := range nums {
		//fmt.Println("round ", num)
		if num == fn+1 {
			fn += 1
		} else if num > fn+1 {
			res = append(res, rangeToString(st, fn)) // Unused variable 'res'
			st = num
			fn = num
		}
	}
	res = append(res, rangeToString(st, fn))

	return res
}

func rangeToString(a int, b int) string {
	if b == a {
		return fmt.Sprintf("%d", a)
	}
	return fmt.Sprintf("%d->%d", a, b)
}

func main() {

	lst := []int{1, 2, 3, 5, 6, 8, 10, 12, 13}
	fmt.Println(summaryRanges(lst))

}
