package problems

import "fmt"

func PalindromeChecker() {
	palindromeValue := "chair"
	left := 0
	right := len(palindromeValue) - 1

	for left < right {
		if palindromeValue[left] != palindromeValue[right] {
			fmt.Println(false)
			return
		}
		left++
		right--
	}
	fmt.Println(true)

}
