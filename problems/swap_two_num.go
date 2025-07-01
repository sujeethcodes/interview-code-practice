package problems

import "fmt"

func SwapTwoNumbers() {
	a := 10
	b := 20

	temp := a
	a = b
	b = temp
	fmt.Println("a-> ", a, "b-> ", b)

}
