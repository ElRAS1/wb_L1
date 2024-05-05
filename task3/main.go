/* Дана последовательность чисел: 2,4,6,8,10. 
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mas = [...]int{2, 4, 6, 8, 10}

	fmt.Print("Первый вариант программы: ")
	startTime := time.Now()
	first(&mas)
	fmt.Println()
	fmt.Printf("Время выполнения первого варианта: %v\n", time.Since(startTime))

	fmt.Print("Второй вариант программы: ")
	startTime = time.Now()
	second(&mas)
	fmt.Println()
	fmt.Printf("Время выполнения второго варианта: %v\n", time.Since(startTime))
}

func first(mas *[5]int) {
	// создаем группу для отслеживания горутин
	var wg sync.WaitGroup
	res := make(chan int, len(mas))
	// добавляем кол-во горутин в группу
	wg.Add(len(mas))
	for _, i := range mas {
		go func(i int) {
			// после выполнения горутины сбрасываем счетчик на один
			defer wg.Done()
			res <- (i * i)
		}(i)
	}
	// ждем пока все горутины закончат выполнение и закрываем канал
	go func() {
		defer close(res)
		wg.Wait()
	}()
	x := 0
	for i := range res {
		x += i
	}
	fmt.Print(x)
}

func second(mas *[5]int) {
	res := make(chan int, len(mas))
	done := make(chan struct{}) // Канал для сигнализации о завершении работы всех горутин

	for _, i := range mas {
		go func(i int) {
			res <- (i * i)
			done <- struct{}{} // Отправляем сигнал о завершении работы горутины
		}(i)
	}

	// Закрываем канал с результатами после получения всех результатов
	go func() {
		for range mas {
			<-done // Ждем сигналы от всех горутин
		}
		close(res)
	}()

	x := 0
	for i := range res {
		x += i
	}
	fmt.Print(x)
}
