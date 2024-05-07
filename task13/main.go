// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	x := 1
	y := 2

	fmt.Printf("x = %d , y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("x = %d , y = %d", x, y)
}
