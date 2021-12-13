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

func callNumber(number int, boards [][5][5]Cell) int {
	for b,board := range boards {
		bingo, newBoard := markBoard(board, number)
		boards[b] = newBoard
		if bingo > 0 {
			return bingo * number
		}
	}
	return 0
}

func AMain(nums []int, boards [][5][5]Cell) int {
	for _,number := range nums {
		bingo := callNumber(number, boards)
		if bingo > 0 {
			return bingo
		}
	}
	return 0
}
