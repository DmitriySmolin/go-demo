package main

import "fmt"

// Указатель - переменные, которые хранят адрес в памяти, а не значение

func main() {
	a := 5
	double(&a)
	fmt.Println(a)
}

func double(num *int) {
	*num = *num * 2
}
