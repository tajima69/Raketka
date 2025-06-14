package handlers

import (
	"log"
	"time"
)

func StartAutoRounds(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)

			if len(AllBets) == 0 {
				continue
			}

			log.Println("Starting automatic round")
			RunRound()
		}
	}()
}
