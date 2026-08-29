package web

import (
	"github.com/google/uuid"
	"tictactoe/internal/domain"
)

func ToDomainGame(id string, board domain.Board) (*domain.Game, error) {
	gameID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &domain.Game{
		ID:    gameID,
		Board: board,
	}, nil
}

func FromDomainGame(game *domain.Game, isFinish bool, winner string) MoveResponse {
	status := StatusActive
	if isFinish {
		if winner != "" {
			status = StatusWon
		} else {
			status = StatusDraw
		}
	}
	return MoveResponse{
		ID:     game.ID.String(),
		Board:  game.Board,
		Status: status,
		Winner: winner,
	}
}
