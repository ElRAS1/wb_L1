// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

type Dog struct {
}

func (d *Dog) Gav() {
	fmt.Println("Gav Gav brother")
}

type Cat struct {
}

func (c *Cat) Meow() {
	fmt.Println("Meow meow bro")
}

// Определение интерфейса Reaction с методом Reaction(), который должен быть реализован всеми адаптируемыми объектами.
type Reaction interface {
	Reaction()
}

// Реализация метода Reaction() для структуры Dog, который вызывает метод Gav().
func (d *Dog) Reaction() {
	d.Gav()
}

// Реализация метода Reaction() для структуры Cat, который вызывает метод Meow().
func (c *Cat) Reaction() {
	c.Meow()
}

// CatAdapter адаптирует объект Cat к интерфейсу Reaction, вызывая метод Meow() при вызове Reaction().
type CatAdapter struct {
	*Cat
}

// DogAdapter адаптирует объект Dog к интерфейсу Reaction, вызывая метод Gav() при вызове Reaction().
type DogAdapter struct {
	*Dog
}

// Функция-конструктор ConstCat принимает объект Cat и возвращает адаптер CatAdapter, соответствующий интерфейсу Reaction.
func ConstCat(cat *Cat) Reaction {
	return &CatAdapter{cat}
}

// Функция-конструктор ConstDog принимает объект Dog и возвращает адаптер DogAdapter, соответствующий интерфейсу Reaction.
func ConstDog(dog *Dog) Reaction {
	return &DogAdapter{dog}
}

func main() {
	// Создаем экземпляры Cat и Dog с помощью функций-конструкторов.
	c := ConstCat(&Cat{})
	d := ConstDog(&Dog{})

	// Вызываем метод Reaction() для адаптеров Cat и Dog, что демонстрирует адаптацию их поведения к общему интерфейсу.
	c.Reaction()
	d.Reaction()
}
