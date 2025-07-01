package problems

import (
	"fmt"
	"strconv"
	"strings"
)

func SumDigits() {
	input := 12345
	splitVal := strings.Split(strconv.Itoa(input), "")
	sum := 0

	for _, each := range splitVal {
		num, _ := strconv.Atoi(each)
		sum += num
	}

	fmt.Println("Sum of Digits : ", sum)
}
