package problems

import "fmt"

func SortArray() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original Array :", arr)

	n := len(arr) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Sorted Array :", arr)

}
