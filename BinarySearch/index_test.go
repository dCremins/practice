package main

import (
	"testing"
	"reflect"
)


var data = []struct {
	nums		[]int // input
	target		int   // input
	expected	int   // expected result
  }{
	{ []int{-1,0,3,5,9,12}, 9, 4 },
	{ []int{-1,0,3,5,9,12}, 2, -1 },
}

  func TestBinarySearch(t *testing.T) {
	for _, testData := range data {
	  actual := BinarySearchMain(testData.nums, testData.target)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("BinarySearchMain()")
	  }
	}
  }