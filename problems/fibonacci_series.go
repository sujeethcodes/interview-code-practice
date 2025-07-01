package problems

import "fmt"

func Fibonacci() {
	a, b := 0, 1
	val := ""
	for i := 0; i <= 10; i++ {
		val += fmt.Sprint(a, " ")
		a, b = a+b, a
	}
	fmt.Println(val)
}
