package main

import (
	"fmt"
	"go-demo/account"

	"github.com/fatih/color"
)

func main() {
	// 1. Создать аккаунт
	// 2. Найти аккаунт
	// 3. Удалить аккаунт
	// 4. Выход

	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
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

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL ")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount)

}

func findAccount(vault *account.Vault) {
	// URL
	// Поиск аккаунтов
	// Вывод
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount() {}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
