package main

import (
	"fmt"
	"go-demo/account"
	"go-demo/files"
)

func main() {
	// 1. Создать аккаунт
	// 2. Найти аккаунт
	// 3. Удалить аккаунт
	// 4. Выход

	fmt.Println("__Менеджер паролей__")

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}

	}

}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&variant)
	return variant
}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL ")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}

	// myAccount.OutputPassword()
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func findAccount() {}

func deleteAccount() {}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
