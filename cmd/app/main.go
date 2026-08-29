package main

import (
	"context"
	"go.uber.org/fx"
	"log"
	"net/http"
	"tictactoe/internal/di"
	"tictactoe/internal/web"
)

func main() {
	fx.New(
		di.Module,
		fx.Invoke(runServer),
	).Run()
}

func runServer(lc fx.Lifecycle, router *web.Router) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting server on :8080")
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server")
			return srv.Shutdown(ctx)
		},
	})
}
