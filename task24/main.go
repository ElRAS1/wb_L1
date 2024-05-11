/* Разработать программу нахождения расстояния между двумя точками
, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

// NewPoint - конструктор для создания нового экземпляра Point.
func NewPoint(x int, y int) *Point {
	return &Point{x: x, y: y}
}

func main() {
	// Создание двух точек.
	first := NewPoint(2, 4)
	second := NewPoint(4, 6)

	// Вычисление и вывод расстояния между точками.
	fmt.Println(PointDistance(first, second))
}

// PointDistance - вычисляет расстояние между двумя точками.
func PointDistance(first *Point, second *Point) float64 {
	return math.Sqrt(math.Pow(float64(second.x-first.x), 2) + math.Pow(float64(second.y-first.y), 2))
}
