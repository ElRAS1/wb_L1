/* Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode. */

package main

import "fmt"

func revers(s interface{}) (string, error) {
	x, ok := s.(string)
	if !ok {
		// Обработка случая, когда i не содержит строку
		return "", fmt.Errorf("input is not a string")
	}
	rn := []rune(x)
	res := []rune{}
	j := len(rn) - 1

	for range rn {
		res = append(res, rn[j])
		j--
	}
	return string(res), nil
}
