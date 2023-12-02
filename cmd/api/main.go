package main

import (
	"context"
	"log"

	"github.com/bxxf/go-template/internal/auth/router"
	"github.com/bxxf/go-template/internal/config"
	"github.com/bxxf/go-template/internal/logger"
	"github.com/bxxf/go-template/internal/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			logger.NewLogger,
			config.NewConfig,
			router.NewAuthRouter,
			server.NewServer,
		),
		fx.Invoke(
			func(s *server.Server) {
			},
			func(c *config.Config) {
			}),
	)

	ctx := context.Background()
	// start the application
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	// wait for the application to stop
	<-app.Done()

	// stop the application
	if err := app.Stop(ctx); err != nil {
		log.Fatal(err)
	}

}
