package main

import (
	"fmt"
)

// medium
// ахАХхахвха я забыла r-l в скобочки убрать

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxAr := 0

	for l < r {
		heightL, heightR := height[l], height[r]
		maxAr = max(min(heightL, heightR)*(r-l), maxAr)
		//fmt.Printf("maxAr: %d, l: %d, r: %d, min: %d \n", maxAr, l, r, min(heightL, heightR))

		// Двигаем границу которая наименьшая из двух на следующую побольше
		// тут двигаем левую пока она не вырастет
		if heightL < heightR {
			for l < r && height[l] <= heightL {
				l++
			}
		} else {
			for l < r && height[r] <= heightR {
				r--
			}
		}
	}
	return maxAr
}

// Пример использования
func main() {
	hist := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(hist))
}
