package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

type Board struct {
	m, n int
	body [][]bool
}

func createBoard(m int,n int) *Board {
	board := make([][]bool, m)
	for i := range m {
		board[i] := make([]bool, n)
	}
	return &Board{m:m, n:n, body:board}
}

func setCell(board *Board) Set(x int, y int, state bool) {
	board.body[x][y] = state
}

func life(board) {

}