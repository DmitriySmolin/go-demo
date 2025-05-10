package main

import (
	"fmt"
)

func main() {
	transactions := []int{0, 20, 35}
	temp := transactions
	// slice := transactions[0:2]
	// slice[0] = 100
	// slice[4] = 100
	// fmt.Println(transactions)
	// fmt.Println(slice)

	// newTransactions := append(transactions, 100)
	// fmt.Println(transactions)
	// fmt.Println(newTransactions)

	transactions = append(transactions, 100)
	fmt.Println(temp)
	fmt.Println(transactions[1:])
}
