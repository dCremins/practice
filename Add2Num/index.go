package main

import (
	"fmt"
    "strings"
	"strconv"
)

func reverse(list []int) []int {
	reverseList := make([]int, 0)
	for i := range list {
		reverseList = append(reverseList, list[len(list) - i - 1])
	}
	return reverseList
}

func concat(list []int) int {
	textSlice := []string{}
	for _, n := range list {
        text := strconv.Itoa(n)
        textSlice = append(textSlice, text)
    }
	numString := strings.Join(textSlice, "")
	number, err := strconv.Atoi(numString)
	if err != nil {
		fmt.Println("err", err)
	}
	return number
}

func split(n int) []int {
	splitList := make([]int, 0)

	if n == 0 {
		splitList = append(splitList, 0)
	}

	for n != 0 {
        splitList = append(splitList, n % 10)
        n /= 10
    }
	return splitList
}

func addLinkedList(L1 []int, L2 []int) []int {
    first := reverse(L1)
	second := reverse(L2)
	firstNum := concat(first)
	secondNum := concat(second)
	total := firstNum + secondNum
	splitTotal := split(total)
	return splitTotal
}