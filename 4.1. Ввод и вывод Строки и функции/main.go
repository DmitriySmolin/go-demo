package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64 // 0.0
	var userKg float64 // 0.0
	fmt.Print("__Калькулятор индекса массы тела \n")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес в кг: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Print(IMT)
}
