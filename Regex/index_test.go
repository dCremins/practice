package main

import (
	"testing"
	"reflect"
)


var data = []struct {
	input1		string // input
	input2		string // input
	expected	bool // expected result
  }{
	{
		"aa",
		"a",
		false,
	},
	{
		"aa",
		"a*",
		true,
	},
	{
		"ab",
		".*",
		true,
	},
	{
		"aab",
		"c*a*b",
		true,
	},
	{
		"mississippi",
		"mis*is*p*",
		false,
	},
}

  func TestRegex(t *testing.T) {
	for _, testData := range data {
	  actual := RegexMain(testData.input1, testData.input2)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("Regex()")
	  }
	}
  }