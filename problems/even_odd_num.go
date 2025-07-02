package problems

import "fmt"

func EvenOddFinder() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even := []int{}
	odd := []int{}
	for _, each := range arr {
		if each%2 == 0 {
			even = append(even, each)
		} else {
			odd = append(odd, each)
		}
	}

	fmt.Println("Even :", even)
	fmt.Println("Odd :", odd)

}
