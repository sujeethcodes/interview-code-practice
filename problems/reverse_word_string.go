package problems

import "fmt"

func ReverseWordString() {
	input := "Golang coding Interview"
	n := len(input) - 1
	output := ""
	for i := n; i >= 0; i-- {
		output += string(input[i])
	}
	fmt.Println(output)
}
