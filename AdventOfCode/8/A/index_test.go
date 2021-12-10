package main

import (
	"testing"
	"reflect"
)


var data = []struct {
	input		string // input
	expected	string // expected result
  }{
	{
		"input",
		"expected",
	},
}

  func TestA(t *testing.T) {
	for _, testData := range data {
	  actual := AMain(testData.input)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("A()")
	  }
	}
  }