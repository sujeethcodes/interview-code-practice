package problems

import (
	"fmt"
	"strings"
)

func RemoveAllWhiteSpace() {
	input := "No Spaces Here"
	inputField := strings.Fields(input)
	output := ""
	for _, each := range inputField {
		output += string(each)
	}
	fmt.Println("Remove White space :", output)
}
