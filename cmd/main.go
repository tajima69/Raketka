package main

import (
	"context"
	"github.com/tajima69/Raketka/database"
	"github.com/tajima69/Raketka/internal/server"
)

func main() {
	DB, err := database.DbConnect()
	if err != nil {
		panic(err)
	}
	DB.Conn(context.Background())
	defer DB.Close()
	server.Server(DB)
}
