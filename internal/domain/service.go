package domain

import (
	"context"
	"errors"
)

var ErrInvalidBoard = errors.New("invalid board state")

type serviceImpl struct{}

func NewService() Service {
	return &serviceImpl{}
}

func (s *serviceImpl) ComputeNextMove(ctx context.Context, game *Game) (*Move, error) {
	finished, _, err := s.CheckCompletion(ctx, game)
	if err != nil {
		return nil, err
	}
	if finished {
		return nil, errors.New("game is over")
	}
	boardCopy := game.Board
	return findBestMove(&boardCopy), nil
}

func (s *serviceImpl) ValidateBoard(ctx context.Context, oldGame, newGame *Game) error {
	if newGame == nil {
		return ErrInvalidBoard
	}
	diffs := 0
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if oldGame.Board[r][c] != CellEmpty && oldGame.Board[r][c] != newGame.Board[r][c] {
				return errors.New("invalid board: previous move was altered")
			}
			if oldGame.Board[r][c] == CellEmpty && newGame.Board[r][c] != CellEmpty {
				if newGame.Board[r][c] != CellX {
					return errors.New("invalid board: only player X can make a move")
				}
				diffs++
			}
		}
	}
	if diffs != 1 {
		return errors.New("invalid board: must place exactly one mark")
	}
	return nil
}

func (s *serviceImpl) CheckCompletion(ctx context.Context, game *Game) (bool, string, error) {
	score := evaluateBoard(&game.Board)
	if score == 10 {
		return true, "O", nil
	} else if score == -10 {
		return true, "X", nil
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board[i][j] == CellEmpty {
				return false, "", nil
			}
		}
	}
	return true, "", nil
}
