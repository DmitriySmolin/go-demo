package main

import "fmt"

func main() {
	// key - значение
	// NY - NewYork
	m := map[string]string{
		"PurpleSchool": "purpleschool.ru",
	}
	fmt.Println(m)
	fmt.Println(m["PurpleSchool"])
	m["PurpleSchool"] = "https://purpleschool.ru"
	fmt.Println(m["PurpleSchool"])
	m["Google"] = "https://google.com"
	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m)
	delete(m, "Yandex")
	fmt.Println(m)
	delete(m, "Y")
	fmt.Println(m["Y"])
	fmt.Println(m)
}
