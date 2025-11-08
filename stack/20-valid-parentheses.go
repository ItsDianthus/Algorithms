package main

import "fmt"

// easy

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ { // бежим по строке
		if s[i] == '{' || s[i] == '[' || s[i] == '(' { // если открывающая, кидаем в стек
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if s[i] != top && s[i]-top < 3 {
				// pop
				stack = stack[:len(stack)-1]
			} else { // ЗАБЫЛАААА ыыыыыыы
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(})"))
}
