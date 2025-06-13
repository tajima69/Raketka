package auth

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/middleware"
	"github.com/tajima69/Raketka/internal/modules/auth/handlers"
)

func Rout(app *fiber.App, database *sql.DB) {
	handler := handlers.Handler{Db: database}

	app.Post("/register", handler.AuthPostHandler)
	app.Post("/login", handler.LoginPostHandler)

	authGroup := app.Group("/", middleware.Protected())
	authGroup.Get("me", handler.AuthGetHandler)
}
