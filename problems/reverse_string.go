package problems

import "fmt"

func ReverseString() {
	reversetr := "helloworld"
	res := ""
	for i := len(reversetr) - 1; i >= 0; i-- {
		res += string(reversetr[i])
	}
	fmt.Println(res)
}
