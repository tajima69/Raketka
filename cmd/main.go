package main

import (
	"github.com/joho/godotenv"
	"github.com/tajima69/Raketka/database"
	"github.com/tajima69/Raketka/internal/modules/roulette/handlers"
	"github.com/tajima69/Raketka/internal/server"
	"log"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	DB, err := database.DbConnect()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	handlers.StartAutoRounds(DB, 10*time.Second)

	server.Server(DB)
}
