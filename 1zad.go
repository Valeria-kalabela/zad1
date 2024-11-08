package main

import (
	"fmt"
	"os"
	"strconv"
)

func des(n, y int) string {
	var a string
	for n != 0 {
		a = strconv.Itoa(n%y) + a
		n = n / y
	}
	return a
}

func vdes(n, x int) int {
	var a int = 0
	var i int = 0
	for n != 0 {
		a += ((n % 10) * (x ^ i))
		n = n / 10
		i++
	}
	return a
}

func main() {
	var n, x, y int
	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &n)
	fmt.Print("Введите начальную систему счисления: ")
	fmt.Fscan(os.Stdin, &x)
	fmt.Print("Введите конечную систему счсиления: ")
	fmt.Fscan(os.Stdin, &y)
	if x == 10 {
		fmt.Print("Ответ: ", des(n, y))
	} else {
		fmt.Print(vdes(n, x))
	}
}
