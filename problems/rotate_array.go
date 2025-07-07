package problems

import (
	"fmt"
	"sort"
)

// Input:  arr = [1, 2, 3, 4, 5, 6, 7], k = 3
// Output: [5, 6, 7, 1, 2, 3, 4]

func RotateArray() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 2
	rotateVal := []int{}
	rotateVal = append(rotateVal, arr[len(arr)-k:]...)
	sort.Ints(rotateVal)
	rotateVal = append(rotateVal, arr[:len(arr)-k]...)
	fmt.Println(rotateVal)

}
