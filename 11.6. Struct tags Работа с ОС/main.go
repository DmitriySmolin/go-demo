package main

import (
	"fmt"
	"go-demo/account"
	"go-demo/files"
)

func main() {
	// defer fmt.Println(1) // 1
	// defer fmt.Println(2) // 2
	// порядок вывода 2 -> 1 // по приниципу LIFO работает defer
	files.ReadFile()
	files.WriteFile("Привет! Я файл!!", "file.txt")

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL ")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}

	myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
