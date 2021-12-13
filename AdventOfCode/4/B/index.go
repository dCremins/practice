package main


type Cell struct {
    num		int
    called	bool
}


func markBoard(board [5][5]Cell, number int) (int, [5][5]Cell) {
	verticals := [5][]int{}
	sum := 0
	winning := false
	for r,row := range board {
		horizontal := []int{}
		for c, cell := range row {
			if cell.num == number {
				cell.called = true
				board[r][c] = cell
			}
			if cell.called {
				verticals[c] = append(verticals[c], cell.num)
				horizontal = append(horizontal, cell.num)
			} else {
				sum+=cell.num
			}
			if len(verticals[c]) == 5 {
				winning = true
			}
		}
		if len(horizontal) == 5 {
			winning = true
		}
	}
	if winning {
		return sum, board
	}
	return 0, board
}

func callNumber(number int, boards [][5][5]Cell) ([][5][5]Cell, int) {
	// last := 0
	lastSum := 0
	newBoards := [][5][5]Cell{}
	for b,board := range boards {
		bingo, newBoard := markBoard(board, number)
		boards[b] = newBoard
		if bingo > 0 {
			lastSum = bingo * number
		} else {
			newBoards = append(newBoards, newBoard)
		}
	}
	return newBoards, lastSum
}

func BMain(nums []int, boards [][5][5]Cell) int {
	// last := 0
	lastSum := 0
	for _,number := range nums {
		b, bingo := callNumber(number, boards)
		if bingo > 0 {
			boards = b 
			lastSum = bingo
		}
	}
	return lastSum
}