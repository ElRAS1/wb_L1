// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

func ChangeBit(num, pos, bit int64) int64 {
	// Создаем маску путем сдвига 1 типа int64 на pos (на позицию, которая нам нужна)
	mask := int64(1 << pos)

	res := num

	// Применяем логическое И между числом и маской, чтобы проверить, какой бит установлен на интересующей нас позиции.
	// Если биты равны, то ничего не делаем.
	// Если биты не равны и бит, который нам нужно поменять, равен 1, то применяем логическое И и инвертируем бит.
	if mask&num == 1 && bit == 0 {
		res = num &^ mask
		// Если бит, который нам нужно поменять, равен 0, то применяем логическое ИЛИ, чтобы установить бит.
	} else if mask&num == 0 && bit == 1 {
		res = num | mask
	}

	return res
}