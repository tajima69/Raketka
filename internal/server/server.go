package server

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tajima69/Raketka/internal/modules/auth"
	"github.com/tajima69/Raketka/internal/modules/crush"
	"github.com/tajima69/Raketka/internal/modules/roulette"
)

func Server(db *sql.DB) {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${time} ${ip} - ${method} ${path} ${status} - ${latency} - error: ${error}\n",
	}))
	auth.Rout(app, db)
	roulette.Rout(app, db)
	crush.Rout(app, db)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
