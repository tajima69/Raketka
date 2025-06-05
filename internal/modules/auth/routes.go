package auth

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func Rout(app *fiber.App, database *sql.DB) {
	Handler := Handler{Db: database}
	app.Post("/login", Handler.AuthPostHandler)
	app.Get("/login/:id", Handler.AuthGetHandler)
}
