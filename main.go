package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {

	fmt.Println("__Калькулятор индекса массы тела__")
	for {
		userHeight, userKg := getUserInput()
		IMT := calculateIMT(userHeight, userKg)
		outputResult(IMT)

		isRepeatCalculation := checkRepeatCalculation()

		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще раз расчет (y/n): ")
	fmt.Scan(&userChoise)

	if userChoise == "y" || userChoise == "Y" {
		return true
	}

	return false
}
