package main

func main() {
	//----- ONLY HERE FOR TESTING -----\\
}

func longestCommonPrefix(strs []string) string {

	// ===== FINDS THE SHORTEST STRING IN THE LIST ===== \\
	smallest_length := len(strs[0])
	for _, val := range strs {
		if len(val) < smallest_length {
			smallest_length = len(val)
		}
	}

	var outer_str string

	for i := 0; i < smallest_length; i++ {
		inner_str := string(strs[0][i])
		for _, value := range strs {
			if string(value[i]) != inner_str {
				return outer_str
			}
		}
		outer_str += inner_str
	}
	return outer_str
}
