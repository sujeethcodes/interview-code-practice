package problems

import "fmt"

func FindingSamllestNumber() {
	arr := []int{10, 5, 25, 8, 15, 3}

	min := arr[0]
	for _, each := range arr {
		if each < min {
			min = each
		}
	}
	fmt.Println("Samllest Number :", min)
}
