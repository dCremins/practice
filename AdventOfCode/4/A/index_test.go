package main

import (
	"testing"
	"reflect"
	"fmt"
)


var data = []struct {
	input1		[]int // input
	input2		[][5][5]Cell // input
	expected	int // expected result
  }{
	{
		[]int{7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1},
		[][5][5]Cell{
			{
				{
					{num: 22, called: false},
					{num: 13, called: false},
					{num: 17, called: false},
					{num: 11, called: false},
					{num: 0, called: false},
				},{
					{num: 8, called: false},
					{num: 2, called: false},
					{num: 23, called: false},
					{num: 4, called: false},
					{num: 24, called: false},
				},{
					{num: 21, called: false},
					{num: 9, called: false},
					{num: 14, called: false},
					{num: 16, called: false},
					{num: 7, called: false},
				},{
					{num: 6, called: false},
					{num: 10, called: false},
					{num: 3, called: false},
					{num: 18, called: false},
					{num: 5, called: false},
				},{
					{num: 1, called: false},
					{num: 12, called: false},
					{num: 20, called: false},
					{num: 15, called: false},
					{num: 19, called: false},
				},
			},

			{
				{
					{num: 3, called: false},
					{num: 15, called: false},
					{num: 0, called: false},
					{num: 2, called: false},
					{num: 22, called: false},
				},{
					{num: 9, called: false},
					{num: 18, called: false},
					{num: 13, called: false},
					{num: 17, called: false},
					{num: 5, called: false},
				},{
					{num: 19, called: false},
					{num: 8, called: false},
					{num: 7, called: false},
					{num: 25, called: false},
					{num: 23, called: false},
				},{
					{num: 20, called: false},
					{num: 11, called: false},
					{num: 10, called: false},
					{num: 24, called: false},
					{num: 4, called: false},
				},{
					{num: 14, called: false},
					{num: 21, called: false},
					{num: 16, called: false},
					{num: 12, called: false},
					{num: 6, called: false},
				},
			},

			{
				{
					{num: 14, called: false},
					{num: 21, called: false},
					{num: 17, called: false},
					{num: 24, called: false},
					{num: 4, called: false},
				},{
					{num: 10, called: false},
					{num: 16, called: false},
					{num: 15, called: false},
					{num: 9, called: false},
					{num: 19, called: false},
				},{
					{num: 18, called: false},
					{num: 8, called: false},
					{num: 23, called: false},
					{num: 26, called: false},
					{num: 20, called: false},
				},{
					{num: 22, called: false},
					{num: 11, called: false},
					{num: 13, called: false},
					{num: 6, called: false},
					{num: 5, called: false},
				},{
					{num: 2, called: false},
					{num: 0, called: false},
					{num: 12, called: false},
					{num: 3, called: false},
					{num: 7, called: false},
				},
			},
		},
		4512,
	},
	{
		[]int{10,80,6,69,22,99,63,92,30,67,28,93,0,50,65,87,38,7,91,60,57,40,84,51,27,12,44,88,64,35,39,74,61,55,31,48,81,89,62,37,94,43,29,14,95,8,78,49,90,97,66,70,25,68,75,45,42,23,9,96,56,72,59,32,85,3,71,79,18,24,33,19,15,20,82,26,21,13,4,98,83,34,86,5,2,73,17,54,1,77,52,58,76,36,16,46,41,47,11,53},
		[][5][5]Cell{
			{{{num: 3, called: false}, {num: 82, called: false}, {num: 18, called: false}, {num: 50, called: false}, {num: 90, called: false},},
			{{num: 16, called: false}, {num: 37, called: false}, {num: 52, called: false}, {num: 67, called: false}, {num: 28, called: false},},
			{{num: 30, called: false}, {num: 54, called: false}, {num: 80, called: false}, {num: 11, called: false}, {num: 10, called: false},},
			{{num: 60, called: false}, {num: 79, called: false}, {num: 7, called: false}, {num: 65, called: false}, {num: 58, called: false},},
			{{num: 76, called: false}, {num: 83, called: false}, {num: 38, called: false}, {num: 51, called: false}, {num: 1, called: false},},},

			{{{num: 83, called: false}, {num: 63, called: false}, {num: 60, called: false}, {num: 88, called: false}, {num: 98, called: false},},
			{{num: 70, called: false}, {num: 87, called: false}, {num: 5, called: false}, {num: 99, called: false}, {num: 14, called: false},},
			{{num: 85, called: false}, {num: 3, called: false}, {num: 11, called: false}, {num: 16, called: false}, {num: 33, called: false},},
			{{num: 72, called: false}, {num: 69, called: false}, {num: 97, called: false}, {num: 36, called: false}, {num: 49, called: false},},
			{{num: 26, called: false}, {num: 17, called: false}, {num: 58, called: false}, {num: 13, called: false}, {num: 2, called: false},},},

			{{{num: 30, called: false}, {num: 80, called: false}, {num: 64, called: false}, {num: 53, called: false}, {num: 69, called: false},},
			{{num: 36, called: false}, {num: 0, called: false}, {num: 32, called: false}, {num: 46, called: false}, {num: 70, called: false},},
			{{num: 13, called: false}, {num: 31, called: false}, {num: 22, called: false}, {num: 95, called: false}, {num: 15, called: false},},
			{{num: 12, called: false}, {num: 35, called: false}, {num: 5, called: false}, {num: 84, called: false}, {num: 21, called: false},},
			{{num: 39, called: false}, {num: 60, called: false}, {num: 68, called: false}, {num: 83, called: false}, {num: 47, called: false},},},

			{{{num: 77, called: false}, {num: 93, called: false}, {num: 26, called: false}, {num: 62, called: false}, {num: 88, called: false},},
			{{num: 87, called: false}, {num: 76, called: false}, {num: 80, called: false}, {num: 10, called: false}, {num: 63, called: false},},
			{{num: 32, called: false}, {num: 7, called: false}, {num: 28, called: false}, {num: 82, called: false}, {num: 44, called: false},},
			{{num: 43, called: false}, {num: 30, called: false}, {num: 31, called: false}, {num: 16, called: false}, {num: 74, called: false},},
			{{num: 33, called: false}, {num: 86, called: false}, {num: 42, called: false}, {num: 45, called: false}, {num: 47, called: false},},},

			{{{num: 95, called: false}, {num: 86, called: false}, {num: 93, called: false}, {num: 45, called: false}, {num: 67, called: false},},
			{{num: 20, called: false}, {num: 58, called: false}, {num: 63, called: false}, {num: 35, called: false}, {num: 97, called: false},},
			{{num: 84, called: false}, {num: 79, called: false}, {num: 10, called: false}, {num: 54, called: false}, {num: 49, called: false},},
			{{num: 48, called: false}, {num: 66, called: false}, {num: 75, called: false}, {num: 23, called: false}, {num: 61, called: false},},
			{{num: 5, called: false}, {num: 30, called: false}, {num: 6, called: false}, {num: 56, called: false}, {num: 71, called: false},},},

			{{{num: 75, called: false}, {num: 8, called: false}, {num: 85, called: false}, {num: 12, called: false}, {num: 98, called: false},},
			{{num: 37, called: false}, {num: 51, called: false}, {num: 91, called: false}, {num: 24, called: false}, {num: 23, called: false},},
			{{num: 50, called: false}, {num: 54, called: false}, {num: 81, called: false}, {num: 53, called: false}, {num: 33, called: false},},
			{{num: 72, called: false}, {num: 57, called: false}, {num: 52, called: false}, {num: 25, called: false}, {num: 6, called: false},},
			{{num: 56, called: false}, {num: 40, called: false}, {num: 95, called: false}, {num: 87, called: false}, {num: 22, called: false},},},

			{{{num: 52, called: false}, {num: 19, called: false}, {num: 53, called: false}, {num: 9, called: false}, {num: 32, called: false},},
			{{num: 23, called: false}, {num: 99, called: false}, {num: 48, called: false}, {num: 26, called: false}, {num: 73, called: false},},
			{{num: 10, called: false}, {num: 8, called: false}, {num: 54, called: false}, {num: 20, called: false}, {num: 79, called: false},},
			{{num: 49, called: false}, {num: 45, called: false}, {num: 34, called: false}, {num: 74, called: false}, {num: 90, called: false},},
			{{num: 27, called: false}, {num: 72, called: false}, {num: 30, called: false}, {num: 13, called: false}, {num: 57, called: false},},},

			{{{num: 1, called: false}, {num: 60, called: false}, {num: 72, called: false}, {num: 85, called: false}, {num: 64, called: false},},
			{{num: 62, called: false}, {num: 39, called: false}, {num: 56, called: false}, {num: 93, called: false}, {num: 78, called: false},},
			{{num: 90, called: false}, {num: 17, called: false}, {num: 87, called: false}, {num: 48, called: false}, {num: 7, called: false},},
			{{num: 9, called: false}, {num: 13, called: false}, {num: 45, called: false}, {num: 23, called: false}, {num: 69, called: false},},
			{{num: 44, called: false}, {num: 80, called: false}, {num: 86, called: false}, {num: 55, called: false}, {num: 21, called: false},},},

			{{{num: 86, called: false}, {num: 57, called: false}, {num: 25, called: false}, {num: 98, called: false}, {num: 18, called: false},},
			{{num: 42, called: false}, {num: 75, called: false}, {num: 38, called: false}, {num: 9, called: false}, {num: 66, called: false},},
			{{num: 54, called: false}, {num: 19, called: false}, {num: 99, called: false}, {num: 87, called: false}, {num: 49, called: false},},
			{{num: 33, called: false}, {num: 32, called: false}, {num: 53, called: false}, {num: 8, called: false}, {num: 6, called: false},},
			{{num: 17, called: false}, {num: 68, called: false}, {num: 24, called: false}, {num: 58, called: false}, {num: 95, called: false},},},

			{{{num: 2, called: false}, {num: 5, called: false}, {num: 80, called: false}, {num: 10, called: false}, {num: 61, called: false},},
			{{num: 27, called: false}, {num: 16, called: false}, {num: 40, called: false}, {num: 67, called: false}, {num: 78, called: false},},
			{{num: 66, called: false}, {num: 13, called: false}, {num: 24, called: false}, {num: 42, called: false}, {num: 75, called: false},},
			{{num: 25, called: false}, {num: 7, called: false}, {num: 35, called: false}, {num: 11, called: false}, {num: 85, called: false},},
			{{num: 93, called: false}, {num: 38, called: false}, {num: 4, called: false}, {num: 31, called: false}, {num: 77, called: false},},},

			{{{num: 52, called: false}, {num: 81, called: false}, {num: 68, called: false}, {num: 88, called: false}, {num: 95, called: false},},
			{{num: 82, called: false}, {num: 50, called: false}, {num: 46, called: false}, {num: 87, called: false}, {num: 20, called: false},},
			{{num: 3, called: false}, {num: 54, called: false}, {num: 59, called: false}, {num: 75, called: false}, {num: 51, called: false},},
			{{num: 92, called: false}, {num: 93, called: false}, {num: 38, called: false}, {num: 72, called: false}, {num: 4, called: false},},
			{{num: 8, called: false}, {num: 77, called: false}, {num: 61, called: false}, {num: 31, called: false}, {num: 56, called: false},},},

			{{{num: 45, called: false}, {num: 16, called: false}, {num: 8, called: false}, {num: 44, called: false}, {num: 62, called: false},},
			{{num: 92, called: false}, {num: 23, called: false}, {num: 42, called: false}, {num: 20, called: false}, {num: 74, called: false},},
			{{num: 73, called: false}, {num: 83, called: false}, {num: 65, called: false}, {num: 9, called: false}, {num: 84, called: false},},
			{{num: 55, called: false}, {num: 13, called: false}, {num: 21, called: false}, {num: 70, called: false}, {num: 59, called: false},},
			{{num: 34, called: false}, {num: 2, called: false}, {num: 98, called: false}, {num: 47, called: false}, {num: 37, called: false},},},

			{{{num: 3, called: false}, {num: 54, called: false}, {num: 85, called: false}, {num: 79, called: false}, {num: 76, called: false},},
			{{num: 42, called: false}, {num: 29, called: false}, {num: 45, called: false}, {num: 12, called: false}, {num: 46, called: false},},
			{{num: 60, called: false}, {num: 59, called: false}, {num: 24, called: false}, {num: 67, called: false}, {num: 80, called: false},},
			{{num: 89, called: false}, {num: 15, called: false}, {num: 99, called: false}, {num: 68, called: false}, {num: 48, called: false},},
			{{num: 40, called: false}, {num: 7, called: false}, {num: 95, called: false}, {num: 44, called: false}, {num: 70, called: false},},},

			{{{num: 75, called: false}, {num: 31, called: false}, {num: 99, called: false}, {num: 16, called: false}, {num: 32, called: false},},
			{{num: 80, called: false}, {num: 56, called: false}, {num: 43, called: false}, {num: 30, called: false}, {num: 36, called: false},},
			{{num: 66, called: false}, {num: 73, called: false}, {num: 35, called: false}, {num: 20, called: false}, {num: 61, called: false},},
			{{num: 67, called: false}, {num: 28, called: false}, {num: 89, called: false}, {num: 23, called: false}, {num: 54, called: false},},
			{{num: 47, called: false}, {num: 26, called: false}, {num: 69, called: false}, {num: 70, called: false}, {num: 50, called: false},},},

			{{{num: 35, called: false}, {num: 91, called: false}, {num: 81, called: false}, {num: 13, called: false}, {num: 15, called: false},},
			{{num: 73, called: false}, {num: 1, called: false}, {num: 37, called: false}, {num: 68, called: false}, {num: 28, called: false},},
			{{num: 98, called: false}, {num: 29, called: false}, {num: 9, called: false}, {num: 22, called: false}, {num: 56, called: false},},
			{{num: 12, called: false}, {num: 59, called: false}, {num: 82, called: false}, {num: 67, called: false}, {num: 31, called: false},},
			{{num: 77, called: false}, {num: 47, called: false}, {num: 32, called: false}, {num: 79, called: false}, {num: 52, called: false},},},

			{{{num: 22, called: false}, {num: 73, called: false}, {num: 39, called: false}, {num: 14, called: false}, {num: 46, called: false},},
			{{num: 99, called: false}, {num: 0, called: false}, {num: 27, called: false}, {num: 34, called: false}, {num: 40, called: false},},
			{{num: 4, called: false}, {num: 5, called: false}, {num: 38, called: false}, {num: 23, called: false}, {num: 18, called: false},},
			{{num: 64, called: false}, {num: 26, called: false}, {num: 89, called: false}, {num: 59, called: false}, {num: 79, called: false},},
			{{num: 71, called: false}, {num: 76, called: false}, {num: 53, called: false}, {num: 49, called: false}, {num: 62, called: false},},},

			{{{num: 14, called: false}, {num: 37, called: false}, {num: 27, called: false}, {num: 67, called: false}, {num: 94, called: false},},
			{{num: 76, called: false}, {num: 16, called: false}, {num: 79, called: false}, {num: 61, called: false}, {num: 83, called: false},},
			{{num: 8, called: false}, {num: 43, called: false}, {num: 36, called: false}, {num: 28, called: false}, {num: 75, called: false},},
			{{num: 10, called: false}, {num: 4, called: false}, {num: 24, called: false}, {num: 56, called: false}, {num: 44, called: false},},
			{{num: 26, called: false}, {num: 1, called: false}, {num: 88, called: false}, {num: 9, called: false}, {num: 86, called: false},},},

			{{{num: 14, called: false}, {num: 78, called: false}, {num: 43, called: false}, {num: 10, called: false}, {num: 30, called: false},},
			{{num: 56, called: false}, {num: 29, called: false}, {num: 1, called: false}, {num: 61, called: false}, {num: 9, called: false},},
			{{num: 7, called: false}, {num: 95, called: false}, {num: 39, called: false}, {num: 35, called: false}, {num: 25, called: false},},
			{{num: 33, called: false}, {num: 87, called: false}, {num: 71, called: false}, {num: 97, called: false}, {num: 21, called: false},},
			{{num: 72, called: false}, {num: 0, called: false}, {num: 4, called: false}, {num: 2, called: false}, {num: 24, called: false},},},

			{{{num: 88, called: false}, {num: 0, called: false}, {num: 72, called: false}, {num: 42, called: false}, {num: 6, called: false},},
			{{num: 53, called: false}, {num: 79, called: false}, {num: 58, called: false}, {num: 80, called: false}, {num: 20, called: false},},
			{{num: 57, called: false}, {num: 84, called: false}, {num: 15, called: false}, {num: 21, called: false}, {num: 64, called: false},},
			{{num: 98, called: false}, {num: 17, called: false}, {num: 43, called: false}, {num: 8, called: false}, {num: 95, called: false},},
			{{num: 2, called: false}, {num: 22, called: false}, {num: 59, called: false}, {num: 63, called: false}, {num: 78, called: false},},},

			{{{num: 78, called: false}, {num: 21, called: false}, {num: 33, called: false}, {num: 57, called: false}, {num: 72, called: false},},
			{{num: 10, called: false}, {num: 69, called: false}, {num: 85, called: false}, {num: 73, called: false}, {num: 16, called: false},},
			{{num: 92, called: false}, {num: 60, called: false}, {num: 87, called: false}, {num: 39, called: false}, {num: 63, called: false},},
			{{num: 40, called: false}, {num: 15, called: false}, {num: 77, called: false}, {num: 80, called: false}, {num: 56, called: false},},
			{{num: 6, called: false}, {num: 62, called: false}, {num: 99, called: false}, {num: 50, called: false}, {num: 3, called: false},},},

			{{{num: 38, called: false}, {num: 72, called: false}, {num: 34, called: false}, {num: 41, called: false}, {num: 74, called: false},},
			{{num: 90, called: false}, {num: 29, called: false}, {num: 9, called: false}, {num: 6, called: false}, {num: 91, called: false},},
			{{num: 94, called: false}, {num: 39, called: false}, {num: 56, called: false}, {num: 71, called: false}, {num: 67, called: false},},
			{{num: 53, called: false}, {num: 21, called: false}, {num: 22, called: false}, {num: 32, called: false}, {num: 10, called: false},},
			{{num: 73, called: false}, {num: 48, called: false}, {num: 79, called: false}, {num: 47, called: false}, {num: 85, called: false},},},

			{{{num: 5, called: false}, {num: 49, called: false}, {num: 73, called: false}, {num: 24, called: false}, {num: 8, called: false},},
			{{num: 75, called: false}, {num: 12, called: false}, {num: 11, called: false}, {num: 47, called: false}, {num: 69, called: false},},
			{{num: 66, called: false}, {num: 70, called: false}, {num: 89, called: false}, {num: 62, called: false}, {num: 48, called: false},},
			{{num: 99, called: false}, {num: 3, called: false}, {num: 29, called: false}, {num: 88, called: false}, {num: 30, called: false},},
			{{num: 10, called: false}, {num: 40, called: false}, {num: 32, called: false}, {num: 33, called: false}, {num: 43, called: false},},},

			{{{num: 61, called: false}, {num: 93, called: false}, {num: 2, called: false}, {num: 58, called: false}, {num: 84, called: false},},
			{{num: 47, called: false}, {num: 62, called: false}, {num: 51, called: false}, {num: 16, called: false}, {num: 82, called: false},},
			{{num: 80, called: false}, {num: 22, called: false}, {num: 50, called: false}, {num: 31, called: false}, {num: 65, called: false},},
			{{num: 76, called: false}, {num: 85, called: false}, {num: 83, called: false}, {num: 4, called: false}, {num: 40, called: false},},
			{{num: 86, called: false}, {num: 59, called: false}, {num: 68, called: false}, {num: 14, called: false}, {num: 69, called: false},},},

			{{{num: 52, called: false}, {num: 5, called: false}, {num: 74, called: false}, {num: 9, called: false}, {num: 72, called: false},},
			{{num: 84, called: false}, {num: 69, called: false}, {num: 38, called: false}, {num: 1, called: false}, {num: 27, called: false},},
			{{num: 78, called: false}, {num: 90, called: false}, {num: 46, called: false}, {num: 97, called: false}, {num: 95, called: false},},
			{{num: 57, called: false}, {num: 21, called: false}, {num: 32, called: false}, {num: 93, called: false}, {num: 29, called: false},},
			{{num: 11, called: false}, {num: 66, called: false}, {num: 20, called: false}, {num: 51, called: false}, {num: 48, called: false},},},

			{{{num: 2, called: false}, {num: 3, called: false}, {num: 58, called: false}, {num: 18, called: false}, {num: 53, called: false},},
			{{num: 11, called: false}, {num: 96, called: false}, {num: 63, called: false}, {num: 33, called: false}, {num: 13, called: false},},
			{{num: 55, called: false}, {num: 47, called: false}, {num: 30, called: false}, {num: 9, called: false}, {num: 46, called: false},},
			{{num: 98, called: false}, {num: 85, called: false}, {num: 79, called: false}, {num: 19, called: false}, {num: 65, called: false},},
			{{num: 87, called: false}, {num: 94, called: false}, {num: 77, called: false}, {num: 27, called: false}, {num: 75, called: false},},},

			{{{num: 54, called: false}, {num: 97, called: false}, {num: 46, called: false}, {num: 33, called: false}, {num: 90, called: false},},
			{{num: 99, called: false}, {num: 93, called: false}, {num: 22, called: false}, {num: 0, called: false}, {num: 51, called: false},},
			{{num: 83, called: false}, {num: 53, called: false}, {num: 34, called: false}, {num: 29, called: false}, {num: 38, called: false},},
			{{num: 35, called: false}, {num: 65, called: false}, {num: 80, called: false}, {num: 82, called: false}, {num: 9, called: false},},
			{{num: 56, called: false}, {num: 30, called: false}, {num: 19, called: false}, {num: 49, called: false}, {num: 15, called: false},},},

			{{{num: 43, called: false}, {num: 40, called: false}, {num: 51, called: false}, {num: 67, called: false}, {num: 37, called: false},},
			{{num: 4, called: false}, {num: 30, called: false}, {num: 85, called: false}, {num: 24, called: false}, {num: 21, called: false},},
			{{num: 83, called: false}, {num: 94, called: false}, {num: 69, called: false}, {num: 91, called: false}, {num: 99, called: false},},
			{{num: 13, called: false}, {num: 32, called: false}, {num: 82, called: false}, {num: 86, called: false}, {num: 12, called: false},},
			{{num: 66, called: false}, {num: 9, called: false}, {num: 60, called: false}, {num: 65, called: false}, {num: 97, called: false},},},

			{{{num: 71, called: false}, {num: 96, called: false}, {num: 55, called: false}, {num: 92, called: false}, {num: 6, called: false},},
			{{num: 83, called: false}, {num: 8, called: false}, {num: 63, called: false}, {num: 56, called: false}, {num: 18, called: false},},
			{{num: 4, called: false}, {num: 0, called: false}, {num: 74, called: false}, {num: 70, called: false}, {num: 34, called: false},},
			{{num: 15, called: false}, {num: 87, called: false}, {num: 44, called: false}, {num: 80, called: false}, {num: 29, called: false},},
			{{num: 68, called: false}, {num: 33, called: false}, {num: 99, called: false}, {num: 14, called: false}, {num: 47, called: false},},},

			{{{num: 76, called: false}, {num: 86, called: false}, {num: 46, called: false}, {num: 6, called: false}, {num: 8, called: false},},
			{{num: 90, called: false}, {num: 80, called: false}, {num: 77, called: false}, {num: 30, called: false}, {num: 62, called: false},},
			{{num: 97, called: false}, {num: 66, called: false}, {num: 55, called: false}, {num: 59, called: false}, {num: 36, called: false},},
			{{num: 1, called: false}, {num: 43, called: false}, {num: 27, called: false}, {num: 15, called: false}, {num: 57, called: false},},
			{{num: 54, called: false}, {num: 70, called: false}, {num: 38, called: false}, {num: 21, called: false}, {num: 89, called: false},},},

			{{{num: 75, called: false}, {num: 96, called: false}, {num: 97, called: false}, {num: 54, called: false}, {num: 29, called: false},},
			{{num: 62, called: false}, {num: 43, called: false}, {num: 69, called: false}, {num: 57, called: false}, {num: 88, called: false},},
			{{num: 36, called: false}, {num: 46, called: false}, {num: 73, called: false}, {num: 84, called: false}, {num: 28, called: false},},
			{{num: 18, called: false}, {num: 98, called: false}, {num: 38, called: false}, {num: 63, called: false}, {num: 4, called: false},},
			{{num: 59, called: false}, {num: 35, called: false}, {num: 99, called: false}, {num: 90, called: false}, {num: 58, called: false},},},

			{{{num: 66, called: false}, {num: 29, called: false}, {num: 60, called: false}, {num: 25, called: false}, {num: 95, called: false},},
			{{num: 91, called: false}, {num: 4, called: false}, {num: 87, called: false}, {num: 76, called: false}, {num: 41, called: false},},
			{{num: 26, called: false}, {num: 19, called: false}, {num: 45, called: false}, {num: 96, called: false}, {num: 74, called: false},},
			{{num: 7, called: false}, {num: 82, called: false}, {num: 3, called: false}, {num: 81, called: false}, {num: 31, called: false},},
			{{num: 17, called: false}, {num: 64, called: false}, {num: 51, called: false}, {num: 93, called: false}, {num: 71, called: false},},},

			{{{num: 41, called: false}, {num: 51, called: false}, {num: 14, called: false}, {num: 70, called: false}, {num: 26, called: false},},
			{{num: 35, called: false}, {num: 38, called: false}, {num: 50, called: false}, {num: 25, called: false}, {num: 13, called: false},},
			{{num: 95, called: false}, {num: 60, called: false}, {num: 88, called: false}, {num: 36, called: false}, {num: 24, called: false},},
			{{num: 66, called: false}, {num: 94, called: false}, {num: 62, called: false}, {num: 97, called: false}, {num: 83, called: false},},
			{{num: 21, called: false}, {num: 10, called: false}, {num: 37, called: false}, {num: 1, called: false}, {num: 96, called: false},},},

			{{{num: 38, called: false}, {num: 90, called: false}, {num: 81, called: false}, {num: 96, called: false}, {num: 14, called: false},},
			{{num: 40, called: false}, {num: 82, called: false}, {num: 71, called: false}, {num: 18, called: false}, {num: 83, called: false},},
			{{num: 78, called: false}, {num: 17, called: false}, {num: 65, called: false}, {num: 46, called: false}, {num: 84, called: false},},
			{{num: 7, called: false}, {num: 92, called: false}, {num: 63, called: false}, {num: 79, called: false}, {num: 49, called: false},},
			{{num: 55, called: false}, {num: 21, called: false}, {num: 89, called: false}, {num: 95, called: false}, {num: 72, called: false},},},

			{{{num: 60, called: false}, {num: 35, called: false}, {num: 26, called: false}, {num: 67, called: false}, {num: 42, called: false},},
			{{num: 59, called: false}, {num: 77, called: false}, {num: 15, called: false}, {num: 6, called: false}, {num: 75, called: false},},
			{{num: 18, called: false}, {num: 16, called: false}, {num: 21, called: false}, {num: 55, called: false}, {num: 4, called: false},},
			{{num: 98, called: false}, {num: 49, called: false}, {num: 38, called: false}, {num: 43, called: false}, {num: 30, called: false},},
			{{num: 0, called: false}, {num: 85, called: false}, {num: 69, called: false}, {num: 96, called: false}, {num: 19, called: false},},},

			{{{num: 93, called: false}, {num: 22, called: false}, {num: 87, called: false}, {num: 66, called: false}, {num: 94, called: false},},
			{{num: 7, called: false}, {num: 43, called: false}, {num: 98, called: false}, {num: 18, called: false}, {num: 57, called: false},},
			{{num: 29, called: false}, {num: 20, called: false}, {num: 91, called: false}, {num: 60, called: false}, {num: 21, called: false},},
			{{num: 28, called: false}, {num: 51, called: false}, {num: 17, called: false}, {num: 10, called: false}, {num: 14, called: false},},
			{{num: 96, called: false}, {num: 12, called: false}, {num: 15, called: false}, {num: 25, called: false}, {num: 37, called: false},},},

			{{{num: 12, called: false}, {num: 69, called: false}, {num: 27, called: false}, {num: 41, called: false}, {num: 36, called: false},},
			{{num: 58, called: false}, {num: 11, called: false}, {num: 42, called: false}, {num: 44, called: false}, {num: 9, called: false},},
			{{num: 8, called: false}, {num: 56, called: false}, {num: 33, called: false}, {num: 7, called: false}, {num: 30, called: false},},
			{{num: 70, called: false}, {num: 64, called: false}, {num: 78, called: false}, {num: 17, called: false}, {num: 61, called: false},},
			{{num: 22, called: false}, {num: 28, called: false}, {num: 94, called: false}, {num: 0, called: false}, {num: 99, called: false},},},

			{{{num: 16, called: false}, {num: 64, called: false}, {num: 88, called: false}, {num: 22, called: false}, {num: 48, called: false},},
			{{num: 65, called: false}, {num: 42, called: false}, {num: 23, called: false}, {num: 7, called: false}, {num: 26, called: false},},
			{{num: 97, called: false}, {num: 12, called: false}, {num: 63, called: false}, {num: 57, called: false}, {num: 45, called: false},},
			{{num: 29, called: false}, {num: 94, called: false}, {num: 91, called: false}, {num: 21, called: false}, {num: 54, called: false},},
			{{num: 95, called: false}, {num: 43, called: false}, {num: 0, called: false}, {num: 85, called: false}, {num: 46, called: false},},},

			{{{num: 9, called: false}, {num: 50, called: false}, {num: 54, called: false}, {num: 98, called: false}, {num: 35, called: false},},
			{{num: 45, called: false}, {num: 37, called: false}, {num: 84, called: false}, {num: 87, called: false}, {num: 5, called: false},},
			{{num: 40, called: false}, {num: 23, called: false}, {num: 14, called: false}, {num: 17, called: false}, {num: 18, called: false},},
			{{num: 73, called: false}, {num: 43, called: false}, {num: 86, called: false}, {num: 41, called: false}, {num: 59, called: false},},
			{{num: 69, called: false}, {num: 77, called: false}, {num: 78, called: false}, {num: 15, called: false}, {num: 60, called: false},},},

			{{{num: 79, called: false}, {num: 9, called: false}, {num: 45, called: false}, {num: 59, called: false}, {num: 85, called: false},},
			{{num: 38, called: false}, {num: 56, called: false}, {num: 64, called: false}, {num: 95, called: false}, {num: 60, called: false},},
			{{num: 39, called: false}, {num: 22, called: false}, {num: 14, called: false}, {num: 57, called: false}, {num: 66, called: false},},
			{{num: 98, called: false}, {num: 53, called: false}, {num: 83, called: false}, {num: 76, called: false}, {num: 16, called: false},},
			{{num: 62, called: false}, {num: 94, called: false}, {num: 72, called: false}, {num: 54, called: false}, {num: 82, called: false},},},

			{{{num: 77, called: false}, {num: 44, called: false}, {num: 6, called: false}, {num: 66, called: false}, {num: 46, called: false},},
			{{num: 9, called: false}, {num: 89, called: false}, {num: 11, called: false}, {num: 84, called: false}, {num: 63, called: false},},
			{{num: 81, called: false}, {num: 94, called: false}, {num: 87, called: false}, {num: 83, called: false}, {num: 21, called: false},},
			{{num: 22, called: false}, {num: 90, called: false}, {num: 1, called: false}, {num: 93, called: false}, {num: 92, called: false},},
			{{num: 24, called: false}, {num: 65, called: false}, {num: 34, called: false}, {num: 45, called: false}, {num: 99, called: false},},},

			{{{num: 36, called: false}, {num: 93, called: false}, {num: 59, called: false}, {num: 5, called: false}, {num: 43, called: false},},
			{{num: 76, called: false}, {num: 49, called: false}, {num: 51, called: false}, {num: 0, called: false}, {num: 68, called: false},},
			{{num: 71, called: false}, {num: 34, called: false}, {num: 55, called: false}, {num: 7, called: false}, {num: 73, called: false},},
			{{num: 14, called: false}, {num: 10, called: false}, {num: 45, called: false}, {num: 63, called: false}, {num: 95, called: false},},
			{{num: 30, called: false}, {num: 94, called: false}, {num: 79, called: false}, {num: 67, called: false}, {num: 11, called: false},},},

			{{{num: 93, called: false}, {num: 98, called: false}, {num: 82, called: false}, {num: 96, called: false}, {num: 91, called: false},},
			{{num: 3, called: false}, {num: 79, called: false}, {num: 55, called: false}, {num: 70, called: false}, {num: 24, called: false},},
			{{num: 68, called: false}, {num: 56, called: false}, {num: 87, called: false}, {num: 12, called: false}, {num: 76, called: false},},
			{{num: 19, called: false}, {num: 31, called: false}, {num: 67, called: false}, {num: 1, called: false}, {num: 54, called: false},},
			{{num: 49, called: false}, {num: 62, called: false}, {num: 23, called: false}, {num: 15, called: false}, {num: 10, called: false},},},

			{{{num: 10, called: false}, {num: 74, called: false}, {num: 98, called: false}, {num: 15, called: false}, {num: 6, called: false},},
			{{num: 14, called: false}, {num: 31, called: false}, {num: 66, called: false}, {num: 38, called: false}, {num: 86, called: false},},
			{{num: 68, called: false}, {num: 84, called: false}, {num: 60, called: false}, {num: 80, called: false}, {num: 26, called: false},},
			{{num: 34, called: false}, {num: 72, called: false}, {num: 87, called: false}, {num: 92, called: false}, {num: 61, called: false},},
			{{num: 81, called: false}, {num: 56, called: false}, {num: 73, called: false}, {num: 12, called: false}, {num: 53, called: false},},},

			{{{num: 11, called: false}, {num: 69, called: false}, {num: 4, called: false}, {num: 6, called: false}, {num: 23, called: false},},
			{{num: 38, called: false}, {num: 47, called: false}, {num: 16, called: false}, {num: 99, called: false}, {num: 96, called: false},},
			{{num: 7, called: false}, {num: 13, called: false}, {num: 40, called: false}, {num: 41, called: false}, {num: 78, called: false},},
			{{num: 12, called: false}, {num: 5, called: false}, {num: 1, called: false}, {num: 18, called: false}, {num: 88, called: false},},
			{{num: 20, called: false}, {num: 42, called: false}, {num: 10, called: false}, {num: 82, called: false}, {num: 73, called: false},},},

			{{{num: 66, called: false}, {num: 97, called: false}, {num: 72, called: false}, {num: 55, called: false}, {num: 99, called: false},},
			{{num: 26, called: false}, {num: 59, called: false}, {num: 6, called: false}, {num: 79, called: false}, {num: 53, called: false},},
			{{num: 74, called: false}, {num: 80, called: false}, {num: 98, called: false}, {num: 28, called: false}, {num: 69, called: false},},
			{{num: 25, called: false}, {num: 95, called: false}, {num: 17, called: false}, {num: 29, called: false}, {num: 34, called: false},},
			{{num: 85, called: false}, {num: 64, called: false}, {num: 84, called: false}, {num: 90, called: false}, {num: 42, called: false},},},

			{{{num: 95, called: false}, {num: 50, called: false}, {num: 58, called: false}, {num: 51, called: false}, {num: 66, called: false},},
			{{num: 48, called: false}, {num: 27, called: false}, {num: 81, called: false}, {num: 94, called: false}, {num: 0, called: false},},
			{{num: 35, called: false}, {num: 82, called: false}, {num: 57, called: false}, {num: 71, called: false}, {num: 16, called: false},},
			{{num: 32, called: false}, {num: 93, called: false}, {num: 70, called: false}, {num: 40, called: false}, {num: 25, called: false},},
			{{num: 31, called: false}, {num: 73, called: false}, {num: 46, called: false}, {num: 12, called: false}, {num: 90, called: false},},},

			{{{num: 39, called: false}, {num: 94, called: false}, {num: 52, called: false}, {num: 9, called: false}, {num: 88, called: false},},
			{{num: 3, called: false}, {num: 23, called: false}, {num: 59, called: false}, {num: 77, called: false}, {num: 29, called: false},},
			{{num: 2, called: false}, {num: 40, called: false}, {num: 93, called: false}, {num: 85, called: false}, {num: 38, called: false},},
			{{num: 74, called: false}, {num: 97, called: false}, {num: 12, called: false}, {num: 50, called: false}, {num: 1, called: false},},
			{{num: 22, called: false}, {num: 36, called: false}, {num: 68, called: false}, {num: 65, called: false}, {num: 37, called: false},},},

			{{{num: 70, called: false}, {num: 15, called: false}, {num: 44, called: false}, {num: 90, called: false}, {num: 55, called: false},},
			{{num: 42, called: false}, {num: 20, called: false}, {num: 82, called: false}, {num: 0, called: false}, {num: 78, called: false},},
			{{num: 10, called: false}, {num: 49, called: false}, {num: 62, called: false}, {num: 3, called: false}, {num: 22, called: false},},
			{{num: 91, called: false}, {num: 73, called: false}, {num: 84, called: false}, {num: 40, called: false}, {num: 28, called: false},},
			{{num: 72, called: false}, {num: 13, called: false}, {num: 11, called: false}, {num: 60, called: false}, {num: 19, called: false},},},

			{{{num: 58, called: false}, {num: 95, called: false}, {num: 66, called: false}, {num: 36, called: false}, {num: 22, called: false},},
			{{num: 91, called: false}, {num: 99, called: false}, {num: 77, called: false}, {num: 94, called: false}, {num: 44, called: false},},
			{{num: 70, called: false}, {num: 14, called: false}, {num: 85, called: false}, {num: 13, called: false}, {num: 52, called: false},},
			{{num: 49, called: false}, {num: 6, called: false}, {num: 2, called: false}, {num: 50, called: false}, {num: 35, called: false},},
			{{num: 47, called: false}, {num: 42, called: false}, {num: 15, called: false}, {num: 98, called: false}, {num: 63, called: false},},},

			{{{num: 35, called: false}, {num: 1, called: false}, {num: 99, called: false}, {num: 21, called: false}, {num: 68, called: false},},
			{{num: 93, called: false}, {num: 32, called: false}, {num: 30, called: false}, {num: 76, called: false}, {num: 5, called: false},},
			{{num: 79, called: false}, {num: 96, called: false}, {num: 10, called: false}, {num: 85, called: false}, {num: 16, called: false},},
			{{num: 19, called: false}, {num: 69, called: false}, {num: 81, called: false}, {num: 78, called: false}, {num: 70, called: false},},
			{{num: 66, called: false}, {num: 36, called: false}, {num: 26, called: false}, {num: 94, called: false}, {num: 39, called: false},},},

			{{{num: 78, called: false}, {num: 51, called: false}, {num: 55, called: false}, {num: 4, called: false}, {num: 97, called: false},},
			{{num: 21, called: false}, {num: 36, called: false}, {num: 53, called: false}, {num: 1, called: false}, {num: 26, called: false},},
			{{num: 77, called: false}, {num: 42, called: false}, {num: 20, called: false}, {num: 12, called: false}, {num: 65, called: false},},
			{{num: 17, called: false}, {num: 52, called: false}, {num: 6, called: false}, {num: 40, called: false}, {num: 16, called: false},},
			{{num: 19, called: false}, {num: 85, called: false}, {num: 2, called: false}, {num: 24, called: false}, {num: 23, called: false},},},

			{{{num: 95, called: false}, {num: 68, called: false}, {num: 76, called: false}, {num: 14, called: false}, {num: 30, called: false},},
			{{num: 19, called: false}, {num: 11, called: false}, {num: 64, called: false}, {num: 99, called: false}, {num: 60, called: false},},
			{{num: 63, called: false}, {num: 55, called: false}, {num: 8, called: false}, {num: 40, called: false}, {num: 65, called: false},},
			{{num: 41, called: false}, {num: 75, called: false}, {num: 62, called: false}, {num: 53, called: false}, {num: 83, called: false},},
			{{num: 26, called: false}, {num: 34, called: false}, {num: 46, called: false}, {num: 72, called: false}, {num: 79, called: false},},},

			{{{num: 68, called: false}, {num: 6, called: false}, {num: 35, called: false}, {num: 62, called: false}, {num: 77, called: false},},
			{{num: 43, called: false}, {num: 14, called: false}, {num: 88, called: false}, {num: 7, called: false}, {num: 11, called: false},},
			{{num: 40, called: false}, {num: 45, called: false}, {num: 98, called: false}, {num: 86, called: false}, {num: 64, called: false},},
			{{num: 3, called: false}, {num: 53, called: false}, {num: 56, called: false}, {num: 87, called: false}, {num: 30, called: false},},
			{{num: 28, called: false}, {num: 37, called: false}, {num: 48, called: false}, {num: 10, called: false}, {num: 72, called: false},},},

			{{{num: 13, called: false}, {num: 69, called: false}, {num: 72, called: false}, {num: 93, called: false}, {num: 17, called: false},},
			{{num: 1, called: false}, {num: 46, called: false}, {num: 8, called: false}, {num: 56, called: false}, {num: 37, called: false},},
			{{num: 78, called: false}, {num: 27, called: false}, {num: 49, called: false}, {num: 64, called: false}, {num: 59, called: false},},
			{{num: 81, called: false}, {num: 99, called: false}, {num: 33, called: false}, {num: 76, called: false}, {num: 79, called: false},},
			{{num: 84, called: false}, {num: 98, called: false}, {num: 51, called: false}, {num: 82, called: false}, {num: 31, called: false},},},

			{{{num: 57, called: false}, {num: 41, called: false}, {num: 45, called: false}, {num: 15, called: false}, {num: 10, called: false},},
			{{num: 65, called: false}, {num: 72, called: false}, {num: 79, called: false}, {num: 17, called: false}, {num: 29, called: false},},
			{{num: 67, called: false}, {num: 0, called: false}, {num: 33, called: false}, {num: 32, called: false}, {num: 69, called: false},},
			{{num: 56, called: false}, {num: 96, called: false}, {num: 92, called: false}, {num: 46, called: false}, {num: 53, called: false},},
			{{num: 88, called: false}, {num: 3, called: false}, {num: 18, called: false}, {num: 87, called: false}, {num: 51, called: false},},},

			{{{num: 97, called: false}, {num: 52, called: false}, {num: 58, called: false}, {num: 67, called: false}, {num: 17, called: false},},
			{{num: 51, called: false}, {num: 69, called: false}, {num: 43, called: false}, {num: 20, called: false}, {num: 63, called: false},},
			{{num: 1, called: false}, {num: 26, called: false}, {num: 27, called: false}, {num: 47, called: false}, {num: 99, called: false},},
			{{num: 53, called: false}, {num: 23, called: false}, {num: 14, called: false}, {num: 90, called: false}, {num: 86, called: false},},
			{{num: 4, called: false}, {num: 56, called: false}, {num: 13, called: false}, {num: 36, called: false}, {num: 11, called: false},},},

			{{{num: 88, called: false}, {num: 11, called: false}, {num: 57, called: false}, {num: 73, called: false}, {num: 89, called: false},},
			{{num: 43, called: false}, {num: 34, called: false}, {num: 91, called: false}, {num: 15, called: false}, {num: 58, called: false},},
			{{num: 9, called: false}, {num: 39, called: false}, {num: 18, called: false}, {num: 12, called: false}, {num: 14, called: false},},
			{{num: 1, called: false}, {num: 98, called: false}, {num: 29, called: false}, {num: 77, called: false}, {num: 52, called: false},},
			{{num: 84, called: false}, {num: 97, called: false}, {num: 96, called: false}, {num: 68, called: false}, {num: 10, called: false},},},

			{{{num: 99, called: false}, {num: 5, called: false}, {num: 69, called: false}, {num: 53, called: false}, {num: 45, called: false},},
			{{num: 10, called: false}, {num: 43, called: false}, {num: 24, called: false}, {num: 60, called: false}, {num: 55, called: false},},
			{{num: 64, called: false}, {num: 57, called: false}, {num: 30, called: false}, {num: 3, called: false}, {num: 0, called: false},},
			{{num: 22, called: false}, {num: 65, called: false}, {num: 68, called: false}, {num: 32, called: false}, {num: 83, called: false},},
			{{num: 52, called: false}, {num: 38, called: false}, {num: 74, called: false}, {num: 97, called: false}, {num: 20, called: false},},},

			{{{num: 27, called: false}, {num: 25, called: false}, {num: 33, called: false}, {num: 41, called: false}, {num: 67, called: false},},
			{{num: 54, called: false}, {num: 42, called: false}, {num: 3, called: false}, {num: 1, called: false}, {num: 55, called: false},},
			{{num: 66, called: false}, {num: 92, called: false}, {num: 44, called: false}, {num: 98, called: false}, {num: 35, called: false},},
			{{num: 14, called: false}, {num: 82, called: false}, {num: 5, called: false}, {num: 10, called: false}, {num: 39, called: false},},
			{{num: 52, called: false}, {num: 79, called: false}, {num: 69, called: false}, {num: 76, called: false}, {num: 48, called: false},},},

			{{{num: 64, called: false}, {num: 58, called: false}, {num: 60, called: false}, {num: 91, called: false}, {num: 42, called: false},},
			{{num: 45, called: false}, {num: 55, called: false}, {num: 35, called: false}, {num: 9, called: false}, {num: 72, called: false},},
			{{num: 36, called: false}, {num: 74, called: false}, {num: 99, called: false}, {num: 33, called: false}, {num: 26, called: false},},
			{{num: 67, called: false}, {num: 4, called: false}, {num: 25, called: false}, {num: 50, called: false}, {num: 14, called: false},},
			{{num: 15, called: false}, {num: 2, called: false}, {num: 96, called: false}, {num: 82, called: false}, {num: 11, called: false},},},

			{{{num: 34, called: false}, {num: 84, called: false}, {num: 90, called: false}, {num: 95, called: false}, {num: 26, called: false},},
			{{num: 8, called: false}, {num: 66, called: false}, {num: 52, called: false}, {num: 43, called: false}, {num: 63, called: false},},
			{{num: 79, called: false}, {num: 98, called: false}, {num: 36, called: false}, {num: 85, called: false}, {num: 41, called: false},},
			{{num: 47, called: false}, {num: 24, called: false}, {num: 33, called: false}, {num: 88, called: false}, {num: 71, called: false},},
			{{num: 86, called: false}, {num: 91, called: false}, {num: 83, called: false}, {num: 40, called: false}, {num: 18, called: false},},},

			{{{num: 79, called: false}, {num: 68, called: false}, {num: 49, called: false}, {num: 64, called: false}, {num: 35, called: false},},
			{{num: 23, called: false}, {num: 57, called: false}, {num: 27, called: false}, {num: 77, called: false}, {num: 71, called: false},},
			{{num: 95, called: false}, {num: 39, called: false}, {num: 43, called: false}, {num: 19, called: false}, {num: 98, called: false},},
			{{num: 78, called: false}, {num: 62, called: false}, {num: 65, called: false}, {num: 58, called: false}, {num: 60, called: false},},
			{{num: 52, called: false}, {num: 73, called: false}, {num: 82, called: false}, {num: 4, called: false}, {num: 32, called: false},},},

			{{{num: 22, called: false}, {num: 54, called: false}, {num: 57, called: false}, {num: 45, called: false}, {num: 3, called: false},},
			{{num: 43, called: false}, {num: 85, called: false}, {num: 30, called: false}, {num: 60, called: false}, {num: 94, called: false},},
			{{num: 35, called: false}, {num: 46, called: false}, {num: 28, called: false}, {num: 55, called: false}, {num: 6, called: false},},
			{{num: 82, called: false}, {num: 42, called: false}, {num: 13, called: false}, {num: 83, called: false}, {num: 59, called: false},},
			{{num: 76, called: false}, {num: 70, called: false}, {num: 41, called: false}, {num: 61, called: false}, {num: 1, called: false},},},

			{{{num: 76, called: false}, {num: 89, called: false}, {num: 34, called: false}, {num: 96, called: false}, {num: 1, called: false},},
			{{num: 95, called: false}, {num: 60, called: false}, {num: 55, called: false}, {num: 23, called: false}, {num: 88, called: false},},
			{{num: 37, called: false}, {num: 13, called: false}, {num: 61, called: false}, {num: 92, called: false}, {num: 62, called: false},},
			{{num: 98, called: false}, {num: 77, called: false}, {num: 32, called: false}, {num: 82, called: false}, {num: 31, called: false},},
			{{num: 33, called: false}, {num: 74, called: false}, {num: 71, called: false}, {num: 58, called: false}, {num: 86, called: false},},},

			{{{num: 73, called: false}, {num: 91, called: false}, {num: 92, called: false}, {num: 49, called: false}, {num: 44, called: false},},
			{{num: 53, called: false}, {num: 6, called: false}, {num: 29, called: false}, {num: 8, called: false}, {num: 95, called: false},},
			{{num: 4, called: false}, {num: 31, called: false}, {num: 54, called: false}, {num: 20, called: false}, {num: 97, called: false},},
			{{num: 98, called: false}, {num: 57, called: false}, {num: 2, called: false}, {num: 65, called: false}, {num: 75, called: false},},
			{{num: 43, called: false}, {num: 88, called: false}, {num: 1, called: false}, {num: 58, called: false}, {num: 0, called: false},},},

			{{{num: 49, called: false}, {num: 91, called: false}, {num: 70, called: false}, {num: 1, called: false}, {num: 79, called: false},},
			{{num: 17, called: false}, {num: 90, called: false}, {num: 33, called: false}, {num: 65, called: false}, {num: 54, called: false},},
			{{num: 56, called: false}, {num: 47, called: false}, {num: 63, called: false}, {num: 83, called: false}, {num: 52, called: false},},
			{{num: 8, called: false}, {num: 45, called: false}, {num: 72, called: false}, {num: 84, called: false}, {num: 39, called: false},},
			{{num: 43, called: false}, {num: 37, called: false}, {num: 71, called: false}, {num: 97, called: false}, {num: 59, called: false},},},

			{{{num: 90, called: false}, {num: 93, called: false}, {num: 20, called: false}, {num: 31, called: false}, {num: 96, called: false},},
			{{num: 98, called: false}, {num: 84, called: false}, {num: 2, called: false}, {num: 87, called: false}, {num: 73, called: false},},
			{{num: 97, called: false}, {num: 16, called: false}, {num: 19, called: false}, {num: 24, called: false}, {num: 38, called: false},},
			{{num: 14, called: false}, {num: 11, called: false}, {num: 94, called: false}, {num: 92, called: false}, {num: 36, called: false},},
			{{num: 4, called: false}, {num: 10, called: false}, {num: 27, called: false}, {num: 44, called: false}, {num: 30, called: false},},},

			{{{num: 20, called: false}, {num: 77, called: false}, {num: 81, called: false}, {num: 80, called: false}, {num: 28, called: false},},
			{{num: 35, called: false}, {num: 51, called: false}, {num: 93, called: false}, {num: 24, called: false}, {num: 62, called: false},},
			{{num: 54, called: false}, {num: 56, called: false}, {num: 41, called: false}, {num: 68, called: false}, {num: 79, called: false},},
			{{num: 29, called: false}, {num: 67, called: false}, {num: 89, called: false}, {num: 60, called: false}, {num: 12, called: false},},
			{{num: 63, called: false}, {num: 91, called: false}, {num: 18, called: false}, {num: 90, called: false}, {num: 99, called: false},},},

			{{{num: 28, called: false}, {num: 48, called: false}, {num: 94, called: false}, {num: 50, called: false}, {num: 73, called: false},},
			{{num: 20, called: false}, {num: 27, called: false}, {num: 34, called: false}, {num: 59, called: false}, {num: 43, called: false},},
			{{num: 66, called: false}, {num: 55, called: false}, {num: 35, called: false}, {num: 98, called: false}, {num: 57, called: false},},
			{{num: 40, called: false}, {num: 53, called: false}, {num: 21, called: false}, {num: 99, called: false}, {num: 4, called: false},},
			{{num: 17, called: false}, {num: 74, called: false}, {num: 80, called: false}, {num: 5, called: false}, {num: 12, called: false},},},

			{{{num: 76, called: false}, {num: 22, called: false}, {num: 6, called: false}, {num: 61, called: false}, {num: 23, called: false},},
			{{num: 70, called: false}, {num: 67, called: false}, {num: 69, called: false}, {num: 33, called: false}, {num: 9, called: false},},
			{{num: 87, called: false}, {num: 2, called: false}, {num: 12, called: false}, {num: 68, called: false}, {num: 27, called: false},},
			{{num: 13, called: false}, {num: 52, called: false}, {num: 82, called: false}, {num: 15, called: false}, {num: 84, called: false},},
			{{num: 24, called: false}, {num: 51, called: false}, {num: 89, called: false}, {num: 53, called: false}, {num: 38, called: false},},},

			{{{num: 96, called: false}, {num: 23, called: false}, {num: 91, called: false}, {num: 97, called: false}, {num: 10, called: false},},
			{{num: 50, called: false}, {num: 8, called: false}, {num: 68, called: false}, {num: 67, called: false}, {num: 0, called: false},},
			{{num: 65, called: false}, {num: 3, called: false}, {num: 92, called: false}, {num: 4, called: false}, {num: 70, called: false},},
			{{num: 53, called: false}, {num: 77, called: false}, {num: 59, called: false}, {num: 86, called: false}, {num: 66, called: false},},
			{{num: 41, called: false}, {num: 78, called: false}, {num: 44, called: false}, {num: 52, called: false}, {num: 71, called: false},},},

			{{{num: 62, called: false}, {num: 19, called: false}, {num: 17, called: false}, {num: 63, called: false}, {num: 75, called: false},},
			{{num: 43, called: false}, {num: 88, called: false}, {num: 15, called: false}, {num: 84, called: false}, {num: 13, called: false},},
			{{num: 41, called: false}, {num: 7, called: false}, {num: 47, called: false}, {num: 16, called: false}, {num: 23, called: false},},
			{{num: 12, called: false}, {num: 71, called: false}, {num: 8, called: false}, {num: 83, called: false}, {num: 50, called: false},},
			{{num: 36, called: false}, {num: 31, called: false}, {num: 22, called: false}, {num: 5, called: false}, {num: 79, called: false},},},

			{{{num: 71, called: false}, {num: 95, called: false}, {num: 17, called: false}, {num: 90, called: false}, {num: 63, called: false},},
			{{num: 64, called: false}, {num: 52, called: false}, {num: 32, called: false}, {num: 3, called: false}, {num: 93, called: false},},
			{{num: 70, called: false}, {num: 13, called: false}, {num: 99, called: false}, {num: 40, called: false}, {num: 5, called: false},},
			{{num: 22, called: false}, {num: 18, called: false}, {num: 83, called: false}, {num: 11, called: false}, {num: 55, called: false},},
			{{num: 47, called: false}, {num: 59, called: false}, {num: 78, called: false}, {num: 45, called: false}, {num: 29, called: false},},},

			{{{num: 9, called: false}, {num: 98, called: false}, {num: 73, called: false}, {num: 46, called: false}, {num: 79, called: false},},
			{{num: 5, called: false}, {num: 51, called: false}, {num: 84, called: false}, {num: 26, called: false}, {num: 40, called: false},},
			{{num: 64, called: false}, {num: 62, called: false}, {num: 0, called: false}, {num: 66, called: false}, {num: 18, called: false},},
			{{num: 33, called: false}, {num: 83, called: false}, {num: 47, called: false}, {num: 1, called: false}, {num: 63, called: false},},
			{{num: 89, called: false}, {num: 31, called: false}, {num: 99, called: false}, {num: 54, called: false}, {num: 36, called: false},},},

			{{{num: 98, called: false}, {num: 15, called: false}, {num: 86, called: false}, {num: 9, called: false}, {num: 50, called: false},},
			{{num: 67, called: false}, {num: 7, called: false}, {num: 75, called: false}, {num: 85, called: false}, {num: 17, called: false},},
			{{num: 96, called: false}, {num: 27, called: false}, {num: 64, called: false}, {num: 81, called: false}, {num: 19, called: false},},
			{{num: 80, called: false}, {num: 30, called: false}, {num: 29, called: false}, {num: 54, called: false}, {num: 52, called: false},},
			{{num: 49, called: false}, {num: 25, called: false}, {num: 36, called: false}, {num: 5, called: false}, {num: 90, called: false},},},

			{{{num: 39, called: false}, {num: 29, called: false}, {num: 40, called: false}, {num: 16, called: false}, {num: 69, called: false},},
			{{num: 38, called: false}, {num: 55, called: false}, {num: 67, called: false}, {num: 71, called: false}, {num: 59, called: false},},
			{{num: 42, called: false}, {num: 72, called: false}, {num: 51, called: false}, {num: 10, called: false}, {num: 45, called: false},},
			{{num: 94, called: false}, {num: 75, called: false}, {num: 21, called: false}, {num: 27, called: false}, {num: 0, called: false},},
			{{num: 84, called: false}, {num: 6, called: false}, {num: 22, called: false}, {num: 33, called: false}, {num: 30, called: false},},},

			{{{num: 33, called: false}, {num: 64, called: false}, {num: 82, called: false}, {num: 97, called: false}, {num: 39, called: false},},
			{{num: 79, called: false}, {num: 7, called: false}, {num: 62, called: false}, {num: 49, called: false}, {num: 99, called: false},},
			{{num: 26, called: false}, {num: 3, called: false}, {num: 13, called: false}, {num: 66, called: false}, {num: 10, called: false},},
			{{num: 37, called: false}, {num: 98, called: false}, {num: 15, called: false}, {num: 80, called: false}, {num: 47, called: false},},
			{{num: 1, called: false}, {num: 35, called: false}, {num: 30, called: false}, {num: 50, called: false}, {num: 43, called: false},},},

			{{{num: 56, called: false}, {num: 92, called: false}, {num: 41, called: false}, {num: 82, called: false}, {num: 34, called: false},},
			{{num: 68, called: false}, {num: 79, called: false}, {num: 11, called: false}, {num: 0, called: false}, {num: 65, called: false},},
			{{num: 70, called: false}, {num: 84, called: false}, {num: 26, called: false}, {num: 76, called: false}, {num: 96, called: false},},
			{{num: 1, called: false}, {num: 72, called: false}, {num: 31, called: false}, {num: 80, called: false}, {num: 8, called: false},},
			{{num: 9, called: false}, {num: 38, called: false}, {num: 98, called: false}, {num: 17, called: false}, {num: 7, called: false},},},

			{{{num: 12, called: false}, {num: 19, called: false}, {num: 6, called: false}, {num: 29, called: false}, {num: 89, called: false},},
			{{num: 96, called: false}, {num: 87, called: false}, {num: 70, called: false}, {num: 75, called: false}, {num: 77, called: false},},
			{{num: 84, called: false}, {num: 74, called: false}, {num: 64, called: false}, {num: 54, called: false}, {num: 13, called: false},},
			{{num: 16, called: false}, {num: 68, called: false}, {num: 44, called: false}, {num: 79, called: false}, {num: 43, called: false},},
			{{num: 61, called: false}, {num: 47, called: false}, {num: 69, called: false}, {num: 26, called: false}, {num: 50, called: false},},},

			{{{num: 43, called: false}, {num: 20, called: false}, {num: 45, called: false}, {num: 21, called: false}, {num: 87, called: false},},
			{{num: 80, called: false}, {num: 50, called: false}, {num: 83, called: false}, {num: 26, called: false}, {num: 49, called: false},},
			{{num: 64, called: false}, {num: 99, called: false}, {num: 71, called: false}, {num: 75, called: false}, {num: 9, called: false},},
			{{num: 18, called: false}, {num: 96, called: false}, {num: 6, called: false}, {num: 94, called: false}, {num: 88, called: false},},
			{{num: 76, called: false}, {num: 97, called: false}, {num: 51, called: false}, {num: 11, called: false}, {num: 74, called: false},},},

			{{{num: 85, called: false}, {num: 47, called: false}, {num: 25, called: false}, {num: 72, called: false}, {num: 93, called: false},},
			{{num: 96, called: false}, {num: 36, called: false}, {num: 81, called: false}, {num: 55, called: false}, {num: 27, called: false},},
			{{num: 63, called: false}, {num: 18, called: false}, {num: 57, called: false}, {num: 1, called: false}, {num: 29, called: false},},
			{{num: 9, called: false}, {num: 35, called: false}, {num: 83, called: false}, {num: 88, called: false}, {num: 98, called: false},},
			{{num: 90, called: false}, {num: 21, called: false}, {num: 3, called: false}, {num: 67, called: false}, {num: 82, called: false},},},

			{{{num: 58, called: false}, {num: 94, called: false}, {num: 55, called: false}, {num: 10, called: false}, {num: 98, called: false},},
			{{num: 2, called: false}, {num: 24, called: false}, {num: 71, called: false}, {num: 93, called: false}, {num: 57, called: false},},
			{{num: 74, called: false}, {num: 34, called: false}, {num: 21, called: false}, {num: 35, called: false}, {num: 73, called: false},},
			{{num: 89, called: false}, {num: 88, called: false}, {num: 6, called: false}, {num: 16, called: false}, {num: 8, called: false},},
			{{num: 76, called: false}, {num: 81, called: false}, {num: 38, called: false}, {num: 28, called: false}, {num: 83, called: false},},},

			{{{num: 36, called: false}, {num: 53, called: false}, {num: 63, called: false}, {num: 67, called: false}, {num: 18, called: false},},
			{{num: 51, called: false}, {num: 74, called: false}, {num: 60, called: false}, {num: 69, called: false}, {num: 85, called: false},},
			{{num: 32, called: false}, {num: 22, called: false}, {num: 80, called: false}, {num: 58, called: false}, {num: 98, called: false},},
			{{num: 34, called: false}, {num: 92, called: false}, {num: 13, called: false}, {num: 12, called: false}, {num: 26, called: false},},
			{{num: 46, called: false}, {num: 61, called: false}, {num: 31, called: false}, {num: 96, called: false}, {num: 47, called: false},},},

			{{{num: 90, called: false}, {num: 1, called: false}, {num: 5, called: false}, {num: 10, called: false}, {num: 48, called: false},},
			{{num: 12, called: false}, {num: 76, called: false}, {num: 95, called: false}, {num: 83, called: false}, {num: 17, called: false},},
			{{num: 24, called: false}, {num: 84, called: false}, {num: 65, called: false}, {num: 44, called: false}, {num: 28, called: false},},
			{{num: 81, called: false}, {num: 80, called: false}, {num: 41, called: false}, {num: 79, called: false}, {num: 15, called: false},},
			{{num: 29, called: false}, {num: 61, called: false}, {num: 75, called: false}, {num: 94, called: false}, {num: 40, called: false},},},

			{{{num: 65, called: false}, {num: 22, called: false}, {num: 40, called: false}, {num: 75, called: false}, {num: 86, called: false},},
			{{num: 93, called: false}, {num: 77, called: false}, {num: 46, called: false}, {num: 35, called: false}, {num: 87, called: false},},
			{{num: 88, called: false}, {num: 5, called: false}, {num: 91, called: false}, {num: 48, called: false}, {num: 74, called: false},},
			{{num: 92, called: false}, {num: 28, called: false}, {num: 66, called: false}, {num: 47, called: false}, {num: 30, called: false},},
			{{num: 69, called: false}, {num: 2, called: false}, {num: 29, called: false}, {num: 67, called: false}, {num: 94, called: false},},},

			{{{num: 78, called: false}, {num: 38, called: false}, {num: 1, called: false}, {num: 53, called: false}, {num: 68, called: false},},
			{{num: 84, called: false}, {num: 70, called: false}, {num: 26, called: false}, {num: 7, called: false}, {num: 72, called: false},},
			{{num: 92, called: false}, {num: 87, called: false}, {num: 55, called: false}, {num: 47, called: false}, {num: 6, called: false},},
			{{num: 51, called: false}, {num: 82, called: false}, {num: 36, called: false}, {num: 73, called: false}, {num: 28, called: false},},
			{{num: 75, called: false}, {num: 58, called: false}, {num: 35, called: false}, {num: 49, called: false}, {num: 56, called: false},},},

			{{{num: 61, called: false}, {num: 85, called: false}, {num: 60, called: false}, {num: 19, called: false}, {num: 24, called: false},},
			{{num: 16, called: false}, {num: 23, called: false}, {num: 71, called: false}, {num: 74, called: false}, {num: 33, called: false},},
			{{num: 42, called: false}, {num: 7, called: false}, {num: 57, called: false}, {num: 82, called: false}, {num: 70, called: false},},
			{{num: 14, called: false}, {num: 97, called: false}, {num: 59, called: false}, {num: 99, called: false}, {num: 49, called: false},},
			{{num: 46, called: false}, {num: 30, called: false}, {num: 89, called: false}, {num: 79, called: false}, {num: 41, called: false},},},

			{{{num: 5, called: false}, {num: 24, called: false}, {num: 92, called: false}, {num: 19, called: false}, {num: 65, called: false},},
			{{num: 0, called: false}, {num: 80, called: false}, {num: 33, called: false}, {num: 78, called: false}, {num: 23, called: false},},
			{{num: 4, called: false}, {num: 37, called: false}, {num: 31, called: false}, {num: 16, called: false}, {num: 41, called: false},},
			{{num: 79, called: false}, {num: 73, called: false}, {num: 88, called: false}, {num: 36, called: false}, {num: 67, called: false},},
			{{num: 86, called: false}, {num: 29, called: false}, {num: 62, called: false}, {num: 61, called: false}, {num: 71, called: false},},},

			{{{num: 40, called: false}, {num: 51, called: false}, {num: 27, called: false}, {num: 57, called: false}, {num: 85, called: false},},
			{{num: 53, called: false}, {num: 68, called: false}, {num: 31, called: false}, {num: 60, called: false}, {num: 83, called: false},},
			{{num: 48, called: false}, {num: 69, called: false}, {num: 24, called: false}, {num: 17, called: false}, {num: 96, called: false},},
			{{num: 54, called: false}, {num: 89, called: false}, {num: 22, called: false}, {num: 77, called: false}, {num: 64, called: false},},
			{{num: 95, called: false}, {num: 26, called: false}, {num: 21, called: false}, {num: 65, called: false}, {num: 41, called: false},},},

			{{{num: 48, called: false}, {num: 7, called: false}, {num: 20, called: false}, {num: 68, called: false}, {num: 21, called: false},},
			{{num: 31, called: false}, {num: 22, called: false}, {num: 1, called: false}, {num: 99, called: false}, {num: 96, called: false},},
			{{num: 82, called: false}, {num: 63, called: false}, {num: 78, called: false}, {num: 2, called: false}, {num: 70, called: false},},
			{{num: 18, called: false}, {num: 83, called: false}, {num: 58, called: false}, {num: 92, called: false}, {num: 51, called: false},},
			{{num: 81, called: false}, {num: 64, called: false}, {num: 98, called: false}, {num: 44, called: false}, {num: 89, called: false},},},

			{{{num: 23, called: false}, {num: 74, called: false}, {num: 99, called: false}, {num: 75, called: false}, {num: 6, called: false},},
			{{num: 81, called: false}, {num: 14, called: false}, {num: 37, called: false}, {num: 8, called: false}, {num: 85, called: false},},
			{{num: 12, called: false}, {num: 36, called: false}, {num: 55, called: false}, {num: 20, called: false}, {num: 47, called: false},},
			{{num: 88, called: false}, {num: 7, called: false}, {num: 90, called: false}, {num: 87, called: false}, {num: 43, called: false},},
			{{num: 3, called: false}, {num: 44, called: false}, {num: 83, called: false}, {num: 15, called: false}, {num: 2, called: false},},},

			{{{num: 97, called: false}, {num: 15, called: false}, {num: 69, called: false}, {num: 76, called: false}, {num: 95, called: false},},
			{{num: 13, called: false}, {num: 44, called: false}, {num: 31, called: false}, {num: 10, called: false}, {num: 14, called: false},},
			{{num: 40, called: false}, {num: 83, called: false}, {num: 49, called: false}, {num: 11, called: false}, {num: 65, called: false},},
			{{num: 43, called: false}, {num: 98, called: false}, {num: 55, called: false}, {num: 92, called: false}, {num: 89, called: false},},
			{{num: 90, called: false}, {num: 33, called: false}, {num: 73, called: false}, {num: 32, called: false}, {num: 53, called: false},},},

			{{{num: 25, called: false}, {num: 79, called: false}, {num: 77, called: false}, {num: 83, called: false}, {num: 68, called: false},},
			{{num: 5, called: false}, {num: 10, called: false}, {num: 23, called: false}, {num: 15, called: false}, {num: 19, called: false},},
			{{num: 18, called: false}, {num: 4, called: false}, {num: 92, called: false}, {num: 51, called: false}, {num: 76, called: false},},
			{{num: 17, called: false}, {num: 90, called: false}, {num: 70, called: false}, {num: 49, called: false}, {num: 39, called: false},},
			{{num: 74, called: false}, {num: 63, called: false}, {num: 75, called: false}, {num: 42, called: false}, {num: 67, called: false},},},

			{{{num: 46, called: false}, {num: 66, called: false}, {num: 18, called: false}, {num: 17, called: false}, {num: 28, called: false},},
			{{num: 75, called: false}, {num: 76, called: false}, {num: 78, called: false}, {num: 72, called: false}, {num: 44, called: false},},
			{{num: 57, called: false}, {num: 39, called: false}, {num: 97, called: false}, {num: 27, called: false}, {num: 99, called: false},},
			{{num: 36, called: false}, {num: 58, called: false}, {num: 62, called: false}, {num: 90, called: false}, {num: 82, called: false},},
			{{num: 14, called: false}, {num: 45, called: false}, {num: 48, called: false}, {num: 64, called: false}, {num: 1, called: false},},},

			{{{num: 51, called: false}, {num: 90, called: false}, {num: 58, called: false}, {num: 6, called: false}, {num: 37, called: false},},
			{{num: 28, called: false}, {num: 30, called: false}, {num: 57, called: false}, {num: 54, called: false}, {num: 10, called: false},},
			{{num: 45, called: false}, {num: 80, called: false}, {num: 39, called: false}, {num: 4, called: false}, {num: 0, called: false},},
			{{num: 29, called: false}, {num: 17, called: false}, {num: 66, called: false}, {num: 18, called: false}, {num: 55, called: false},},
			{{num: 96, called: false}, {num: 44, called: false}, {num: 36, called: false}, {num: 76, called: false}, {num: 34, called: false},},},

			{{{num: 94, called: false}, {num: 50, called: false}, {num: 0, called: false}, {num: 71, called: false}, {num: 99, called: false},},
			{{num: 11, called: false}, {num: 67, called: false}, {num: 96, called: false}, {num: 87, called: false}, {num: 64, called: false},},
			{{num: 48, called: false}, {num: 30, called: false}, {num: 31, called: false}, {num: 68, called: false}, {num: 40, called: false},},
			{{num: 89, called: false}, {num: 55, called: false}, {num: 23, called: false}, {num: 92, called: false}, {num: 42, called: false},},
			{{num: 16, called: false}, {num: 62, called: false}, {num: 37, called: false}, {num: 83, called: false}, {num: 33, called: false},},},

			{{{num: 66, called: false}, {num: 42, called: false}, {num: 91, called: false}, {num: 70, called: false}, {num: 72, called: false},},
			{{num: 28, called: false}, {num: 69, called: false}, {num: 96, called: false}, {num: 17, called: false}, {num: 71, called: false},},
			{{num: 99, called: false}, {num: 5, called: false}, {num: 2, called: false}, {num: 26, called: false}, {num: 19, called: false},},
			{{num: 60, called: false}, {num: 87, called: false}, {num: 51, called: false}, {num: 83, called: false}, {num: 76, called: false},},
			{{num: 77, called: false}, {num: 33, called: false}, {num: 64, called: false}, {num: 61, called: false}, {num: 54, called: false},},},

			{{{num: 61, called: false}, {num: 93, called: false}, {num: 90, called: false}, {num: 82, called: false}, {num: 88, called: false},},
			{{num: 80, called: false}, {num: 11, called: false}, {num: 25, called: false}, {num: 40, called: false}, {num: 28, called: false},},
			{{num: 60, called: false}, {num: 29, called: false}, {num: 34, called: false}, {num: 39, called: false}, {num: 21, called: false},},
			{{num: 24, called: false}, {num: 13, called: false}, {num: 72, called: false}, {num: 77, called: false}, {num: 2, called: false},},
			{{num: 19, called: false}, {num: 95, called: false}, {num: 47, called: false}, {num: 17, called: false}, {num: 0, called: false},},},

			{{{num: 86, called: false}, {num: 97, called: false}, {num: 50, called: false}, {num: 42, called: false}, {num: 87, called: false},},
			{{num: 7, called: false}, {num: 18, called: false}, {num: 80, called: false}, {num: 23, called: false}, {num: 30, called: false},},
			{{num: 41, called: false}, {num: 6, called: false}, {num: 96, called: false}, {num: 92, called: false}, {num: 98, called: false},},
			{{num: 36, called: false}, {num: 45, called: false}, {num: 77, called: false}, {num: 71, called: false}, {num: 38, called: false},},
			{{num: 19, called: false}, {num: 40, called: false}, {num: 47, called: false}, {num: 39, called: false}, {num: 53, called: false},},},

			{{{num: 83, called: false}, {num: 7, called: false}, {num: 80, called: false}, {num: 86, called: false}, {num: 65, called: false},},
			{{num: 37, called: false}, {num: 22, called: false}, {num: 19, called: false}, {num: 84, called: false}, {num: 92, called: false},},
			{{num: 29, called: false}, {num: 17, called: false}, {num: 76, called: false}, {num: 4, called: false}, {num: 33, called: false},},
			{{num: 97, called: false}, {num: 50, called: false}, {num: 1, called: false}, {num: 12, called: false}, {num: 21, called: false},},
			{{num: 15, called: false}, {num: 58, called: false}, {num: 39, called: false}, {num: 38, called: false}, {num: 74, called: false},},},
		},
		39984,
	},
}

  func TestA(t *testing.T) {
	for _, testData := range data {
	  actual := AMain(testData.input1, testData.input2)
	  if !reflect.DeepEqual(actual, testData.expected) {
		fmt.Println(actual, testData.expected)
		t.Errorf("A()")
	  }
	}
  }