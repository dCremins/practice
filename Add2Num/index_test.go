package main

import (
	"testing"
	"reflect"
)

var linkedData = []struct {
	L1			*ListNode // input
	L2			*ListNode// input
	expected	*ListNode // expected result
  }{
	{
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:4,
				Next: &ListNode{
					Val:3,
				},
			},
		}, 
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val:6,
				Next: &ListNode{
					Val:4,
				},
			},
		},
		&ListNode{
			Val: 7,
			Next: &ListNode{
				Val:0,
				Next: &ListNode{
					Val:8,
				},
			},
		},
	},
	{
		&ListNode{
			Val: 0,
		}, 
		&ListNode{
			Val: 0,
		}, 
		&ListNode{
			Val: 0,
		},
	},
	{
		&ListNode{
			Val: 9,
			Next: &ListNode{
				Val:9,
				Next: &ListNode{
					Val:9,
					Next: &ListNode{
						Val:9,
						Next: &ListNode{
							Val:9,
							Next: &ListNode{
								Val:9,
								Next: &ListNode{
									Val:9,
								},
							},
						},
					},
				},
			},
		}, 
		&ListNode{
			Val: 9,
			Next: &ListNode{
				Val:9,
				Next: &ListNode{
					Val:9,
					Next: &ListNode{
						Val:9,
					},
				},
			},
		}, 
		&ListNode{
			Val: 8,
			Next: &ListNode{
				Val:9,
				Next: &ListNode{
					Val:9,
					Next: &ListNode{
						Val:9,
						Next: &ListNode{
							Val:0,
							Next: &ListNode{
								Val:0,
								Next: &ListNode{
									Val:0,
									Next: &ListNode{
										Val:1,
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

  func TestAdd(t *testing.T) {
	for _, testData := range linkedData {
	  actual := addTwoNumbers(testData.L1, testData.L2)
	  if !reflect.DeepEqual(actual, testData.expected) {
		t.Errorf("addTwoNumbers()")
	  }
	}
  }