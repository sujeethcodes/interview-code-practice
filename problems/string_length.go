package problems

import "fmt"

func FindStringLength() {
	input := "hello world"
	count := 0
	for range input {
		count++
	}
	fmt.Println(count)
}
