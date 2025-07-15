package handlers

import (
	"database/sql"
	"log"
	"time"
)

func StartAutoRounds(db *sql.DB, interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)

			if len(AllBets) == 0 {
				continue
			}

			log.Println("⚙️ Starting automatic round")

			result := RunRound()
			if err := saveRoundResultToDB(db, result); err != nil {
				log.Printf("❌ Auto round DB save error: %v", err)
			}
		}
	}()
}
