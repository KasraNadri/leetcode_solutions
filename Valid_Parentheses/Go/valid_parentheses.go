package main

func main() {
	// ----- ONLY HERE FOR TESTING ----- \\
}

// ========== THE SOLUTION FUNCTION ========== \\
func isValid(s string) bool {
	symbols := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{} // Stack to store opening brackets

	for _, char := range s {

		if closing, exists := symbols[char]; exists {
			stack = append(stack, closing)
		} else {

			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			// Pop from stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
