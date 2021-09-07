package main

import (
	"testing"
	"reflect"
)


var data = []struct {
	n			int // input
	expected	int // expected result
  }{
	{ 5, 4 },
	{ 1, 1 },
}

  func TestFirstBad(t *testing.T) {
	for _, testData := range data {
	  actual := FirstBadMain(testData.n, testData.expected)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("FirstBad("+testData.n+"): "+testData.expected)
	  }
	}
  }