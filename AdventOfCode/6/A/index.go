package main

func AMain(fish []int)int {
	for i := 0; i<80; i++ {
		// after 80 days
		newFish := []int{}
		for j,f := range fish {
			if f == 0 {
				fish[j] = 6
				newFish = append(newFish, 8)
			} else {
				fish[j] -= 1
			}
		}
		fish = append(fish, newFish...)
	}
	return len(fish)
}