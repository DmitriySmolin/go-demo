package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {

	// i := 0
	// for i < 10 {
	// 	fmt.Printf("%d\n", i)
	// 	i++
	// }

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	fmt.Println("__Калькулятор индекса массы тела__")
	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
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
