package main

import (
	"testing"
	"reflect"
)


var data = []struct {
	nums		[]int // input
	k			int	  // input
	expected	[]int // expected result
  }{
	{
		[]int{1,2,3,4,5,6,7},
		3,
		[]int{5,6,7,1,2,3,4},
	},
	{
		[]int{-1,-100,3,99},
		2,
		[]int{3,99,-1,-100},
	},
}

  func TestRotate(t *testing.T) {
	for _, testData := range data {
	  actual := RotateMain(testData.nums, testData.k)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("Rotate()")
	  }
	}
  }