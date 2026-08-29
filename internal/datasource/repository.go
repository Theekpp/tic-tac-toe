package datasource

import (
	"context"
	"errors"
	"tictactoe/internal/domain"
)

var ErrGameNotFound = errors.New("game not found")

type repositoryImpl struct {
	storage *Storage
}

func NewRepository(storage *Storage) domain.Repository {
	return &repositoryImpl{storage: storage}
}

func (r *repositoryImpl) Save(ctx context.Context, game *domain.Game) error {
	return r.storage.Save(fromDomain(game))
}

func (r *repositoryImpl) Get(ctx context.Context, id string) (*domain.Game, error) {
	game, ok := r.storage.Get(id)
	if !ok {
		return nil, ErrGameNotFound
	}
	return toDomain(game), nil
}
