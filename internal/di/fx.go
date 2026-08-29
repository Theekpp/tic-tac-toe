package di

import (
	"go.uber.org/fx"
	"tictactoe/internal/datasource"
	"tictactoe/internal/domain"
	"tictactoe/internal/web"
)

var Module = fx.Options(
	fx.Provide(
		datasource.NewStorage,
		datasource.NewRepository,
		domain.NewService,
		web.NewHandler,
		web.NewRouter,
	),
)
