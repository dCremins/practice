package main


func ZeroesMain(nums []int) []int {
	sorted := make([]int, len(nums))
	zCount := 0
	sInd := 0

	for _, num := range nums {
		if num == 0 {
			zCount += 1
		} else {
			sorted[sInd] = num
			sInd += 1
		}
	}
	for x:= 0; x<zCount;x++ {
		sorted[sInd] = 0
		sInd += 1
	}
	return sorted
}