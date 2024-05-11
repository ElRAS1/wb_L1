/* Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false */

package main

import "fmt"

func main() {

	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("%v = %v\n", str1, fo(str1))
	fmt.Printf("%v = %v\n", str2, fo(str2))
	fmt.Printf("%v = %v\n", str3, fo(str3))
}

func fo(s string) bool {
	// создаем карту для хранения каждого элемента и значением true если он уже есть или false
	mp := make(map[string]bool)

	for _, i := range s {
		// если ключ повторяеться то выходим с false потому что символы в строке не уникальные
		if _, ok := mp[string(i)]; ok {
			return false
		}
		// если встречаем первый раз добавялем ключ
		mp[string(i)] = true
	}
	return true
}
