package problems

import "fmt"

func MoveZeroToEnd() {
	arr := []int{1, 3, 0, 2, 5, 0, 4, 0, 6, 0}

	pos := 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != 0 {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos++
		}
	}
	fmt.Println(arr)

}
