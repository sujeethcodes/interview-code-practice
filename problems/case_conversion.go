package problems

import (
	"fmt"
	"strings"
	"unicode"
)

func CaseConvertion() {
	input := "Golang Programming"
	output := ""
	for _, each := range input {
		if unicode.IsUpper(each) {
			output += strings.ToLower(string(each))
		} else {
			output += strings.ToUpper(string(each))
		}
	}
	fmt.Println("Case converted : ", output)
}
