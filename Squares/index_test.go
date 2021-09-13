package main

import (
	"testing"
	"reflect"
)

var data = []struct {
	input		[]int // input
	expected	[]int // expected result
  }{
	{
		[]int{-4, -1, 0, 3, 10},
		[]int{0, 1, 9, 16, 100},
	},
	{
		[]int{-7, -3, 2, 3, 11},
		[]int{4, 9, 9, 49, 121},
	},
}

  func TestSquares(t *testing.T) {
	for _, testData := range data {
	  actual := SquaresMain(testData.input)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("Squares()")
	  }
	}
  }