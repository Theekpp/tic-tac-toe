package domain

import "github.com/google/uuid"

type Cell int

const (
	CellEmpty Cell = 0
	CellX     Cell = 1
	CellO     Cell = 2
)

type Board [3][3]Cell

type Game struct {
	ID    uuid.UUID `json:"id"`
	Board Board     `json:"board"`
}

type Move struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
