package problems

import "fmt"

func FindMissingElements() {
	arr := []int{1, 2, 4, 6, 3, 7, 8}
	found := make(map[int]bool)
	n := arr[len(arr)-1]

	for _, each := range arr {
		found[each] = true
	}

	for i := 1; i <= n; i++ {
		if !found[i] {
			fmt.Println("Missing number : ", i)
		}
	}

}
