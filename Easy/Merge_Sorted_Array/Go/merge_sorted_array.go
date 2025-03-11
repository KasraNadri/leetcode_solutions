package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	nums1_ := trimTrailingZeros(nums1)
	nums2_ := trimTrailingZeros(nums2)
	merged_array := []int{}
	for i < len(nums1_) || j < len(nums2_) {
		if i < len(nums1_) && j < len(nums2_) {
			if nums1_[i] > nums2_[j] {
				merged_array = append(merged_array, nums2_[j])
				j++
			} else if nums1_[i] < nums2_[j] {
				merged_array = append(merged_array, nums1_[i])
				i++
			} else {
				merged_array = append(merged_array, nums1_[i], nums2_[j])
				i++
				j++
			}
		}
		if i >= len(nums1_) && j < len(nums2_) {
			merged_array = append(merged_array, nums2_[j])
			j++
		}
		if j >= len(nums2_) && i < len(nums1_) {

			merged_array = append(merged_array, nums1_[i])

			i++
		}
	}

	copy(nums1, merged_array)
}

func trimTrailingZeros(s []int) []int {
	i := len(s) - 1
	for i >= 0 && s[i] == 0 {
		i--
	}
	return s[:i+1]
}
