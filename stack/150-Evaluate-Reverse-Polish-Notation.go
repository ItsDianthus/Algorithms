package main

import (
	"fmt"
	"strconv"
)

// medium

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, item := range tokens {
		fmt.Println(stack)
		switch item {
		case "/":
			last := stack[len(stack)-1]
			prev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, prev/last)
		case "+":
			last := stack[len(stack)-1]
			prev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, prev+last)
		case "-":
			last := stack[len(stack)-1]
			prev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, prev-last)
		case "*":
			last := stack[len(stack)-1]
			prev := stack[len(stack)-2]

			stack = stack[:len(stack)-2]
			stack = append(stack, prev*last)

		default:
			num, _ := strconv.Atoi(item)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

// Оптимальный вариант: поделить прямо В СТЕКЕ, удалить крайний элемент
func evalRPN_optimal(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, item := range tokens {
		switch item {
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		default:
			num, _ := strconv.Atoi(item)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"1", "2", "+", "5", "-"}))
}
