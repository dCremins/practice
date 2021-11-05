package main

import (
	"testing"
	"reflect"
)


var data = []struct {
	input1		[]int // input
	input2		[]int // input
	expected	float64 // expected result
  }{
	{
		[]int{1,3},
		[]int{2},
		2.00000,
	},
	{
		[]int{1,2},
		[]int{3,4},
		2.50000,
	},
	{
		[]int{0,0},
		[]int{0,0},
		0.00000,
	},
	{
		[]int{},
		[]int{1},
		1.00000,
	},
	{
		[]int{2},
		[]int{},
		2.00000,
	},
}

  func TestMedian(t *testing.T) {
	for _, testData := range data {
	  actual := MedianMain(testData.input1, testData.input2)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("Median()")
	  }
	}
  }