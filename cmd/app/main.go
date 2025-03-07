package main

import (
	"context"
	"log"
	"time"

	"github.com/crlnravel/go-fiber-template/internal/config"
	"github.com/crlnravel/go-fiber-template/platform/database"
)

// @title COMPFEST 17 Mailer "Minecart" API
// @version 1.0
// @description API Endpoints for COMPFEST Mailer System. Built with Golang, Fiber, and Love <3
// @termsOfService
// @contact.name API Maintainer (Tim DevOps COMPFEST 17)
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host github.com/crlnravel/go-fiber-template
// @BasePath /
func main() {
	dbtype := config.GetEnv("DATABASE_TYPE", "pgx")
	ctx := context.Background()

	if dbtype == "pgx" {
		database.ConnectPostgres(ctx)
	} else {
		panic("invalid database type")
	}

	prod := config.GetStageStatus() == config.EnvironmentProduction

	app := NewApp(&appConfig{
		prod: prod,
		db:   database.DB,
	})

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			database.DB.Close()
			return nil
		},
		"http-server": func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	port := ":" + config.GetEnv("SERVER_PORT", "8080")

	if err := app.Listen(port); err != nil {
		log.Panic(err)
	}

	<-wait
}
