package problems

import "fmt"

func CheckArraySort() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				fmt.Println("unsort array : ", false)
				return
			}
		}
	}
	fmt.Println("sort array : ", true)
}
