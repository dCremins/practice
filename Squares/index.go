package main

func SquaresMain(nums []int) []int {
	sorted := make([]int, len(nums))
    a := 0
    z := len(nums) - 1
	for i:=len(nums)-1; i>=0; i-- {
		squaredA := nums[a]*nums[a]
		squaredZ := nums[z]*nums[z]
		if squaredA > squaredZ {
			sorted[i] = squaredA
			a+=1
		} else {
			sorted[i] = squaredZ
			z-=1
		}
	}
	return sorted
}