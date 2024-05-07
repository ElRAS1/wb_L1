/* Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool,
channel из переменной типа interface{}. */

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Объявляем переменные разных типов
	num := 9             // int
	pi := 3.14           // float64
	str := "lalalal"     // string
	bl := true           // bool
	ch := make(chan int) // chan int

	// Вызываем функцию typeofRef() для каждой переменной, чтобы определить и вывести их типы
	typeofRef(num)
	typeofRef(pi)
	typeofRef(str)
	typeofRef(bl)
	typeofRef(ch)

}

func typeofRef(r interface{}) {
	// Используем reflect.TypeOf(r).Kind() для определения базового типа переменной r
	switch reflect.TypeOf(r).Kind() {

	// Проверяем тип переменной и выводим соответствующее сообщение
	case reflect.Int:
		fmt.Printf("type == %v\n", "int")
	case reflect.Chan:
		fmt.Printf("type == %v\n", "chan")
	case reflect.Bool:
		fmt.Printf("type == %v\n", "bool")
	case reflect.String:
		fmt.Printf("type == %v\n", "string")

	}

}
