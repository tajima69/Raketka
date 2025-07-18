package dto

type Users struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password,omitempty"`
	Balance  float64 `json:"balance"`
}

type Bet struct {
	UserID int     `json:"id"`
	Color  string  `json:"color"`
	Amount float64 `json:"amount"`
}

type WinnerResult struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

type RoundResult struct {
	WinnerColor string         `json:"winner_color"`
	Winners     []WinnerResult `json:"winners"`
}

type RoundResultDB struct {
	ID          int            `json:"id"`
	WinnerColor string         `json:"winner_color"`
	Winners     []WinnerResult `json:"winners"`
	CreatedAt   string         `json:"created_at"`
}
