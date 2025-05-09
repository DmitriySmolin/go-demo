package main

import "fmt"

func main() {
	transactions := [6]int{1, 2, 3, 4, 5, 6}
	// transactionsNew := transactions
	// fmt.Println(transactions == transactionsNew)
	// transactionsNew[0] = 30
	// fmt.Println(transactionsNew)

	transactionsPartial := transactions[1:]
	transactionsNewPartial := transactionsPartial[:1]
	transactionsNewPartial[0] = 30
	transactionsNewPartial = transactionsPartial[0:4]
	fmt.Println(transactions)
	fmt.Println(transactionsPartial)
	fmt.Println(transactionsNewPartial)

	fmt.Println(len(transactionsPartial), cap(transactionsPartial))
	fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial))
}
