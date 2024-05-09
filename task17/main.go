// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
)

func main() {

	sl := []int{1, 4, 7, 10, 11, 12, 15, 19, 26}

	res := binarySearch(sl, 12, 0, len(sl)-1)

	if res > 0 {
		fmt.Println("Индекс искомого элемента", res)
	} else {
		fmt.Println("Элемент отсутсвует в слайсе")
	}

}

func binarySearch(sl []int, x int, low int, up int) int {
	// проверяем если левый индекс больше правого то элемент не найден и возвращаем - 1
	if low > up {
		return -1
	}
	// находим индекс среднего элемента
	middle := (low + up) / 2

	// если индекс среднего элемента равен искомому возвращаем индекс
	if x == sl[middle] {
		return middle
	}
	// если искомое меньше индекса среднего элемента
	// отправяляем рекурсивно с начала до середины слайса
	if x < sl[middle] {
		return binarySearch(sl, x, low, middle-1)
		// если х больше то наоборот
	} else {
		return binarySearch(sl, x, middle+1, up)
	}
}
