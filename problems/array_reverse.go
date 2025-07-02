package problems

import "fmt"

func ArrayReverse() {
	arr := []int{1, 2, 3, 4, 5}
	output := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		output = append(output, arr[i])
	}
	fmt.Println("Array Reverse Value : ", output)
}
