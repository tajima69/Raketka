package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/tajima69/Raketka/internal/middleware"
	"github.com/tajima69/Raketka/internal/modules/auth/dto"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Handler struct {
	Db *sql.DB
}

func (h *Handler) AuthGetHandler(c *fiber.Ctx) error {
	var user dto.Users

	id := c.Params("id")
	query := "SELECT id, username, password FROM users WHERE id = $1"

	row := h.Db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	} else if err != nil {
		log.Printf("DB scan error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to read user",
		})
	}

	user.Password = ""

	return c.JSON(user)
}

func (h *Handler) AuthPostHandler(c *fiber.Ctx) error {
	var user dto.Users
	if err := c.BodyParser(&user); err != nil {
		log.Printf("Body parse error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username and password are required",
		})
	}

	var exists bool
	err := h.Db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)", user.Username).Scan(&exists)
	if err != nil {
		log.Printf("Check user exists error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal error",
		})
	}
	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "username already exists",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password hash error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not hash password",
		})
	}

	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err = h.Db.QueryRow(query, user.Username, string(hashedPassword)).Scan(&user.ID)
	if err != nil {
		log.Printf("Insert error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to register user",
		})
	}

	token, err := middleware.GenerateJWT(user.ID, user.Username)
	if err != nil {
		log.Printf("JWT generation error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func (h *Handler) LoginPostHandler(c *fiber.Ctx) error {
	var creds dto.Users
	if err := c.BodyParser(&creds); err != nil {
		log.Printf("Body parse error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	if creds.Username == "" || creds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username and password are required",
		})
	}

	var user dto.Users
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := h.Db.QueryRow(query, creds.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid username or password",
		})
	} else if err != nil {
		log.Printf("DB query error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal error",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid username or password",
		})
	}

	token, err := middleware.GenerateJWT(user.ID, user.Username)
	if err != nil {
		log.Printf("JWT generation error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
