package handlers

import (
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"time"
)

type SpinResult struct {
	Sectors     []string `json:"sectors"`
	WinnerIndex int      `json:"winner_index"`
	WinnerColor string   `json:"winner_color"`
}

func (s *SpinResult) SpinWheelHandler(c *fiber.Ctx) error {
	rand.Seed(time.Now().UnixNano())

	var sectors []string

	for i := 0; i < 15; i++ {
		sectors = append(sectors, "blue")
	}

	for i := 0; i < 15; i++ {
		sectors = append(sectors, "green")
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
