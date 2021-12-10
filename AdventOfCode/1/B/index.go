package main

import "fmt"

func TwoMain(depths []int) int {
	windows := make([]int, len(depths))
	count := 0

	for i,d := range depths { 
		if i > 2 {
			windows[i-2]+=d 
			windows[i-1]+=d 
			windows[i]+=d
		} else if i > 1 {
			windows[i-1]+=d 
			windows[i]+=d
		} else {
			windows[i]+=d
		}
		if i > 3 && windows[i-2] > windows[i-3] {
			fmt.Println(windows[i-2], windows[i-3])
			count++
		}
	}
	return count
}