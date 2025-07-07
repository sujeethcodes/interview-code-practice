package problems

import "fmt"

func TwoSum() {
	input := []int{2, 7, 11, 15}
	target := 9

	left := 0
	right := len(input) - 1

	for left < right {
		sum := input[left] + input[right]
		if sum == target {
			fmt.Println(left, right)
			return
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	fmt.Println(-1)

}
