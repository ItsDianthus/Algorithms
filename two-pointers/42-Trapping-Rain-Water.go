package main

import (
	"fmt"
)

// hard
// fuck

func trap1(height []int) int {
	st, collect := 0, 0
	res := 0
	for true {
		// st ставим на высокую точку
		for st < len(height)-1 && height[st] <= height[st+1] {
			st++
		}
		if st == len(height)-1 {
			continue
		}

		// collect пробегается пока не найдет такую же
		collect = st + 1
		for collect < len(height)-1 && height[collect] < height[st] {
			res += height[st] - height[collect]
			collect++
		}
		if collect == len(height)-1 {
			continue
		}
		// теперь коллект вырос до ст или выше
		st = collect
	}
	return res
}

func trap2(height []int) int {
	st, st2 := 0, 0
	stack := []int{}
	res := 0
	for true {
		// st ставим на высокую точку
		for st < len(height)-1 && height[st] <= height[st+1] {
			st++
		}
		if st == len(height)-1 {
			break
		}
		//fmt.Println("st found: ", st)

		// st2 пробегается пока не найдет такую же
		st2 = st + 1
		// ищем следующий st2
		for st2 < len(height)-1 {
			// st2 ищет максимум
			for st2 < len(height)-1 && height[st2] > height[st2+1] {
				st2++
			}
			for st2 < len(height)-1 && height[st2] <= height[st2+1] {
				st2++
			}
			if st2 == len(height)-1 {
				stack = stackAppend(stack, height, st2)
				// вышли за границы, но ст2 все еще локальный максимум
				break
			}
			//fmt.Println("st2 local found: ", st2)

			// тут st2+1 < st2, мы в локальном максимуме
			if height[st2] < height[st] { // Если это лок максимум
				stack = stackAppend(stack, height, st2)
			} else { // НАШЛИИИ СТ2
				stack = stackAppend(stack, height, st2)
				//fmt.Println("st2 final found: ", st2)
				break
			}
		}
		// обрабатываем стек
		for _, i := range stack {
			minH := min(height[st], height[i])
			for st < i {
				st++
				if height[st] < minH {
					res += minH - height[st]
				}
			}
		}

		st = st2
		// обрабатываем стек от st до st2
		// st = st2
	}
	return res
}

// Функция убирает все значения в стеке, которые были ниже новой вершины. Стек убывающий
func stackAppend(stack []int, heights []int, newInd int) []int {
	if len(stack) == 0 {
		return append(stack, newInd)
	}
	last := stack[len(stack)-1]
	for len(stack) > 0 && heights[last] < heights[newInd] {
		stack = stack[:len(stack)-1]
	}
	return append(stack, newInd)
}

// Есть три решения: стэк, префиксные/постфиксные суммы, два указателя
// делаем два указателя йоу
func trap(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	maxL, maxR := height[l], height[r]

	for l < r {
		if maxL < maxR {
			l++
			if height[l] > maxL {
				maxL = height[l]
			} else {
				res += maxL - height[l]
			}
		} else {
			r--
			if height[r] > maxR {
				maxR = height[r]
			} else {
				res += maxR - height[r]
			}
		}
	}
	return res
}

// Пример использования
func main() {
	hist := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(hist))
}
