/* Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	mas := [...]int{2, 4, 6, 8}
	fmt.Printf("Изначальный массив: %v\n", mas)
	fmt.Printf("Конечный массив: ")
	twoChan(&mas)
	fmt.Printf("\nВремя выполнения программы: %v", time.Since(startTime))
}

func twoChan(mas *[4]int) {
	chan1 := make(chan int) // Создаем канал для записи чисел из массива
	chan2 := make(chan int) // Создаем канал для записи чисел из chan1 и умноженные на 2 и последующий вывод
	var wg sync.WaitGroup

	wg.Add(2) // Добавляем 2 задачи в WaitGroup
	// Запускаем горутину для записи в первый канал из массива.
	worker1(&chan1, mas, &wg)
	// Запускаем горутину для записи в первого канала во второй числа умноженные на 2.
	worker2(&chan1, &chan2, &wg)

	// Ожидаем завершения всех горутин
	go func() {
		wg.Wait()
	}()

	// Выводим умноженные числа из канала в stdout
	for i := range chan2 {
		fmt.Print(i, " ")
	}
}

func worker1(chan1 *chan int, mas *[4]int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		defer close(*chan1) // Закрываем chan1 после записи всех данных
		for _, i := range mas {
			*chan1 <- i
		}
	}()
}

func worker2(chan1 *chan int, chan2 *chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		defer close(*chan2) // Закрываем chan2 после записи всех данных
		for i := range *chan1 {
			*chan2 <- i * 2
		}
	}()
}
