package main

import (
	"fmt"
	"math/rand/v2"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.password = string(res)
}

func main() {

	// fmt.Println(generatePassword(12))

	login := promptData("Введите логин: ")
	// password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount := account{
		login: login,
		url:   url,
	}
	myAccount.generatePassword(12)
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
