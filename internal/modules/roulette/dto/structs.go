package dto

import "time"

type Users struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password,omitempty"`
	Balance  float64 `json:"balance"`
}

type Games struct {
	ID           int       `json:"id"`
	started_at   time.Time `json:"started_at"`
	result_color string    `json:"result_color"`
	total_bank   int       `json:"bank"`
}

type Bet struct {
	UserID string  `json:"id"`
	Color  string  `json:"color"`
	Amount float64 `json:"amount"`
}

type WinnerResult struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type RoundResult struct {
	WinnerColor string         `json:"winner_color"`
	Winners     []WinnerResult `json:"winners"`
}
