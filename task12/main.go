// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {

	// создаем слайс строк
	mas := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем множество из слайса
	res := createSet(mas)

	fmt.Println(res)

}

func createSet(m []string) map[string]struct{} {

	// создаем множество
	res := make(map[string]struct{})
	for _, i := range m {
		// проверяем если элемент уже есть в множестве то не добавялем, так как в множестве нету дублей
		if _, ok := res[i]; ok {
			continue
		}
		res[i] = struct{}{}
	}

	return res
}
