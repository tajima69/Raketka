package roulette

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/modules/roulette/handlers"
)

func Rout(app *fiber.App, database *sql.DB) {
	SpinResult := handlers.SpinResult{}
	app.Post("/bet", handlers.PlaceBetHandler)
	app.Post("/round", handlers.StartRoundHandler)
	app.Get("/spin", SpinResult.SpinWheelHandler)
}
