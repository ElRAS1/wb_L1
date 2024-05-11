// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func main() {
	sec := 0

	fmt.Printf("Введите сколько секунд будет ждать программа: ")
	fmt.Scan(&sec)
	for i := 0; i < 20; i++ {
		if i == 10 {
			fmt.Print("Ждем... ")
			MySleep(time.Duration(sec) * time.Second)
		}
		fmt.Print(i, " ")
	}

}

func MySleep(t time.Duration) {
	r := time.After(t)
	<-r
}
