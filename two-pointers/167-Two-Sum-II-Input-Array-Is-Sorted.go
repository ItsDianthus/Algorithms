package main

// medium??
//

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		if target == numbers[i]+numbers[j] {
			return []int{i + 1, j + 1}
		} else if target < numbers[i]+numbers[j] {
			i--
		} else {
			j++
		}
	}
	return []int{-1}
}

// Пример использования
func main() {
	//hist := []int{1, 3, 4, 1, 1, 1, 1, 1, 1, 1, 1}
	// fmt.Println(twoSum(hist)) // Выводит максимальную площадь
}
