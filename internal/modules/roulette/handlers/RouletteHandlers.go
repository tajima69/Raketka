package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/modules/roulette/dto"
	"math/rand"
	"time"
)

type SpinResult struct {
	Sectors     []string `json:"sectors"`
	WinnerIndex int      `json:"winner_index"`
	WinnerColor string   `json:"winner_color"`
}

var allBets []dto.Bet

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

	allBets = append(allBets, bet)

	return c.JSON(fiber.Map{
		"message": "bet placed",
	})
}

func StartRoundHandler(c *fiber.Ctx) error {
	if len(allBets) == 0 {
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

	var winners []dto.WinnerResult
	for _, bet := range allBets {
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

	allBets = []dto.Bet{}

	return c.JSON(dto.RoundResult{
		WinnerColor: winnerColor,
		Winners:     winners,
	})
}

func (s *SpinResult) SpinWheelHandler(c *fiber.Ctx) error {
	rand.Seed(time.Now().UnixNano())

	sectors := make([]string, 0, 31)

	for i := 0; i < 15; i++ {
		sectors = append(sectors, "blue", "green")
	}

	sectors = append(sectors, "red")

	rand.Shuffle(len(sectors), func(i, j int) {
		sectors[i], sectors[j] = sectors[j], sectors[i]
	})

	winnerIndex := rand.Intn(len(sectors))
	winnerColor := sectors[winnerIndex]

	result := SpinResult{
		Sectors:     sectors,
		WinnerIndex: winnerIndex,
		WinnerColor: winnerColor,
	}
	return c.JSON(result)
}
