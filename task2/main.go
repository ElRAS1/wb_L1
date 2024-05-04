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
	// создаем канал для сихронизации и блокировки горутин
	res := make(chan int)

	for _, i := range mas {
		// запускаем отдельную горутину для каждого элемента массива
		go func(i int) {
			// после записи канал блокируеться
			res <- i * i
		}(i)
	}
	for i := 0; i < len(mas); i++ {
		// после того как читаем из канала он мы снова можем в него писать
		fmt.Printf("%v ", <-res)
	}
}

func second(mas *[5]int) {
	wg := sync.WaitGroup{}

	wg.Add(len(mas)) // добавляем количество горутин в счетчки
	for _, i := range mas {
		// запускаем каждое вычесление квадрата в отдельной горутине
		go func(i int) {
			defer wg.Done() // после выполнение горутины сбрасываем счетчик на один
			fmt.Printf("%v ", i*i)
		}(i)
	}
	wg.Wait() // ждем пока все горутины завершат работу
}
