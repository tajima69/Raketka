package routes

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/modules/roulette/handlers"
)

func Rout(app *fiber.App, database *sql.DB) {
	SpinResult := handlers.SpinResult{}

	app.Get("/spin", SpinResult.SpinWheelHandler)
}
