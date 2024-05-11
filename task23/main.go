// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	sl2 := []int{1, 2, 3, 4, 5, 6, 7}

	Delete(&sl, 2)
	Remove(&sl2, 2)
	fmt.Println(sl)
	fmt.Println(sl2)
}

// удаление элемента из слайса в ситуации когда порядок элементов не важен
func Delete(sl *[]int, ind int) {
	(*sl)[ind], (*sl)[len(*sl)-1] = (*sl)[len(*sl)-1], (*sl)[ind]
	*sl = (*sl)[:len(*sl)-1]
}

// удаление элемента из слайса в ситуации когда порядок элементов важен
func Remove(sl *[]int, ind int) {
	*sl = append((*sl)[:ind], (*sl)[ind+1:]...)
}
