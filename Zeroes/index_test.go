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
		[]int{0, 1, 0, 3, 12},
		[]int{1, 3, 12, 0, 0},
	},
	{
		[]int{0},
		[]int{0},
	},
}

  func TestZeroes(t *testing.T) {
	for _, testData := range data {
	  actual := ZeroesMain(testData.input)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("Zeroes()")
	  }
	}
  }