package main

import (
	"fmt"
    "strings"
	"strconv"
)

func main() {
	AddLinkedList([]int{2,4,3}, []int{5,6,4})
	AddLinkedList([]int{0}, []int{0})
	AddLinkedList([]int{9,9,9,9,9,9,9}, []int{9,9,9,9})
}

func Reverse(list []int) []int {
	reverseList := make([]int, 0)
	for i := range list {
		reverseList = append(reverseList, list[len(list) - i - 1])
	}
	return reverseList
}

func Concat(list []int) int {
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

func Split(n int) []int {
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

func AddLinkedList(L1 []int, L2 []int) []int {
    first := Reverse(L1)
	second := Reverse(L2)
	firstNum := Concat(first)
	secondNum := Concat(second)
	total := firstNum + secondNum
	splitTotal := Split(total)
	return splitTotal
}