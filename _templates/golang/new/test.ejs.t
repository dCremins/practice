---
to: <%= h.changeCase.pascal(name) %>/index_test.go
---
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

  func Test<%= h.changeCase.pascal(name) %>(t *testing.T) {
	for _, testData := range data {
	  actual := <%= h.changeCase.pascal(name) %>Main(testData.input)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("<%= h.changeCase.pascal(name) %>()")
	  }
	}
  }