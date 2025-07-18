package crush

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/middleware"
	"github.com/tajima69/Raketka/internal/modules/crush/handlers"
)

func Rout(app *fiber.App, db *sql.DB) {
	app.Post("/crush", middleware.Protected(), handlers.CrashPostBet)
}
