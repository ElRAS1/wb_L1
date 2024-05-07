/* Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. \
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc. */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем слайс с последовательностью температурных колебаний
	sl := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	startTime := time.Now()
	first(&sl)
	fmt.Println("Время выполнения первой программы без использования горутин ", time.Since(startTime))

	startTime = time.Now()
	second(&sl)
	fmt.Println("Время выполнения второй программы с  использованием горутин ", time.Since(startTime))

}

func first(sl *[]float32) {
	// создаем словарь для объеденения значений в группы с шагом в 10 градусов
	res := make(map[int][]float32)

	for _, i := range *sl {
		// Вычисляем группу
		x := (int(i) / 10) * 10
		// Добавляем в группу значение
		res[int(x)] = append(res[int(x)], i)
	}

	fmt.Println(res)
}

func second(sl *[]float32) {
	res := make(map[int][]float32)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, i := range *sl {
			// Вычисляем группу
			mu.Lock()
			x := (int(i) / 10) * 10
			// Добавляем в группу значение
			res[int(x)] = append(res[int(x)], i)
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println(res)
}
