/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow». */

package main

import (
	"fmt"
	"strings"
)

func main() {
	st := "snow dog sun"

	fmt.Print(ReverseString(st))
}

func ReverseString(st string) string {
	// Используем strings.Builder для эффективного склеивания строк
	var builder strings.Builder
	// Разделяем строку на слова и возвращаем слайс
	tmp := strings.Split(st, " ")
	// Узнаем количество слов
	j := len(tmp) - 1

	sep := " " // Переменная для разделителя между словами
	for range tmp {
		if j == 0 { // Если это последнее слово, не добавляем разделитель
			sep = ""
		}
		// Добавляем слово с разделителем в начало результирующей строки
		builder.WriteString(tmp[j] + sep)
		j-- // Сдвигаемся к следующему слову
	}
	return builder.String() // Возвращаем итоговую строку
}
