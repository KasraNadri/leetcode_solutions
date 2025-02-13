package main

import (
	"strconv"
)

func main() {
	//----- ONLY HERE FOR TESTING -----\\
}

// ========== THE SOLUTION FUNCTION ==========\\
func isPalindrome(x int) bool {

	str_x := strconv.Itoa(x) //----- CONVERTING THE INTEGER TO A STRING -----\\

	len_x := len(str_x)

	digit_end := len_x - 1
	for digit_start := 0; digit_start <= len_x/2; digit_start++ {
		if str_x[digit_start] == str_x[digit_end] {
			digit_end--
		} else {
			return false
		}
	}
	return true
}
