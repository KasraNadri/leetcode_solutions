package main

func main() {
	//----- ONLY HERE FOR TESTING -----\\
}

// ========== THE SOLUTION FUNCTION ========== \\
func strStr(haystack string, needle string) int {
	i := 0
	for i <= len(haystack)-len(needle) {
		if haystack[i] == needle[0] {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
		i++
	}
	return -1
}
