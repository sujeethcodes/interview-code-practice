package problems

import "fmt"

func CommonElementsOfTwoAray() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4, 5, 6, 7}
	arr3 := append(arr1, arr2...)
	freq := make(map[int]int)
	output := []int{}
	for _, each := range arr3 {
		freq[each]++
	}

	for num, count := range freq {
		if count != 1 {
			output = append(output, num)
		}
	}
	fmt.Println(output)

}
