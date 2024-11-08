package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a float64
	var b float64
	var c float64
	var d float64
	fmt.Print("Введите коэффициент a: ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Print("Введите коэффициент b: ")
	fmt.Fscan(os.Stdin, &b)
	fmt.Print("Введите коэффициент c: ")
	fmt.Fscan(os.Stdin, &c)
	if a == 0 {
		fmt.Print("Корень уравнения: ", -c/b)
	} else {
		d = b*b - 4*a*c
		if d < 0 {
			fmt.Print("Нет корней")
		} else if d == 0 {
			fmt.Print("Корень уравнения: ", -b/(2*a))
		} else {
			fmt.Println("Первый корень уравнения: ", (-b+math.Sqrt(d))/(2*a))
			fmt.Print("Второй корень уравнения: ", (-b-math.Sqrt(d))/(2*a))
		}
	}
}
