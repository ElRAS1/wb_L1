/* Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика. */

package main

import (
	"fmt"
	"sync"
)

type Incr struct {
	mu sync.Mutex
	wg sync.WaitGroup
	i  int
}

// Добавляем к счетчику 1
func (c *Incr) Counter() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.i++
	c.wg.Done()
}

func main() {

	test := Incr{}

	// запускаем 25 горутин для добавление к счетчику
	for i := 0; i < 25; i++ {
		test.wg.Add(1)
		go test.Counter()
	}

	// ждем завершение всех горутин
	test.wg.Wait()

	fmt.Println("Значение счетчика:", test.i)
}
