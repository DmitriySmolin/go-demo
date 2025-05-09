package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64 // 0.0
	var userKg float64     // 0.0

	fmt.Println("__Калькулятор индекса массы тела__")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес в кг: ")
	fmt.Scan(&userKg)
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userHeight float64, userKg float64) float64 {
	const IMTPower = 2

	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}
