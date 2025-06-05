package database

import (
	"database/sql"
	"fmt"
	"time"
)

func DbConnect() (*sql.DB, error) {
	var db *sql.DB
	user := "postgres"
	password := "postgres"
	host := "localhost"
	port := "5432"
	dbname := "todo"
	sslmode := "disable"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}
