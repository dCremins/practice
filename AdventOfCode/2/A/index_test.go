package main

import (
	"testing"
	"reflect"
	"fmt"
)


var data = []struct {
	input		[]Move // input
	expected	int // expected result
  }{
	{
		[]Move{
			{angle:"forward", amount:5},
			{angle:"down", amount:5},
			{angle:"forward", amount:8},
			{angle:"up", amount:3},
			{angle:"down", amount:8},
			{angle:"forward", amount:2},
		},
		150,
	},
	{
		[]Move{
			{angle:"forward", amount: 4},
			{angle:"down", amount: 8},
			{angle:"down", amount: 8},
			{angle:"up", amount: 2},
			{angle:"up", amount: 7},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 7},
			{angle:"down", amount: 6},
			{angle:"down", amount: 3},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 9},
			{angle:"up", amount: 2},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 4},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 4},
			{angle:"up", amount: 3},
			{angle:"down", amount: 9},
			{angle:"up", amount: 8},
			{angle:"down", amount: 2},
			{angle:"down", amount: 6},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 6},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 9},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 3},
			{angle:"down", amount: 8},
			{angle:"up", amount: 2},
			{angle:"down", amount: 3},
			{angle:"up", amount: 4},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 7},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 3},
			{angle:"down", amount: 6},
			{angle:"down", amount: 6},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 1},
			{angle:"up", amount: 6},
			{angle:"up", amount: 7},
			{angle:"forward", amount: 8},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 5},
			{angle:"up", amount: 3},
			{angle:"up", amount: 4},
			{angle:"down", amount: 9},
			{angle:"up", amount: 4},
			{angle:"down", amount: 9},
			{angle:"up", amount: 4},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 8},
			{angle:"up", amount: 2},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 9},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 8},
			{angle:"up", amount: 1},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 2},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 5},
			{angle:"down", amount: 1},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 8},
			{angle:"up", amount: 7},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 6},
			{angle:"up", amount: 1},
			{angle:"up", amount: 6},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 4},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 5},
			{angle:"down", amount: 2},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 8},
			{angle:"down", amount: 2},
			{angle:"down", amount: 5},
			{angle:"down", amount: 6},
			{angle:"up", amount: 8},
			{angle:"down", amount: 1},
			{angle:"up", amount: 7},
			{angle:"up", amount: 4},
			{angle:"up", amount: 1},
			{angle:"up", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 8},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"up", amount: 1},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 9},
			{angle:"down", amount: 1},
			{angle:"up", amount: 7},
			{angle:"up", amount: 3},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 5},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 5},
			{angle:"down", amount: 3},
			{angle:"down", amount: 3},
			{angle:"up", amount: 7},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 9},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 1},
			{angle:"up", amount: 9},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 6},
			{angle:"up", amount: 6},
			{angle:"down", amount: 2},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 4},
			{angle:"down", amount: 4},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 1},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 6},
			{angle:"up", amount: 6},
			{angle:"down", amount: 5},
			{angle:"up", amount: 6},
			{angle:"down", amount: 3},
			{angle:"down", amount: 8},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 6},
			{angle:"down", amount: 9},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 8},
			{angle:"down", amount: 6},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 2},
			{angle:"down", amount: 5},
			{angle:"down", amount: 3},
			{angle:"down", amount: 4},
			{angle:"up", amount: 5},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 6},
			{angle:"up", amount: 8},
			{angle:"down", amount: 3},
			{angle:"down", amount: 5},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 4},
			{angle:"down", amount: 3},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 6},
			{angle:"up", amount: 1},
			{angle:"up", amount: 2},
			{angle:"down", amount: 5},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 3},
			{angle:"down", amount: 6},
			{angle:"up", amount: 3},
			{angle:"down", amount: 9},
			{angle:"down", amount: 2},
			{angle:"up", amount: 4},
			{angle:"down", amount: 3},
			{angle:"up", amount: 7},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 9},
			{angle:"down", amount: 3},
			{angle:"down", amount: 9},
			{angle:"down", amount: 1},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 6},
			{angle:"down", amount: 8},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 1},
			{angle:"up", amount: 4},
			{angle:"forward", amount: 8},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 9},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 4},
			{angle:"up", amount: 2},
			{angle:"down", amount: 6},
			{angle:"down", amount: 5},
			{angle:"down", amount: 8},
			{angle:"down", amount: 2},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 8},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 4},
			{angle:"up", amount: 6},
			{angle:"down", amount: 7},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 5},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 7},
			{angle:"up", amount: 3},
			{angle:"up", amount: 6},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 3},
			{angle:"down", amount: 6},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 4},
			{angle:"down", amount: 3},
			{angle:"down", amount: 1},
			{angle:"down", amount: 4},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 5},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 5},
			{angle:"down", amount: 8},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 9},
			{angle:"up", amount: 1},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 4},
			{angle:"up", amount: 8},
			{angle:"up", amount: 7},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 9},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 3},
			{angle:"down", amount: 3},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 9},
			{angle:"down", amount: 9},
			{angle:"up", amount: 2},
			{angle:"down", amount: 7},
			{angle:"down", amount: 6},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 7},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 4},
			{angle:"up", amount: 6},
			{angle:"up", amount: 1},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 1},
			{angle:"up", amount: 7},
			{angle:"down", amount: 8},
			{angle:"up", amount: 5},
			{angle:"down", amount: 8},
			{angle:"up", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 8},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 5},
			{angle:"down", amount: 7},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 8},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 7},
			{angle:"up", amount: 8},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 5},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 9},
			{angle:"down", amount: 5},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 3},
			{angle:"down", amount: 3},
			{angle:"down", amount: 7},
			{angle:"up", amount: 3},
			{angle:"down", amount: 3},
			{angle:"down", amount: 5},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 6},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 8},
			{angle:"up", amount: 1},
			{angle:"down", amount: 7},
			{angle:"down", amount: 6},
			{angle:"down", amount: 3},
			{angle:"down", amount: 6},
			{angle:"forward", amount: 8},
			{angle:"up", amount: 7},
			{angle:"down", amount: 7},
			{angle:"up", amount: 7},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 9},
			{angle:"up", amount: 8},
			{angle:"down", amount: 2},
			{angle:"down", amount: 3},
			{angle:"down", amount: 7},
			{angle:"down", amount: 2},
			{angle:"up", amount: 2},
			{angle:"down", amount: 1},
			{angle:"down", amount: 7},
			{angle:"up", amount: 6},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 8},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 1},
			{angle:"up", amount: 7},
			{angle:"up", amount: 9},
			{angle:"up", amount: 9},
			{angle:"down", amount: 5},
			{angle:"down", amount: 7},
			{angle:"down", amount: 2},
			{angle:"down", amount: 6},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 8},
			{angle:"forward", amount: 4},
			{angle:"up", amount: 3},
			{angle:"down", amount: 9},
			{angle:"up", amount: 3},
			{angle:"down", amount: 6},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 7},
			{angle:"up", amount: 5},
			{angle:"down", amount: 1},
			{angle:"down", amount: 3},
			{angle:"up", amount: 4},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 7},
			{angle:"down", amount: 3},
			{angle:"down", amount: 7},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 8},
			{angle:"up", amount: 4},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 7},
			{angle:"up", amount: 6},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 4},
			{angle:"down", amount: 4},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 7},
			{angle:"down", amount: 8},
			{angle:"up", amount: 5},
			{angle:"down", amount: 4},
			{angle:"up", amount: 6},
			{angle:"up", amount: 6},
			{angle:"up", amount: 4},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 4},
			{angle:"down", amount: 2},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 5},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 1},
			{angle:"down", amount: 4},
			{angle:"up", amount: 2},
			{angle:"down", amount: 7},
			{angle:"down", amount: 5},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 6},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 9},
			{angle:"up", amount: 4},
			{angle:"down", amount: 1},
			{angle:"down", amount: 3},
			{angle:"up", amount: 7},
			{angle:"up", amount: 5},
			{angle:"down", amount: 9},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 1},
			{angle:"up", amount: 9},
			{angle:"down", amount: 4},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 5},
			{angle:"down", amount: 5},
			{angle:"down", amount: 8},
			{angle:"up", amount: 4},
			{angle:"up", amount: 1},
			{angle:"down", amount: 9},
			{angle:"up", amount: 4},
			{angle:"forward", amount: 9},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 4},
			{angle:"up", amount: 2},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 9},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 8},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 1},
			{angle:"up", amount: 9},
			{angle:"down", amount: 3},
			{angle:"down", amount: 2},
			{angle:"down", amount: 8},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 8},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 3},
			{angle:"up", amount: 6},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 9},
			{angle:"up", amount: 7},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 6},
			{angle:"down", amount: 6},
			{angle:"up", amount: 9},
			{angle:"up", amount: 9},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 8},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 2},
			{angle:"up", amount: 5},
			{angle:"up", amount: 9},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 7},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 1},
			{angle:"down", amount: 6},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 1},
			{angle:"down", amount: 7},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 2},
			{angle:"up", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 2},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 1},
			{angle:"down", amount: 9},
			{angle:"down", amount: 4},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 9},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 8},
			{angle:"up", amount: 8},
			{angle:"down", amount: 8},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 9},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 8},
			{angle:"up", amount: 6},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 6},
			{angle:"down", amount: 2},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 6},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 5},
			{angle:"down", amount: 9},
			{angle:"down", amount: 7},
			{angle:"down", amount: 9},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 2},
			{angle:"down", amount: 7},
			{angle:"down", amount: 7},
			{angle:"up", amount: 2},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"up", amount: 7},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 1},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 3},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 8},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 4},
			{angle:"up", amount: 3},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 9},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 3},
			{angle:"down", amount: 9},
			{angle:"down", amount: 5},
			{angle:"up", amount: 6},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 8},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 6},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 3},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 8},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 4},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 1},
			{angle:"up", amount: 4},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 5},
			{angle:"down", amount: 5},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 3},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 5},
			{angle:"up", amount: 4},
			{angle:"down", amount: 7},
			{angle:"down", amount: 9},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 7},
			{angle:"down", amount: 5},
			{angle:"down", amount: 2},
			{angle:"down", amount: 9},
			{angle:"down", amount: 6},
			{angle:"down", amount: 7},
			{angle:"up", amount: 6},
			{angle:"up", amount: 7},
			{angle:"up", amount: 6},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 8},
			{angle:"down", amount: 7},
			{angle:"down", amount: 8},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 1},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 7},
			{angle:"down", amount: 9},
			{angle:"up", amount: 1},
			{angle:"down", amount: 1},
			{angle:"up", amount: 8},
			{angle:"up", amount: 6},
			{angle:"down", amount: 9},
			{angle:"up", amount: 9},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 9},
			{angle:"down", amount: 7},
			{angle:"up", amount: 1},
			{angle:"down", amount: 2},
			{angle:"up", amount: 2},
			{angle:"down", amount: 3},
			{angle:"down", amount: 1},
			{angle:"up", amount: 4},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 3},
			{angle:"up", amount: 8},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 1},
			{angle:"up", amount: 4},
			{angle:"up", amount: 7},
			{angle:"down", amount: 5},
			{angle:"down", amount: 9},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 7},
			{angle:"down", amount: 2},
			{angle:"up", amount: 9},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 8},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 1},
			{angle:"down", amount: 6},
			{angle:"down", amount: 7},
			{angle:"down", amount: 5},
			{angle:"up", amount: 3},
			{angle:"up", amount: 9},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 3},
			{angle:"down", amount: 2},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 6},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 4},
			{angle:"up", amount: 5},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 3},
			{angle:"down", amount: 4},
			{angle:"down", amount: 3},
			{angle:"down", amount: 1},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 6},
			{angle:"down", amount: 4},
			{angle:"down", amount: 9},
			{angle:"down", amount: 3},
			{angle:"up", amount: 8},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 3},
			{angle:"up", amount: 7},
			{angle:"down", amount: 3},
			{angle:"up", amount: 1},
			{angle:"up", amount: 1},
			{angle:"up", amount: 2},
			{angle:"up", amount: 1},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 8},
			{angle:"down", amount: 8},
			{angle:"up", amount: 5},
			{angle:"down", amount: 4},
			{angle:"down", amount: 4},
			{angle:"down", amount: 6},
			{angle:"down", amount: 9},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 6},
			{angle:"up", amount: 4},
			{angle:"forward", amount: 2},
			{angle:"up", amount: 3},
			{angle:"down", amount: 2},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 1},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 3},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 6},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 6},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 6},
			{angle:"down", amount: 4},
			{angle:"down", amount: 4},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 6},
			{angle:"up", amount: 8},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 3},
			{angle:"down", amount: 8},
			{angle:"down", amount: 9},
			{angle:"up", amount: 3},
			{angle:"down", amount: 7},
			{angle:"forward", amount: 4},
			{angle:"forward", amount: 2},
			{angle:"down", amount: 2},
			{angle:"up", amount: 6},
			{angle:"down", amount: 1},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 1},
			{angle:"down", amount: 7},
			{angle:"down", amount: 7},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 3},
			{angle:"down", amount: 7},
			{angle:"down", amount: 5},
			{angle:"up", amount: 2},
			{angle:"forward", amount: 9},
			{angle:"down", amount: 8},
			{angle:"down", amount: 5},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 2},
			{angle:"forward", amount: 7},
			{angle:"up", amount: 8},
			{angle:"down", amount: 2},
			{angle:"down", amount: 5},
			{angle:"down", amount: 8},
			{angle:"down", amount: 9},
			{angle:"down", amount: 9},
			{angle:"down", amount: 1},
			{angle:"up", amount: 4},
			{angle:"forward", amount: 5},
			{angle:"up", amount: 1},
			{angle:"up", amount: 4},
			{angle:"forward", amount: 1},
			{angle:"down", amount: 1},
			{angle:"down", amount: 7},
			{angle:"up", amount: 9},
			{angle:"up", amount: 7},
			{angle:"down", amount: 5},
			{angle:"down", amount: 9},
			{angle:"down", amount: 9},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 7},
			{angle:"down", amount: 3},
			{angle:"up", amount: 4},
			{angle:"down", amount: 7},
			{angle:"down", amount: 8},
			{angle:"forward", amount: 7},
			{angle:"forward", amount: 4},
			{angle:"up", amount: 9},
			{angle:"down", amount: 2},
			{angle:"up", amount: 7},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 3},
			{angle:"forward", amount: 5},
			{angle:"forward", amount: 5},
			{angle:"down", amount: 2},
			{angle:"down", amount: 2},
			{angle:"down", amount: 7},
			{angle:"up", amount: 8},
			{angle:"up", amount: 9},
			{angle:"down", amount: 1},
			{angle:"forward", amount: 9},
			{angle:"forward", amount: 3},
			{angle:"up", amount: 3},
			{angle:"forward", amount: 9},
			{angle:"up", amount: 2},
			{angle:"down", amount: 7},
			{angle:"down", amount: 3},
			{angle:"forward", amount: 4},
			{angle:"down", amount: 5},
			{angle:"down", amount: 3},
			{angle:"up", amount: 5},
			{angle:"forward", amount: 4},
		},
		1690020,
	},
}

  func TestA(t *testing.T) {
	for _, testData := range data {
	  actual := AMain(testData.input)
	  if !reflect.DeepEqual(actual, testData.expected) {
		fmt.Println(actual, testData.expected)
		t.Errorf("A()")
	  }
	}
  }