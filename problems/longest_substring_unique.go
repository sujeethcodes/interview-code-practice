package problems

import (
	"fmt"
)

// Input: "abcabcbb" → Output: 3 (abc)

func LongestSubstringUnique() {
	charIndex := make(map[byte]int)
	char := "abcabcbb"
	maxLen := 0
	start := 0

	for i := 0; i < len(char); i++ {
		index, found := charIndex[char[i]]
		if found && index >= start {
			start = index + 1
		}
		charIndex[char[i]] = i
		currentLength := i - start + 1
		if currentLength > maxLen {
			maxLen = currentLength

		}
	}

	fmt.Println("Len :", maxLen)

}
