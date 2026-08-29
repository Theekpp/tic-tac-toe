package datasource

import "tictactoe/internal/domain"

func toDomain(game *Game) *domain.Game {
	if game == nil {
		return nil
	}
	return &domain.Game{
		ID:    game.ID,
		Board: game.Board,
	}
}

func fromDomain(game *domain.Game) *Game {
	if game == nil {
		return nil
	}
	return &Game{
		ID:    game.ID,
		Board: game.Board,
	}
}
