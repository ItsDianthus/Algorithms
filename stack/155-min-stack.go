package main

import "fmt"

// medium

// Работает, но можно оптимизировать по памяти
//type MinStack struct {
//	stack       []int
//	minIdxStack []int
//}
//
//func Constructor() MinStack {
//	return MinStack{[]int{}, []int{}}
//}
//
//func (this *MinStack) Push(val int) {
//	this.stack = append(this.stack, val)
//	// {10, 8}
//	// {0, 1}
//	if len(this.stack) == 1 {
//		this.minIdxStack = append(this.minIdxStack, 0)
//	} else {
//		lastMinIdx := this.minIdxStack[len(this.minIdxStack)-1]
//
//		if this.stack[lastMinIdx] < val {
//			this.minIdxStack = append(this.minIdxStack, lastMinIdx)
//		} else {
//			this.minIdxStack = append(this.minIdxStack, len(this.stack)-1)
//		}
//	}
//}
//
//func (this *MinStack) Pop() {
//	// вытащить элемент и удалить его
//	this.stack = this.stack[:len(this.stack)-1]
//	this.minIdxStack = this.minIdxStack[:len(this.minIdxStack)-1]
//}
//
//func (this *MinStack) Top() int {
//	return this.stack[len(this.stack)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	return this.stack[this.minIdxStack[len(this.stack)-1]]
//}

// Нечитаемая штука с индексами, зато мемори экономит и работает быстро
type MinStack struct {
	stack       []int
	minIdxStack []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// {-5, -4}
	// {0, 1}
	if len(this.stack) == 1 {
		this.minIdxStack = append(this.minIdxStack, 0)
	} else {
		lastMinIdx := this.minIdxStack[len(this.minIdxStack)-1]

		if this.stack[lastMinIdx] < val {
			// this.minIdxStack = append(this.minIdxStack, lastMinIdx)
		} else {
			this.minIdxStack = append(this.minIdxStack, len(this.stack)-1)
		}
	}
}

func (this *MinStack) Pop() {
	// вытащить элемент и удалить его
	if len(this.stack) == 0 || len(this.minIdxStack) == 0 {
		fmt.Println("ssdjjsdjd")
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.stack) == this.minIdxStack[len(this.minIdxStack)-1] {
		this.minIdxStack = this.minIdxStack[:len(this.minIdxStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.minIdxStack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minIdxStack) == 0 {
		return -1
	}
	return this.stack[this.minIdxStack[len(this.minIdxStack)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(10)
	fmt.Println(obj.GetMin())
	obj.Push(8)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
}
