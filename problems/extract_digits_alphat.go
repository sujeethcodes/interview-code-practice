package problems

import (
	"fmt"
	"unicode"
)

func ExtractDigitsAndAlphat() {
	input := "abc123def456ghi789"
	digit := ""
	alphat := ""
	for _, each := range input {
		if unicode.IsDigit(each) {
			digit += string(each)
		} else {
			alphat += string(each)
		}
	}
	fmt.Println("Original Inputs :", input)
	fmt.Println("Extracted digits :", digit)
	fmt.Println("Extracted alphat :", alphat)

}
