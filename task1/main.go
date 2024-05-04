package main

import (
	"fmt"
)

// первый вариант решения
type Human struct {
	name string
	age  int
}

// в структуру Action встраиваем структуру Human с ее полями и методами
type Action struct {
	admin bool
	*Human
}

func (a *Action) getAdmin() bool {
	return a.admin
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
	admin bool
	HumanInterface
}

func (h *Human2) GetAge() int {
	return h.age
}

func (a *Action2) GetAdmin() bool {
	return a.admin
}

func (h *Human2) GetName() string {
	return h.name
}

func main() {
	// первое решение
	per1 := Action{admin: true, Human: &Human{age: 32, name: "Ivan"}}

	fmt.Println(per1.getName())
	fmt.Println(per1.getAge())
	fmt.Println(per1.getAdmin())
	// второе решение

	per2 := Action2{HumanInterface: &Human2{name: "Vlad", age: 15}, admin: false}

	fmt.Println(per2.GetName())
	fmt.Println(per2.GetAge())
	fmt.Println(per2.GetAdmin())
}
