package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func App(app fiber.Router, db *pgxpool.Pool) {
	r := app.Group("/test")

	ctr := NewController()

	// Route listing

	r.Get("/", ctr.GetTestMsg)
}
