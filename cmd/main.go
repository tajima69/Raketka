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

	handlers.StartAutoRounds(30 * time.Second)

	DB, err := database.DbConnect()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	server.Server(DB)
}
