package main

func main() {
	// ----- ONLY HERE FOR TESTING ----- \\
}

// ========== THE OLUTION FUNCTION ========== \\
func twoSum(nums []int, target int) []int {
	indexes := []int{}
	for index1, value1 := range nums {
		for index2, value2 := range nums {
			if index1 != index2 && value1+value2 == target {
				indexes = append(indexes, index1, index2)
				return indexes
			}
		}
	}
	return indexes
}
