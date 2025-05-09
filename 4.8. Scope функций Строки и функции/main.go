package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")
	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userHeight float64, userKg float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, IMTPower)

	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64 // 0.0
	var userKg float64     // 0.0

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес в кг: ")
	fmt.Scan(&userKg)

	return userHeight, userKg
}

// func calculateIMT(userHeight float64, userKg float64) (IMT float64) {

// 	IMT = userKg / math.Pow(userHeight/100, IMTPower)
// 	return
// }

// func getUserInput() (userHeight float64, userKg float64) {

// 	fmt.Print("Введите свой рост в сантиметрах: ")
// 	fmt.Scan(&userHeight)
// 	fmt.Print("Введите ваш вес в кг: ")
// 	fmt.Scan(&userKg)

// 	return
// }
