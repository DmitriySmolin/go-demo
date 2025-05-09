package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64 // 0.0
	var userKg float64     // 0.0
	fmt.Println("__Калькулятор индекса массы тела")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес в кг: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
	// fmt.Print("Ваш индекс массы тела: ", IMT)
}
