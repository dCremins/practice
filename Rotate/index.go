package main
// import "fmt"

func RotateMain(nums []int, k int) []int {
	rotated := make([]int, len(nums))
	for i, num := range nums {
		rotated[(i+k)%len(nums)] = num
	}
	return rotated
}