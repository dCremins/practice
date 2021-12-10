package main


func OneMain(depths []int) int {
	last := depths[0]
	count := 0

	for _,d := range depths {
		if d > last {
			count++
		}
		last = d
	}
	return count
}