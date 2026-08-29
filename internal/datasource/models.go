package datasource

import (
	"github.com/google/uuid"
	"tictactoe/internal/domain"
)

type Game struct {
	ID    uuid.UUID    `json:"id"`
	Board domain.Board `json:"board"`
}
