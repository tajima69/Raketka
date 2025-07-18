package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/modules/roulette/dto"
	"log"
	"math/rand"
	"strings"
	"time"
)

var LastRoundResult dto.RoundResult
var AllBets []dto.Bet

func PostBetHandler(c *fiber.Ctx) error {
	userIDRaw := c.Locals("userID")
	userID, ok := userIDRaw.(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid user ID",
		})
	}

	var bet dto.Bet
	if err := c.BodyParser(&bet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body",
		})
	}

	bet.UserID = userID
	bet.Color = strings.ToLower(bet.Color)

	validColors := map[string]bool{"blue": true, "green": true, "red": true}
	if !validColors[bet.Color] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid color",
		})
	}
	if bet.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid amount",
		})
	}

	AllBets = append(AllBets, bet)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "bet placed",
		"bet":     bet,
	})
}

func StartRoundHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(AllBets) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "no bets placed",
			})
		}

		result := RunRound()
		LastRoundResult = result

		if err := saveRoundResultToDB(db, result); err != nil {
			log.Printf("DB insert error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save result"})
		}

		AllBets = nil

		return c.JSON(result)
	}
}

func GetUserBetsHandler(c *fiber.Ctx) error {
	userIDRaw := c.Locals("userID")
	userID, ok := userIDRaw.(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	var userBets []dto.Bet
	for _, bet := range AllBets {
		if bet.UserID == userID {
			userBets = append(userBets, bet)
		}
	}
	return c.JSON(userBets)
}

func GetLastResultHandler(c *fiber.Ctx) error {
	if LastRoundResult.WinnerColor == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "no round played yet",
		})
	}
	return c.JSON(LastRoundResult)
}

func RunRound() dto.RoundResult {
	sectors := make([]string, 0, 31)
	for i := 0; i < 15; i++ {
		sectors = append(sectors, "blue", "green")
	}
	sectors = append(sectors, "red")

	rand.Seed(time.Now().UnixNano())
	winnerColor := sectors[rand.Intn(len(sectors))]

	var winners []dto.WinnerResult
	for _, bet := range AllBets {
		if bet.Color == winnerColor {
			multiplier := 2.0
			if winnerColor == "red" {
				multiplier = 14.0
			}
			winners = append(winners, dto.WinnerResult{
				UserID: bet.UserID,
				Amount: bet.Amount * multiplier,
			})
		}
	}

	result := dto.RoundResult{
		WinnerColor: winnerColor,
		Winners:     winners,
	}
	LastRoundResult = result
	AllBets = nil

	log.Printf("🎯 Round completed! Winner color: %s | Winners: %+v\n", winnerColor, winners)

	return result
}

func GetAllResultsHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query(`SELECT id, winner_color, winners, created_at FROM roulette_results ORDER BY created_at DESC`)
		if err != nil {
			log.Printf("DB query error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to query results"})
		}
		defer rows.Close()

		var results []dto.RoundResultDB

		for rows.Next() {
			var r dto.RoundResultDB
			var winnersJSON []byte

			err := rows.Scan(&r.ID, &r.WinnerColor, &winnersJSON, &r.CreatedAt)
			if err != nil {
				log.Printf("Row scan error: %v", err)
				continue
			}

			err = json.Unmarshal(winnersJSON, &r.Winners)
			if err != nil {
				log.Printf("Unmarshal error: %v", err)
				continue
			}

			results = append(results, r)
		}

		return c.JSON(results)
	}
}
func saveRoundResultToDB(db *sql.DB, result dto.RoundResult) error {
	if result.Winners == nil {
		result.Winners = []dto.WinnerResult{}
	}
	winnersJSON, err := json.Marshal(result.Winners)
	if err != nil {
		return err
	}

	query := `INSERT INTO roulette_results (winner_color, winners) VALUES ($1, $2)`
	_, err = db.Exec(query, result.WinnerColor, winnersJSON)
	return err
}
