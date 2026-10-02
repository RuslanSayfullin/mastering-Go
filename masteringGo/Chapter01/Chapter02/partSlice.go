package main

import (
	"fmt"
)

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)

	l := len(aSlice)

	// первые 5 элементов
	fmt.Println(aSlice[0:5])
	// первые 5 элементов
	fmt.Println(aSlice[:5])

	// последние два элемента
	fmt.Println(aSlice[l-2 : l])
	// последние два элемента
	fmt.Println(aSlice[l-2:])

	// первые пять элементов
	t := aSlice[0:5:10]
	fmt.Println(len(t), cap(t))
	// Элементы с индексами 2, 3, 4 емкость составит 10-2
	t = aSlice[2:5:10]
	fmt.Println(len(t), cap(t))
	// элементы с индексами 2 3 4 емкость составит 10-2
	t = aSlice[2:5:10]
	fmt.Println(len(t), cap(t))

	// элементы с индексами 0, 1, 2, 3, 4 новая емкость составит 6-0
	t = aSlice[:5:5]
	fmt.Println(len(t), cap(t))

}
