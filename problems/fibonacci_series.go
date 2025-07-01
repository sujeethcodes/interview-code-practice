package problems

import "fmt"

func Fibonacci() {
	a, b := 0, 1
	for i := 0; i <= 10; i++ {
		fmt.Println(a, " ")
		a, b = a+b, a

	}
}
