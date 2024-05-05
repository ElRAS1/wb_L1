// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{} // Используем WaitGroup для ожидания завершения всех горутин
	mu := sync.Mutex{}     // Используем Mutex для синхронизации доступа к map
	dict := make(map[int]int)
	y := 0
	for i := 0; i < 10; i++ {
		y++
		wg.Add(1) // Увеличиваем счетчик горутин
		mu.Lock() // Блокируем доступ к map
		go func(y int) {
			defer mu.Unlock() // Освобождаем Mutex после завершения горутины
			defer wg.Done()   // Уменьшаем счетчик горутин после завершения горутины
			dict[y] = y
			fmt.Print(y, " ")                  // Выводим y для демонстрации
			time.Sleep(100 * time.Millisecond) // Добавляем задержку для демонстрации
		}(y)
	}
	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println()
	fmt.Println(dict)
}
