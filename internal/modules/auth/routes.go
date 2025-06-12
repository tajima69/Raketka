package auth

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/modules/auth/handlers"
)

func Rout(app *fiber.App, database *sql.DB) {
	Handler := handlers.Handler{Db: database}
	app.Post("/register", Handler.AuthPostHandler)
	app.Post("/login", Handler.LoginPostHandler)
	app.Get("/register/:id", Handler.AuthGetHandler)
}
