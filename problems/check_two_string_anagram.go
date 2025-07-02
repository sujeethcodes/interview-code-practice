package problems

import (
	"fmt"
	"strings"
)

func CheckTwoStringAnagram() {

	str1 := "heart"
	str2 := "earth"

	str3 := strings.ToLower(str1 + str2)

	freq := make(map[string]int)

	for _, each := range str3 {
		freq[string(each)]++
	}

	for _, count := range freq {
		if count != 2 {
			fmt.Println(str1, "&", str2, false)
			return
		}
	}
	fmt.Println(str1, "&", str2, true)
}
