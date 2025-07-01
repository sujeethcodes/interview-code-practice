package problems

import "fmt"

func ReverseString(s string) {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	fmt.Println(res)
}
