package datasource

import "sync"

type Storage struct {
	games sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Save(game *Game) error {
	s.games.Store(game.ID.String(), game)
	return nil
}

func (s *Storage) Get(id string) (*Game, bool) {
	value, ok := s.games.Load(id)
	if !ok {
		return nil, false
	}
	game, ok := value.(*Game)
	if !ok {
		return nil, false
	}
	return game, true
}

func (s *Storage) Delete(id string) error {
	s.games.Delete(id)
	return nil
}
