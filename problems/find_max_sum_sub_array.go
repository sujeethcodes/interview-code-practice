package problems

import "fmt"

func FindTheMaximumSumOfSubArray() {
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3
	// output = 9
	windows := 0
	for i := 0; i < k; i++ {
		windows += arr[i]
	}
	maxSum := windows
	for j := k; j < len(arr)-1; j++ {
		windows = windows - arr[j-k] + arr[j]
		if windows > maxSum {
			maxSum = windows
		}
	}
	fmt.Println(maxSum)

}
