package problems

import (
	"fmt"
	"sort"
)

func PairSumFinder() {
	arr := []int{2, 4, 3, 5, 6, -2, 8, 7, 1}
	sort.Ints(arr)
	target := 6
	output := [][]int{}

	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			output = append(output, []int{arr[left], arr[right]})
			left++
			right--
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	fmt.Println("Target :", target)
	fmt.Println("Output :", output)

}
