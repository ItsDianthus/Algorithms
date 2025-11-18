package main

// medium
// монотонный стек

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))

	for ind := len(temperatures) - 1; ind >= 0; ind-- {
		num := temperatures[ind]

		for len(stack) > 0 {
			if num < temperatures[stack[len(stack)-1]] {
				res[ind] = stack[len(stack)-1] - ind
				stack = append(stack, ind)
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if len(stack) == 0 {
			stack = append(stack, ind)
			res[ind] = 0
		}
	}
	return res
}

func main() {
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})

}
