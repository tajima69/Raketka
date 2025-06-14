package roulette

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/middleware"
	"github.com/tajima69/Raketka/internal/modules/roulette/handlers"
)

func Rout(app *fiber.App, database *sql.DB) {
	app.Post("/start-round", handlers.StartRoundHandler)
	app.Post("/bets", middleware.Protected(), handlers.PostBetHandler)
	app.Get("/bet", handlers.GetUserBetsHandler)
	app.Get("/result", handlers.GetLastResultHandler)
}
