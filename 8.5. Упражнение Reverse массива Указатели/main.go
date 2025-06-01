package main

import (
	"fmt"
)

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

// №1
func reverse(arr *[4]int) {
	i := 0
	j := len(arr) - 1

	for i < j {

		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

		i = i + 1
		j = j - 1
	}
}

// №2
// func reverse(arr *[4]int) {
// 	for index, value := range *arr {
// 		(*arr)[len(arr)-1-index] = value
// 	}
// }
