package problems

import "fmt"

func RemoveElements() {
	arr := []int{3, 2, 2, 3}
	val := 3

	pos := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != val {
			arr[pos] = arr[i]
			pos++

		}

	}
	fmt.Println("length :", len(arr[:pos]))
}
