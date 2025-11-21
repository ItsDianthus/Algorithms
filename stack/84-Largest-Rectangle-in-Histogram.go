package main

import (
	"fmt"
)

// hard
// сортировка по индексам и всё

func largestRectangleArea(heights []int) int {
	stack := []int{} // Стек хранит индексы баров
	maxres := 0      // Максимальная площадь

	for i := 0; i < len(heights); i++ {
		// Пока текущая высота < высоты на вершине стека
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]   // Индекс верхнего элемента
			stack = stack[:len(stack)-1] // Убираем его из стека

			right := i // Правая граница = текущий индекс
			left := -1 // Левая граница
			if len(stack) > 0 {
				left = stack[len(stack)-1] // Левая граница = предыдущий элемент стека
			}

			// Вычисляем площадь прямоугольника высоты top
			width := right - left - 1
			area := heights[top] * width
			maxres = max(maxres, area)
		}

		// Добавляем текущий индекс в стек
		stack = append(stack, i)
	}

	// Обрабатываем оставшиеся элементы в стеке
	right := len(heights)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		left := -1
		if len(stack) > 0 {
			left = stack[len(stack)-1]
		}

		width := right - left - 1
		area := heights[top] * width
		maxres = max(maxres, area)
	}

	return maxres
}

// Пример использования
func main() {
	hist := []int{1, 3, 4, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(largestRectangleArea(hist)) // Выводит максимальную площадь
}
