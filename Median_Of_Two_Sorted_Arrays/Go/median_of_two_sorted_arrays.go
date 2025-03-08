package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println(median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	max_slice := []int{}
	min_slice := []int{}
	sum_slice := []int{}

	if len(nums1) >= len(nums2) {
		max_slice = nums1
		min_slice = nums2
	} else {
		max_slice = nums2
		min_slice = nums1
	}
	if len(min_slice) == 0 {
		sum_slice = max_slice
	} else {
		for i := 0; i < len(max_slice); {
			if min_slice != nil {
				if max_slice[i] > min_slice[0] {
					sum_slice = append(sum_slice, min_slice[0])
					min_slice = deleteFirst(min_slice)
				} else if max_slice[i] < min_slice[0] {
					sum_slice = append(sum_slice, max_slice[i])
					i++
				} else {
					sum_slice = append(sum_slice, max_slice[i], min_slice[0])
					min_slice = deleteFirst(min_slice)
					i++
				}
			} else {
				sum_slice = append(sum_slice, max_slice[i:]...)
				break
			}
		}
	}

	if min_slice != nil {
		sum_slice = append(sum_slice, min_slice...)
	}
	if len(sum_slice)%2 == 0 {
		return float64(sum_slice[len(sum_slice)/2-1]+sum_slice[len(sum_slice)/2]) / float64(2)
	} else {
		return float64(sum_slice[len(sum_slice)/2])
	}
}

func deleteFirst(slice []int) []int {
	if len(slice) > 1 {
		return slice[1:] // Return everything except the first element
	} else {
		return nil
	}
}
