package problems

import "fmt"

func CountingVowelsConsonants() {
	input := "Automation World"
	vowles := "aeiouAEIOU"
	vowelStr := ""

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(vowles); j++ {
			if input[i] == vowles[j] {
				vowelStr += string(input[i])
				break
			}

		}
	}
	fmt.Println("Vowel Count      : ", len(vowelStr))
	fmt.Println("Non-Vowel Count  : ", len(input)-len(vowelStr))
}
