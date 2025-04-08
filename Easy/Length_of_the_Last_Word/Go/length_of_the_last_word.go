package main

//----- REQUIRED IMPORTS -----\\
import (
	"strings"
)

func main() {
	//----- ONLY HERE FOR TESTING -----\\
}

// ========== THE SOLUTION FUNCTION ========== \\
func lengthOfLastWord(s string) int {
	trimmed_s := strings.TrimSpace(s)
	sliced_s := strings.Fields(trimmed_s)
	return len(sliced_s[len(sliced_s)-1])
}
