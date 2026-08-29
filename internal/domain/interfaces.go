package domain

import "context"

type Service interface {
	ComputeNextMove(ctx context.Context, game *Game) (*Move, error)
	ValidateBoard(ctx context.Context, oldGame, newGame *Game) error
	CheckCompletion(ctx context.Context, game *Game) (bool, string, error)
}

type Repository interface {
	Save(ctx context.Context, game *Game) error
	Get(ctx context.Context, id string) (*Game, error)
}
