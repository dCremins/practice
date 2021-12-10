package main

import (
    "strconv"
)

func AMain(nums []string)int64 {
	positions:=make([]int,len(nums[0]))
	gamma:=""
	epsilon:=""
	for _,b := range nums {
		for i:=0;i<len(nums[0]);i++ {
			if b[i]==49 {
				positions[i]+=1
			}
		}
	}
	for i:=0;i<len(nums[0]);i++ {
		if positions[i]>=(len(nums)/2) {
			gamma+="1"
			epsilon+="0"
		} else {
			gamma+="0"
			epsilon+="1"
		}
	}
	gRate,_:=strconv.ParseInt(gamma, 2, 64)
	eRate,_:=strconv.ParseInt(epsilon, 2, 64)
	return gRate*eRate
}