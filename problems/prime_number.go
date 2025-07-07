package problems

import "fmt"

func PrimeNumbers() {

	for i := 0; i <= 10; i++ {
		isPrime := true

		if i <= 1 {
			isPrime = false
		}

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i, "is Prime")
		} else {
			fmt.Println(i, "is Not Prime")
		}
	}
}
