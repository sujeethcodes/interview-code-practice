package problems

import (
	"fmt"
	"strconv"
)

func StringToInt() {
	str := "12345"
	int, _ := strconv.Atoi(str)
	fmt.Println(int)
}
