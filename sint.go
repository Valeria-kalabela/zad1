/* package main

import "fmt"

func main() {
	var hello string = "hello world"
	hello = "hello GO"
	fmt.Println(hello)
}









package main

import "fmt"

var n int64
var numbers [5]int = [5]int{1, 2, 3, 4, 5}
var numbers1 = [...]int{1, 2, 3, 4, 5}
var numbers2 = []int{1, 2, 3, 4, 5}

func main() {
	numbers2 = append(numbers2, 6)
	//fmt.Scan(&n)
	fmt.Println(numbers2)
}


package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i * i)
	}

	var i = 10
	for i < 20 {
		fmt.Println(i * i)
		i++
	}
}





ЗАПУСК ПРОГРАММЫ
go run main.go

ДЛЯ КОМПИЛЯЦИИ
go build 1zad.go


//var x=6
//var n int32 = 5


var hello string
hello = "Hello world"

\n: переход на новую строку
\t табуляция, например fmt.Print(i * j, "\t")



const (
    pi float64 = 3.1415
    e float64 = 2.7182
)






SWITCH

package main
import "fmt"

func main() {

    a := 8
    switch(a) {
        case 9:
            fmt.Println("a = 9")
        case 8:
            fmt.Println("a = 8")
        case 7,6,5:
            fmt.Println("a = 7")
		default:
            fmt.Println("значение переменной a не определено")
    }
}






ПЕРЕБОР МАССИВОВ
var users = [3]string{"Tom", "Alice", "Kate"}
for index, value := range users{
    fmt.Println(index, value)
}

ИЛИ

var users = [3]string{"Tom", "Alice", "Kate"}
for i:= 0; i < len(users); i++{
    fmt.Println(users[i])
}











Параметры функции(Если несколько параметров подряд имеют один и тот
 же тип,то мы можем указать тип только для последнего параметра, а
 предыдущие параметры также будут представлять этот тип:)

package main
import "fmt"

func main() {
    add(4, 5)   // x + y = 9
    add(20, 6)  // x + y = 26
}

func add(x int, y int){
    var z = x + y
    fmt.Println("x + y = ", z)
}








Неопределенное количество параметров

package main
import "fmt"

func main() {
    add(1, 2, 3)        // sum = 6
    add(1, 2, 3, 4)     // sum = 10
    add(5, 6, 7, 2, 3)  // sum = 23
}

func add(numbers ...int){
    var sum = 0
    for _, number := range numbers{
        sum += number
    }






Если мы хотим передать срез
add([]int{1, 2, 3}...)

add([]int{1, 2, 3, 4}...)

var nums = []int{5, 6, 7, 2, 3}
add(nums...)





Возвращение результата из функции
func add(x, y int) int {
    return x + y
}



func add(x, y int, firstName, lastName string) (int , string) {
    var z int = x + y
    var fullName = firstName + " " + lastName
    return z, fullName
}
*/