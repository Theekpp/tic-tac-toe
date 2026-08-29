package web

import (
	"tictactoe/internal/domain"
)

type MoveRequest struct {
	Board domain.Board `json:"board"`
}

type MoveResponse struct {
	ID     string       `json:"id"`
	Board  domain.Board `json:"board"`
	Status GameStatus   `json:"status"`
	Winner string       `json:"winner,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type GameStatus string

const (
	StatusActive GameStatus = "active"
	StatusWon    GameStatus = "won"
	StatusDraw   GameStatus = "draw"
)
