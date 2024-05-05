// реализация отстановки горутины с помощью WaitGroup

package main

import (
	"fmt"
	"sync"
)

func main() {
	var mas = [...]int{2, 4, 6, 8, 10}
	// создаем группу для отслеживания горутин
	var wg sync.WaitGroup
	res := make(chan int, len(mas))
	// добавляем кол-во горутин в группу
	wg.Add(len(mas))
	worker(&mas, &wg, res)

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

func worker(mas *[5]int, wg *sync.WaitGroup, res chan int) {
	for _, i := range mas {
		go func(i int) {
			// после выполнения горутины сбрасываем счетчик на один
			defer wg.Done()
			res <- (i * i)
		}(i)
	}
}
