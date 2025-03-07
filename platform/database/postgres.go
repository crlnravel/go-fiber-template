package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/crlnravel/go-fiber-template/internal/config"
)

var DB *pgxpool.Pool

func ConnectPostgres(ctx context.Context) {
	dsn := config.GetEnv("DATABASE_URL", "")
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB = conn

	log.Print("Successfully connected to database")
}
