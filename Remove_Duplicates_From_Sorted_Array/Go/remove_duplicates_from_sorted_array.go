package main

func main() {
	// ----- ONLY HERE FOR TESTING ----- \\
}

// ========== THE SOLUTION FUNCTION ========== \\
func removeDuplicates(nums []int) int {
	pureNums := map[int]int{}
	newNums := []int{}
	for _, v := range nums {
		if _, exists := pureNums[v]; exists {
			pureNums[v] += 1
		} else {
			newNums = append(newNums, v)
			pureNums[v] = 1
		}
	}
	copy(nums, newNums)
	return len(newNums)
}
