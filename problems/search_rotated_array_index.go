package problems

import "fmt"

func SearchRotatedSortIndex() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			fmt.Println(mid)
			return
		}
		if arr[left] <= arr[mid] {
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= arr[right] && arr[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

}
