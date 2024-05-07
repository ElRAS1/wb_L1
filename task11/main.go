// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func main() {
	// Создаем две карты для представления множеств
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	// Создаем два тестовых множества
	initSet(set1, set2)

	// Вычисляем пересечение множеств с помощью функции interSet
	r := interSet(set1, set2)

	fmt.Println(r)
}

// interSet принимает две карты и возвращает карту, представляющую пересечение этих множеств.
func interSet(set1 map[int]struct{}, set2 map[int]struct{}) map[int]struct{} {
	// Создаем карту для хранения результата
	resSet := make(map[int]struct{})

	// Проходим по каждому ключу в первой карте
	for i := range set1 {
		// Проверяем, содержит ли вторая карта данный ключ
		if _, ok := set2[i]; !ok {
			continue
		}
		// Если ключ присутствует в обеих картах, добавляем его в результирующую карту
		resSet[i] = struct{}{}
	}

	return resSet
}

// initSet создает два тестовых множест
func initSet(set1 map[int]struct{}, set2 map[int]struct{}) {

	set1[1] = struct{}{}
	set1[2] = struct{}{}
	set1[3] = struct{}{}
	set1[4] = struct{}{}
	set1[5] = struct{}{}

	set2[5] = struct{}{}
	set2[6] = struct{}{}
	set2[1] = struct{}{}
	set2[7] = struct{}{}
	set2[2] = struct{}{}
}
