package main

import (
	"fmt"
	"sort"
)

// medium
// сортировка по индексам и всё

func carFleet(target int, position []int, speed []int) int {
	indx := make([]int, len(position))
	for i := 0; i < len(indx); i++ {
		indx[i] = i
	}

	sort.Slice(indx, func(i, j int) bool { return position[indx[i]] > position[indx[j]] })

	res := []float64{calcTime(target, position[indx[0]], speed[indx[0]])}
	for _, ind := range indx {
		hours := calcTime(target, position[ind], speed[ind])

		if res[len(res)-1] < hours {
			res = append(res, hours)
		}
	}
	return len(res)
}

func calcTime(target int, pos int, sp int) float64 {
	return float64(target-pos) / float64(sp)
}

func main() {
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
}
