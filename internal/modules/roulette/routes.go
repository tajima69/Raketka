package roulette

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/middleware"
	"github.com/tajima69/Raketka/internal/modules/roulette/handlers"
)

func Rout(app *fiber.App) {
	app.Post("/start-round", middleware.Protected(), handlers.StartRoundHandler)
	app.Post("/bets", middleware.Protected(), handlers.PostBetHandler)
	app.Get("/user-bets", middleware.Protected(), handlers.GetUserBetsHandler)

	app.Get("/result", handlers.GetLastResultHandler)
}
