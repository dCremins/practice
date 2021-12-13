package main

// import "math"

type Vent struct {
	x1	int
	y1	int
	x2	int
	y2	int
}

func AMain(vents []Vent) int {
	used := map[int]map[int]int{}
	crossed := 0
	for _, vent := range vents {
		if vent.x1 == vent.x2 {
			// Horizontal
			_, exists := used[vent.x1]
			if !exists {
				used[vent.x1] = map[int]int{}
			}
			min := vent.y1 
			max := vent.y2
			if vent.y1 > vent.y2 {
				min = vent.y2
				max = vent.y1
			}
			
			for i:=min; i<=max; i++ {
				_, exists = used[vent.x1][i]
				if !exists {
					used[vent.x1][i] = 0
				}
				used[vent.x1][i]++;
				if used[vent.x1][i]==2 {
					crossed++
				}
			}
		} else if vent.y1 == vent.y2 {
			// Vertical
			min := vent.x1 
			max := vent.x2
			if vent.x1 > vent.x2 {
				min = vent.x2
				max = vent.x1
			}
			for i:=min; i<=max; i++ {
				_, exists := used[i]
				if !exists {
					used[i] = map[int]int{}
				}
				_, exists = used[i][vent.y1]
				if !exists {
					used[i][vent.y1] = 0
				}
				used[i][vent.y1]++;
				if used[i][vent.y1]==2 {
					crossed++
				}
			}
		}
	}
	return crossed
}

/*
{
	1: {
		1: 1,
		2: 1,
		3: 0,
		4: 2,
	}
}
*/