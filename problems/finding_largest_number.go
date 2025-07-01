package problems

import "fmt"

func FindingLargestNumber() {
	arr := []int{10, 5, 25, 8, 15, 3}
	max := arr[0]

	for _, each := range arr {
		if each > max {
			max = each
		}
	}
	fmt.Println("Largest Number :", max)
}
