package problems

import "fmt"

func SecondLargestElement() {
	arr := []int{12, 35, 1, 10, 34, 1}
	max := arr[0]
	secondMax := 0

	for _, each := range arr {
		if each > max {
			secondMax = max
			max = each
		} else if each > secondMax {
			secondMax = each
		}
	}
	fmt.Println("Second Largest Element : ", secondMax)
}
