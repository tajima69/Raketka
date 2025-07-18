package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tajima69/Raketka/internal/modules/crush/dto"
)

var (
	currentBets []dto.CrashBet
)

func CrashPostBet(c *fiber.Ctx) error {
	var bet dto.CrashBet
	if err := c.BodyParser(&bet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	if bet.Amount <= 0 || bet.AutoCashout < 1.0 || bet.AutoCashout > 1000.0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid bet"})
	}

	bet.UserID = c.Locals("userID").(int)
	currentBets = append(currentBets, bet)

	return c.JSON(fiber.Map{"message": "bet accepted"})
}
