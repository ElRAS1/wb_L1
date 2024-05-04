package main

import "fmt"

// первый вариант решения
type Human struct {
	name string
	age  int
}

// в структуру Action встраиваем структуру Human с ее полями и методами
type Action struct {
	Human
}

func (h *Human) getAge() int {
	return h.age
}

func (h *Human) getName() string {
	return h.name
}

// второй вариант решения

// Определение интерфейса с необходимыми методами
type HumanInterface interface {
	GetName() string
	GetAge() int
}

// структура Human2 реализует интерфейс
type Human2 struct {
	name string
	age  int
}

// структура Action2 так же реализует интерфейс
type Action2 struct {
	HumanInterface
}

func (h Human2) GetAge() int {
	return h.age
}

func (h Human2) GetName() string {
	return h.name
}

func main() {
	// первое решение
	per1 := Action{Human{age: 32, name: "Ivan"}}

	fmt.Println(per1.getName())
	fmt.Println(per1.getAge())

	// второе решение

	per3 := Action2{HumanInterface: Human2{name: "Vlad", age: 15}}
	fmt.Println(per3.GetName())
	fmt.Println(per3.GetAge())
}
