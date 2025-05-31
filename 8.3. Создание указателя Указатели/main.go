package main

import "fmt"

// Указатель - переменные, которые хранят адрес в памяти, а не значение

func main() {
	a := 5
	pointerA:= &a
	res := double(a)
	fmt.Println(pointerA)
	fmt.Println(res)
}

func double(num int) int {
	return num * 2
}
