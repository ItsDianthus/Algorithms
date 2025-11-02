package main

import "fmt"

type RecentCounter struct {
	allCalls []int
}

func Constructor() RecentCounter {
	return RecentCounter{nil}
}

// НЕВКУСНО, ПОТОМУ ЧТО КВАДРАТ ДЕЙСТВИЙ
//func (this *RecentCounter) Ping(t int) int {
//	cnt := 0
//	this.allCalls = append(this.allCalls, t)
//	for _, num := range this.allCalls {
//		if num >= t-3000 && num <= t {
//			cnt++
//		}
//	}
//	this.allCalls = this.allCalls[len(this.allCalls)-cnt:]
//	fmt.Println(this.allCalls)
//	return cnt
//}

func (this *RecentCounter) Ping(t int) int {
	this.allCalls = append(this.allCalls, t)
	cnt := 0

	// НЕВКУСНО потому что рейндж очень долгииий
	//for _, v := range this.allCalls {
	//	if v < t-3000 {
	//		cnt++
	//	} else {
	//		break
	//	}
	//}

	// Нямка
	for len(this.allCalls) > 0 && this.allCalls[0] < t-3000 {
		this.allCalls = this.allCalls[1:]
	}

	this.allCalls = this.allCalls[cnt:]

	return len(this.allCalls)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
}
