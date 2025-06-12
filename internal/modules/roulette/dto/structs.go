package dto

import "time"

type Users struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password,omitempty"`
	Balance  float64 `json:"balance"`
}

type games struct {
	ID           int       `json:"id"`
	started_at   time.Time `json:"started_at"`
	result_color string    `json:"result_color"`
	total_bank   int       `json:"bank"`
}

type bets struct {
	ID int `json:"id"`
}
