package main

import (
	"testing"
	"reflect"
)


var data = []struct {
	numbers		[]int // input
	target		int // input
	expected	[]int // expected result
  }{
	{
		[]int{2,7,11,15},
		9,
		[]int{1,2},
	},
	{
		[]int{2,3,4},
		6,
		[]int{1,3},
	},
	{
		[]int{-1,0},
		-1,
		[]int{1,2},
	},
	{
		[]int{-3,3,4,90},
		0,
		[]int{1,2},
	},
	{
		[]int{5,25,75},
		100,
		[]int{2,3},
	},
	{
		[]int{1,2,3,4,4,9,56,90},
		8,
		[]int{4,5},
	},
}

  func TestTwo(t *testing.T) {
	for i, testData := range data {
	  actual := TwoMain(testData.numbers, testData.target)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("Test %d expected %d, received %d", i, testData.expected, actual)
	  }
	}
  }