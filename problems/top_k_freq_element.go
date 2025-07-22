package problems

import (
	"fmt"
	"sort"
)

func TopKfreqElement() {

	arr := []int{1, 1, 1, 2, 2, 3}
	k := 2
	res := []int{}
	freq := make(map[int]int)

	for _, each := range arr {
		freq[each]++
	}

	type kv struct {
		Key   int
		Value int
	}
	frqSlice := []kv{}

	for key, value := range freq {
		if value > 1 {
			frqSlice = append(frqSlice, kv{key, value})
		}
	}

	sort.Slice(frqSlice, func(i, j int) bool {
		return frqSlice[i].Value > frqSlice[j].Value
	})

	for i := 0; i < k; i++ {
		res = append(res, frqSlice[i].Key)
	}
	fmt.Println(res)
}
