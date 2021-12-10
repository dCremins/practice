package main

import (
	// "fmt"
	"strconv"
)

func filterVals(vals []string, index int, t string)[]string {
	if len(vals) == 1 {
		return vals
	}
	O := []string{}
	I := []string{}
	for _,v := range vals {
		if v[index] == 49 {
			I = append(I, v)
		} else {
			O = append(O, v)
		}
	}
	if t == "most" {
		if len(I) >= len(O) {
			if len(I) == 1 {
				return I
			}
			return filterVals(I, index+1, t)
		} else if len(O) == 1 {
			return O
		}
		return filterVals(O, index+1, t)
	}

	if len(I) < len(O) {
		if len(I) == 1 {
			return I
		}
		return filterVals(I, index+1, t)
	} else if len(O) == 1 {
		return O
	}
	return filterVals(O, index+1, t)
}

func BMain(nums []string)int64 {
	O2 := filterVals(nums, 0, "most")
	CO2 := filterVals(nums, 0, "least")
	// fmt.Println(O2)
	// fmt.Println(CO2)
	
	oRate,_:=strconv.ParseInt(O2[0], 2, 64)
	cRate,_:=strconv.ParseInt(CO2[0], 2, 64)
	return oRate*cRate
}
