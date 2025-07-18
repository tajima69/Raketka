package dto

type CrashBet struct {
	UserID      int     `json:"user_id"`
	Amount      float64 `json:"amount"`
	AutoCashout float64 `json:"auto_cashout"`
}
