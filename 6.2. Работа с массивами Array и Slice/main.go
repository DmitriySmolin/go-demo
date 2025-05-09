package main

import "fmt"

func main() {
	// transactions := [3]int{5, 10, -7}
	// banks := [2]string{"Тинькофф", "Альфа"}

	// fmt.Println(transactions)
	// fmt.Println(banks)

	transactions := [3]int{1, 2, 3}
	banks := [2]string{}

	fmt.Println(transactions[1])
	banks[0] = "Тинькофф"
	fmt.Println(banks)
}
