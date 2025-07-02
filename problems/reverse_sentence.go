package problems

import (
	"fmt"
	"strings"
)

func ReverseSentence() {
	input := "Golang is a programming language"
	words := strings.Fields(input)
	n := len(words) - 1
	output := ""
	for i := n; i >= 0; i-- {
		output += words[i] + " "
	}
	fmt.Println("Original :", input)
	fmt.Println("Reversed :", output)
}
