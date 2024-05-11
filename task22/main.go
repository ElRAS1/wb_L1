// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем две переменные, значения которых больше 2^20
	a := big.NewInt(4194304) // 2 в 21 степени
	b := big.NewInt(8388608) // 2 в 22 степени

	// Выполняем операции с a и b
	sum := addition(a, b)
	sub := subtraction(b, a)
	mul := multiplication(a, b)
	div := division(b, a)

	fmt.Printf("%v + %v = %v\n", a, b, sum)
	fmt.Printf("%v - %v = %v\n", b, a, sub)
	fmt.Printf("%v * %v = %v\n", a, b, mul)
	fmt.Printf("%v / %v = %v\n", b, a, div)
}

func addition(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Add(a, b)
}

func subtraction(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Sub(a, b)
}

func division(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Div(a, b)
}

func multiplication(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(a, b)
}
