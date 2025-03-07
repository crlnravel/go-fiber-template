package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/crlnravel/go-fiber-template/internal/test"
)

type appConfig struct {
	prod bool
	port int
	db   *pgxpool.Pool
}

func NewApp(cfg *appConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork: cfg.prod,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
	}))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("hello from minecart!")
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		if err := c.JSON(fiber.Map{
			"message": "🐣 v1",
		}); err != nil {
			return err
		}
		return c.Next()
	})

	// your routes here

	test.App(v1, cfg.db)

	return app
}
