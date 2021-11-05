package main

import "fmt"
import "sort"


func MedianMain(nums1 []int, nums2 []int) float64 {
	combined := append(nums1, nums2...)
	sort.Ints(combined)
	length := len(combined)
	
	if length % 2 != 0 {
		return float64(combined[(length-1)/2])
	}
	mid := length/2
	return float64(combined[mid]+combined[mid-1])/2
}