package problems

import "fmt"

func MergeTwoSortedArray() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8, 10}
	i, j := 0, 0
	merged := []int{}

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}
	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}
	fmt.Println(merged)

}
