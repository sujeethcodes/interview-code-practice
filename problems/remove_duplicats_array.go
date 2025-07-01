package problems

import "fmt"

func RemoveDuplicatsArray() {
	arr := []int{1, 2, 3, 2, 5, 1, 6, 3, 7}
	seen := make(map[int]bool)
	duplicate := make(map[int]bool)
	res := []int{}
	for _, each := range arr {
		if !seen[each] {
			seen[each] = true
			res = append(res, each)
		} else {
			duplicate[each] = true
		}
	}
	fmt.Println(res)

}
