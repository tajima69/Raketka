package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/modules/roulette/dto"
	"math/rand"
	"strings"
	"time"
)

var AllBets []dto.Bet

func PlaceBetHandler(c *fiber.Ctx) error {
	var bet dto.Bet
	if err := c.BodyParser(&bet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid bet",
		})
	}

	if bet.Color != "blue" && bet.Color != "green" && bet.Color != "red" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid color",
		})
	}

	AllBets = append(AllBets, bet)

	return c.JSON(fiber.Map{
		"message": "bet placed",
	})
}

func StartRoundHandler(c *fiber.Ctx) error {
	if len(AllBets) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "no bets placed",
		})
	}

	var sectors []string
	for i := 0; i < 15; i++ {
		sectors = append(sectors, "blue")
		sectors = append(sectors, "green")
	}
	sectors = append(sectors, "red")

	rand.Seed(time.Now().UnixNano())
	winnerIndex := rand.Intn(len(sectors))
	winnerColor := sectors[winnerIndex]

	colorTotals := map[string]float64{
		"blue":  0,
		"green": 0,
		"red":   0,
	}

	var winners []dto.WinnerResult
	for _, bet := range AllBets {
		colorTotals[bet.Color] += bet.Amount

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

	AllBets = []dto.Bet{}

	return c.JSON(dto.RoundResult{
		WinnerColor: winnerColor,
		Winners:     winners,
	})
}

func PostBetHandler(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var bet dto.Bet
	if err := c.BodyParser(&bet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}

	bet.UserID = userID
	bet.Color = strings.ToLower(bet.Color)

	validColors := map[string]bool{"blue": true, "green": true, "red": true}
	if !validColors[bet.Color] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid color"})
	}
	if bet.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid amount"})
	}

	AllBets = append(AllBets, bet)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "bet placed",
		"bet":     bet,
	})
}

func GetUserBetsHandler(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user_id is required",
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
