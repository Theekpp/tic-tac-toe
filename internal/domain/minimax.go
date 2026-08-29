package domain

func evaluateBoard(board *Board) int {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			if board[i][0] == CellO {
				return 10
			} else if board[i][0] == CellX {
				return -10
			}
		}
	}

	for j := 0; j < 3; j++ {
		if board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			if board[0][j] == CellO {
				return 10
			} else if board[0][j] == CellX {
				return -10
			}
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != CellEmpty {
		if board[1][1] == CellO {
			return 10
		} else if board[1][1] == CellX {
			return -10
		}
	}

	if board[2][0] == board[1][1] && board[1][1] == board[0][2] && board[0][0] != CellEmpty {
		if board[1][1] == CellO {
			return 10
		} else if board[1][1] == CellX {
			return -10
		}
	}

	return 0
}

func findBestMove(board *Board) *Move {
	bestVal := -1000
	bestMove := &Move{-1, -1}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == CellEmpty {
				board[i][j] = CellO
				moveVal := minimax(board, 0, false)
				board[i][j] = CellEmpty

				if moveVal > bestVal {
					bestVal = moveVal
					bestMove.Row = i
					bestMove.Col = j
				}
			}
		}
	}
	return bestMove
}

func minimax(board *Board, depth int, isMax bool) int {
	score := evaluateBoard(board)
	if score == 10 {
		return score - depth
	} else if score == -10 {
		return score + depth
	}
	hasEmpty := false
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == CellEmpty {
				hasEmpty = true
				break
			}
		}
	}
	if !hasEmpty {
		return 0
	}

	if isMax {
		best := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == CellEmpty {
					board[i][j] = CellO
					best = maxVal(best, minimax(board, depth+1, isMax))
					board[i][j] = CellEmpty
				}
			}
		}
		return best
	}

	best := 1000
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == CellEmpty {
				board[i][j] = CellX
				best = minVal(best, minimax(board, depth+1, isMax))
				board[i][j] = CellEmpty
			}
		}
	}
	return best
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FindBestMove(board *Board) *Move {
	return findBestMove(board)
}

func EvaluateBoard(board *Board) int {
	return evaluateBoard(board)
}
