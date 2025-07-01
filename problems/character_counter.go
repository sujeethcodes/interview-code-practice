package problems

import "fmt"

func CharacterCounter() {
	str := "developer"
	chrMap := make(map[string]int)

	for _, each := range str {
		chrMap[string(each)]++
	}
	fmt.Println(chrMap)
}
