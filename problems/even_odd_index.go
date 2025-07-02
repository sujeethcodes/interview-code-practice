package problems

import "fmt"

func EvenOddIndex() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := []int{}
	odd := []int{}

	for i, each := range arr {
		if each%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}
	fmt.Println("Even value Index : ", even)
	fmt.Println("odd value Index : ", odd)

}
