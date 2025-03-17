package main

func main() {
	//----- ONLY HERE FOR TESTING -----\\
}

// ========== THE SOLUTION FUNCTION ========== \\
func rotate(nums []int, k int) {
	rotated_array := []int{}
	if len(nums) > 1 {
		k = k % len(nums)
		for i := len(nums) - k; i <= len(nums)-1; i++ {
			rotated_array = append(rotated_array, nums[i])
		}
		rotated_array = append(rotated_array, nums[:len(nums)-k]...)
		copy(nums, rotated_array)
	}
}
