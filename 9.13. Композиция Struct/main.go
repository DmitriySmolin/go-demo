package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

// Zoo
// Animal-> Crocodile

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
	// acc account
}

func newAccountWithTimeStamp(login string, password string, urlString string) (*accountWithTimeStamp, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}

	if password == "" {
		newAcc.generatePassword(12) // более упрощенный вызов встроенного метода
		// newAcc.account.generatePassword(12) // допустимое обращение при композиции
	}

	return newAcc, nil
}

// 1. Если логина нет, то ошибка
// 2. Если нет пароля, то генерим
// func newAccount(login string, password string, urlString string) (*account, error) {

// 	if login == "" {
// 		return nil, errors.New("INVALID_LOGIN")
// 	}

// 	_, err := url.ParseRequestURI(urlString)

// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}

// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}

// 	if password == "" {
// 		newAcc.generatePassword(12)
// 	}

// 	return newAcc, nil
// }

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

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

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL ")

	myAccount, err := newAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}

	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
