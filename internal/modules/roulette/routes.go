package roulette

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/middleware"
	"github.com/tajima69/Raketka/internal/modules/roulette/handlers"
)

func Rout(app *fiber.App, db *sql.DB) {
	app.Post("/start-round", middleware.Protected(), handlers.StartRoundHandler(db))
	app.Post("/bets", middleware.Protected(), handlers.PostBetHandler)
	app.Get("/user-bets", middleware.Protected(), handlers.GetUserBetsHandler)

	app.Get("/lastresult", handlers.GetLastResultHandler)
	app.Get("/results", handlers.GetAllResultsHandler(db))
}
