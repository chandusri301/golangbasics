package main

import (
	"fmt"
)

func isPalindrome(number int) bool {
	var remainder, temp int
	var reverse int = 0

	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {
		return true
	} else {
		return false
	}

}

func main() {
	var number int
	fmt.Println(" Enter any positive number")
	fmt.Scanln(&number)

	fmt.Println(" Result of palindrom=", isPalindrome(number))
}
