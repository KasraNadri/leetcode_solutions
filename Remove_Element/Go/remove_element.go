package main

func main() {
	// ----- ONLY HERE FOR TESTING ----- \\
}

// ========== THE SOLUTION FUNCTION ========== \\
func removeElement(nums []int, val int) int {
	refined_list := []int{}
	for _, v := range nums {
		if v != val {
			refined_list = append(refined_list, v) //ADDING ANY VALUE THAT'S NOT 'val'
		}
	}
	copy(nums, refined_list) //POURING THE 'refined_list' INTO 'nums' WHILE KEEPING 'nums' LENGTH
	return len(refined_list)
}
