package main
// import "fmt"

func babyMaker(fish []int, day int) int {
	// fmt.Println(day)
	if day == 0 {
		sum := 0
		for _, count := range fish {
			sum += count
		}
		return sum
	}

	school := []int{0,0,0,0,0,0,0,0,0}
	for i,c := range fish {
		if i == 0 {
			school[6] += c
			school[8] += c
		} else {
			school[i-1] +=c
		}
	}
	return babyMaker(school, day-1)
}

func BMain(fish []int)int {
	school := []int{0,0,0,0,0,0,0,0,0}
	for _, age := range fish {
		school[age]++
	}
	return babyMaker(school, 256)
}