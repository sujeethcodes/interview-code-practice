package problems

import "fmt"

func FirstNonRepeatingChar() {
	str := "programming"
	freq := make(map[string]int)

	for _, each := range str {
		freq[string(each)]++
	}

	for _, each := range str {
		if freq[string(each)] == 1 {
			fmt.Println("First Non Repeating Character :", string(each))
			return
		}
	}
}
