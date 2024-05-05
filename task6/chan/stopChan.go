// реализации с остановкой горутины с помощью каналов

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// создали канал для записи значений
	ch := make(chan int)
	// создали канал для завершения программы
	done := make(chan struct{})

	// запустили горутину и засыпаем на 2 секунды, а потом записываем в канал done.
	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	// в отдельно горутине запускаем бесконечный цикл где записываем в канал ch случайные числа и ждем когда из done можем прочитать и завершить программу
	go func() {
		for {
			select {
			case <-done:
				close(ch)
				return
			default:
				ch <- rand.Intn(25)
			}
		}
	}()

	for i := range ch {
		fmt.Println(i)
	}

}
