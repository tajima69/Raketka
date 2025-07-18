package handlers

import (
	"database/sql"
	"github.com/tajima69/Raketka/internal/modules/crush/dto"
	"log"
	"math"
	"math/rand"
	"time"
)

func StartCrushRounds(db *sql.DB) {
	for {
		time.Sleep(20 * time.Second)

		crashPoint := generateCrashPoint()
		log.Printf("🧨 Crash point: %.2fx", crashPoint)

		var roundID int
		err := db.QueryRow(`INSERT INTO crash_rounds (crash_point) VALUES ($1) RETURNING id`, crashPoint).Scan(&roundID)
		if err != nil {
			log.Printf("DB error: %v", err)
			continue
		}

		for _, bet := range currentBets {
			won := bet.AutoCashout <= crashPoint
			_, err := db.Exec(`INSERT INTO crash_bets (round_id, user_id, amount, auto_cashout, won) 
				VALUES ($1, $2, $3, $4, $5)`,
				roundID, bet.UserID, bet.Amount, bet.AutoCashout, won)
			if err != nil {
				log.Printf("Error saving bet: %v", err)
			}
		}

		currentBets = []dto.CrashBet{} // очистить
	}
}

func generateCrashPoint() float64 {
	min := 1.00
	max := 1000.00
	rand.Seed(time.Now().UnixNano())
	return math.Round((min+rand.Float64()*(max-min))*100) / 100
}
