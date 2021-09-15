package main

func TwoMain(numbers []int, target int) []int {
	answer := make([]int, 2)
	for i:=0;i<len(numbers);i++ {
		for j:=i+1;j<len(numbers);j++ {
			if numbers[i]+numbers[j] == target {
				answer[0]=i+1
				answer[1]=j+1
				return answer
			}
		}
	}
	return answer
}