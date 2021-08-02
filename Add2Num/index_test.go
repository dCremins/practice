package main

import (
	"testing"
	"reflect"
)

var reverseData = []struct {
	list		[]int // input
	expected	[]int // expected result
  }{
	{[]int{2,4,3}, []int{3,4,2}},
	{[]int{5,2,7,8,1}, []int{1,8,7,2,5}},
	{[]int{0}, []int{0}},
}


var concatData = []struct {
	list		[]int // input
	expected	int // expected result
  }{
	{[]int{2,4,3}, 243},
	{[]int{0}, 0},
	{[]int{9,9,9,9,9,9,9}, 9999999},
}


var splitData = []struct {
	n			int // input
	expected	[]int // expected result
  }{
	{243, []int{3,4,2}},
	{13579, []int{9,7,5,3,1}},
	{0, []int{0}},
	{9999999, []int{9,9,9,9,9,9,9}},
}


var linkedData = []struct {
	L1			[]int // input
	L2			[]int // input
	expected	[]int // expected result
  }{
	{[]int{2,4,3}, []int{5,6,4}, []int{7,0,8}},
	{[]int{0}, []int{0}, []int{0}},
	{[]int{9,9,9,9,9,9,9}, []int{9,9,9,9}, []int{8,9,9,9,0,0,0,1}},
}

  func TestReverse(t *testing.T) {
	for _, testData := range reverseData {
	  actual := Reverse(testData.list)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("reverse(%d): expected %d, actual %d", testData.list, testData.expected, actual)
	  }
	}
  }

  func TestConcat(t *testing.T) {
	for _, testData := range concatData {
		actual := Concat(testData.list)
	  if actual != testData.expected {
		t.Errorf("concat(%d): expected %d, actual %d", testData.list, testData.expected, actual)
	  }
	}
  }

  func TestSplit(t *testing.T) {
	for _, testData := range splitData {
		actual := Split(testData.n)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("split(%d): expected %d, actual %d", testData.n, testData.expected, actual)
	  }
	}
  }

  func TestAdd(t *testing.T) {
	for _, testData := range linkedData {
	  actual := AddLinkedList(testData.L1, testData.L2)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("addLinkedList(%d, %d): expected %d, actual %d", testData.L1, testData.L2, testData.expected, actual)
	  }
	}
  }